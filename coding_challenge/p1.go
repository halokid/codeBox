package main

import (
  "log"
  "math"
  "strings"
)

func solution(digits string, num string) int {
  totalSec := 2
  tmpIndex := 0     // todo: the tmp index of char loop
  for i, char := range num {
    log.Println("char --->>>", string(char))
    charIndex := strings.Index(digits, string(char))
    // todo: the first time is press the key, dont need to add abs index
    if i != 0 {
      tmpSec := int(math.Abs(float64(charIndex - tmpIndex)))
      log.Println("tmpSec --->>>", tmpSec)
      totalSec += tmpSec
    }
    tmpIndex = charIndex
  }

  return totalSec
}

func main() {
  res := solution("0123456789", "210")
  log.Println("res -->>>", res)
}
