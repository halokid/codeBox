package main 

import (
  "fmt"
)

// 初始化hash计算需要的基础map table
func initCryptTable() {
  var seed, index1, index2 uint64 0x00100001, 0, 0
  i := 0
  for index1 = 0; index1 < 0x100; index1 += 1 {
    seed = (seed * 125 + 3) % 0x2aaaaab
    temp1 := (seed & 0xffff) << 0x10
    seed = (seed * 125 + 3) % 0x2aaaaab
    temp2 := seed & 0xffff
    cryptTable[index2] = temp1 | temp2
    i += 1
  }
}

// hash, 以及相关校验hash值
func HashKey(lpszString string, dwHashType int) uint64 {
    i, ch := 0, 0
    var seed1, seed2 uint64 = 0x7FED7FED, 0xEEEEEEEE
    var key uint8
    strLen := len(lpszString)
    for i < strLen {
        key = lpszString[i]
        ch = int(toUpper(rune(key)))
        i += 1
        seed1 = cryptTable[(dwHashType<<8)+ch] ^ (seed1 + seed2)
        seed2 = uint64(ch) + seed1 + seed2 + (seed2 << 5) + 3
    }
    return uint64(seed1)
}


func main() {
  
}