package main

import (
  "bytes"
  "fmt"
  "runtime"
  "strconv"
  "sync"
)

var goroutineSpace = []byte("goroutine ")

func curGoroutineID() uint64 {
  b := make([]byte, 64)
  b = b[:runtime.Stack(b, false)]
  // Parse the 4707 out of "goroutine 4707 ["
  b = bytes.TrimPrefix(b, goroutineSpace)
  i := bytes.IndexByte(b, ' ')
  if i < 0 {
    panic(any(fmt.Sprintf("No space found in %q", b)))
  }

  b = b[:i]
  n, err := strconv.ParseUint(string(b), 10, 64)
  if err != nil {
    panic(any(fmt.Sprintf("Failed to parse goroutine ID out of %q, %v", b, err)))
  }
  return n
}

/*
func Tracex() func() {
  pc, _, _, ok := runtime.Caller(1)
  if !ok {
    panic(any("not found caller"))
  }

  fn := runtime.FuncForPC(pc)
  name := fn.Name()

  gid := curGoroutineID()
  fmt.Printf("g[%05d]: enter: [%s]\n", gid, name)
  return func() { fmt.Printf("g[%05d]: exit: [%s]\n", gid, name) }
}
 */

//加入printTrace做到格式化输出
func printTrace(id uint64, name, arrow string, indent int) {
  indents := ""
  for i := 0; i < indent; i++ {
    //indents += "    "
    indents += "|    "
  }
  fmt.Printf("g[%05d]:%s%s%s\n", id, indents, arrow, name)

  // TODO: show the file path
  //c := stack.Caller(0)
  //c := stack.Caller(1)
  //fmt.Printf("g[%05d]:%s%s%s        %+v\n", id, indents, arrow, name, fmt.Sprintf("%+v", c))
}

var mu sync.Mutex
var m = make(map[uint64]int)

func Tracex() func() {
  pc, _, _, ok := runtime.Caller(1)
  if !ok {
    panic(any("not found caller"))
  }

  fn := runtime.FuncForPC(pc)
  name := fn.Name()
  gid := curGoroutineID()

  mu.Lock()
  indents := m[gid]    // 获取当前gid对应的缩进层次
  m[gid] = indents + 1 // 缩进层次+1后存入map
  mu.Unlock()

  //printTrace(gid, name, "->", indents+1)
  printTrace(gid, name, "|-", indents+1)
  return func() {
    mu.Lock()
    indents := m[gid]    // 获取当前gid对应的缩进层次
    m[gid] = indents - 1 // 缩进层次-1后存入map
    mu.Unlock()
    //printTrace(gid, name, "<-", indents)
    printTrace(gid, name, "|-", indents)
  }
}

func A1() {
  defer Tracex()()
  B1()
}

func B1() {
  defer Tracex()()
  C1()
}

func C1() {
  defer Tracex()()
  D()
}

func D() {
  defer Tracex()()
}

func A2() {
  defer Tracex()()
  B2()
}
func B2() {
  defer Tracex()()
  C2()
}
func C2() {
  defer Tracex()()
  D()
}

func main() {
  var wg sync.WaitGroup
  wg.Add(1)

  go func() {
    A2()
    wg.Done()
  }()

  A1()
  wg.Wait()
}









