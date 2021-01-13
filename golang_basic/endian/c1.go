package main

import (
  "encoding/binary"
  "fmt"
)

func BigEndian() {
  // 大端序
  // 二进制: 0000 0001 - 0000 0010 - 0000 0011 - 0000 0100， 一共32bit， 4byte
  var testInt int32 = 0x01020304     // 十六进制表示
  fmt.Printf("使用大端位: %d \n", testInt)

  var testBytes = make([]byte, 4)
  binary.BigEndian.PutUint32(testBytes, uint32(testInt))
  // todo: 输出 [1 2 3 4],  下标0(对应数据1）是内存最低位， 所以储存了数据最高位 0x01。 下标3(对应数据4)是内存最高位， 所以储存了数据最低位0x04
  fmt.Println("大端位模式, int32 转为 bytes", testBytes)

  convInt := binary.BigEndian.Uint32(testBytes)
  // todo: 输出 16909060, 实际上数据值都是一样的, 只是在内存中的储存方式不同
  fmt.Printf("大端序的字节转为int32, bytes 转为 int32: %d\n", convInt)
}

func LittleEndian() {
  // 小端序
  // 二进制: 0000 0001 - 0000 0010 - 0000 0011 - 0000 0100， 一共32bit， 4byte
  var testInt int32 = 0x01020304
  fmt.Printf("使用小端位: %d \n", testInt)

  var testBytes = make([]byte, 4)
  // todo: 输出 [4 3 2 1],  下标0(对应数据4）是内存最低位， 所以储存了数据最低位 0x04。 下标3(对应数据1)是内存最高位， 所以储存了数据最高位0x01
  binary.LittleEndian.PutUint32(testBytes, uint32(testInt))
  fmt.Println("小端位模式, int32 转为 bytes", testBytes)

  convInt := binary.LittleEndian.Uint32(testBytes)
  // todo: 输出 16909060, 实际上数据值都是一样的, 只是在内存中的储存方式不同
  fmt.Printf("小端序的字节转为int32, bytes 转为 int32: %d\n", convInt)
}

func main() {
  BigEndian()
  LittleEndian()
}



