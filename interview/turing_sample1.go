package main

import (
  "log"
  "sync"
  "time"
)

func main() {
  //c1()
  //c2()
  //c3()
  //c4()
  //c5()
  //c6()
  //c7()
  //c8()
  //c9()

  /*
  m := make(map[int]int)
  wg := &sync.WaitGroup{}
  mu := &sync.Mutex{}
  wg.Add(10)

  for i := 0; i < 10; i++ {
    go func() {
      defer wg.Done()
      mu.Lock()
      m[i] = i
      mu.Unlock()
    }()
  }

  wg.Wait()
  log.Println(m)
  log.Println(len(m))
   */

  // ------------------------------------------
  a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
  log.Printf("point outer makeSquares -->>> %+v", &a[0])
  makeSquares(a)
  log.Println(a)
}

func c1() {
  var s []int
  a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
  s = a[2:4]
  log.Println(s)
  log.Println(cap(s))
  log.Println(cap(a))
}

func c2() {
  for i := 0; i < 3; i++ {
    break
    for j := 0; j < 3; j++ {
      print(i, ",", j, " ")
      break
      //break
    }
    continue
    println()
  }
}

func c3() {
  a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
  s := a[2: 4]
  s[0] = 33
  log.Println(a[2])
}

// --------------------------------------------
type Shape interface {
  Area() float64
  Perimeter() float64
}

type Rect struct {
  width float64
  height float64
}

func (r Rect) Area() float64 {
  return r.width * r.height
}

func (r Rect) Perimeter() float64 {
  return 2 * (r.width + r.height)
}

func c4() {
  var s Shape
  s = Rect{5.0, 4.0}
  r := Rect{5.0, 4.0}
  log.Println(s == r)
}

func c5() {
  a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
  s := a[2: 4]
  newS := append(s, 55, 66)
  log.Println(newS)
  log.Println(len(newS), cap(newS))
}

// -----------------------------------------
const N = 3
func c6() {
  m := make(map[int]*int)

  for i := 0; i < N; i++ {
    // todo: the point point to the `i` varible, at the end the `i` is what, the point will be the same memory address
    log.Printf("%+v --- %d", &i, i)   // todo: every loop `i` is the same varible, `&i` is the same memory address
    m[i] = &i
  }

  log.Printf("%+v", m)
  for _, v := range m {
    log.Println(*v)
  }
}

// ------------------------------------------
var start time.Time

func init() {
  start = time.Now()
}

func service1(c chan string) {
  c <- "Hello from service1"
}

func service2(c chan string) {
  c <- "Hello from service2"
}

func c7() {
  chan1 := make(chan string)
  chan2 := make(chan string)

  go service1(chan1)
  go service2(chan2)

  // todo: the select for listening channel, forever prcess `the last match case`, here always `<-chan2`
  // todo: if has `default`, always go to `default`
  select {
  case res := <-chan1:
    log.Println("Response from service 1", res)
  case res := <-chan2:
   log.Println("Response from service 2", res)
  default:
   log.Println("No response service")
  }

  time.Sleep(3 * time.Second)
}

func c8() {
  m := make(map[int]int)
  wg := &sync.WaitGroup{}
  mu := &sync.Mutex{}
  wg.Add(10)

  for i := 0; i < 10; i++ {
    go func() {
      defer wg.Done()
      mu.Lock()
      m[i] = i
      mu.Unlock()
    }()
  }

  wg.Wait()
  log.Println(m)
  log.Println(len(m))
}

func makeSquares(array [10]int) {
  log.Printf("point in makeSquares -->>> %+v", &array[0])
  for index, elem := range array {
    array[index] = elem * elem
  }
}

func c9() {
  a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
  // todo: below just only change the `a` in func scope, because the `a` varible is not pass from outer scope
  makeSquares(a)
  log.Println(a)
}


