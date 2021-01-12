package main

import (
  "encoding/binary"
  "fmt"
)

func BigEndian() {
  // 大端序
  // 二进制: 0000 0001 0000 0010 0000 0011 0000 0100， 一共32bit， 4byte
  var testInt int32 = 0x010203040     // 十六进制表示
  fmt.Printf("%d 使用大端位: \n", testInt)

  var testBytes []byte = make([]byte, 4)
  binary.BigEndian.PutUint32(testBytes, uint32(testInt))    // 大端位模式
  fmt.Println("int32 转为 bytes", testBytes)


}

func main() {
  BigEndian()
}