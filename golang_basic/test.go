package main 

import (
  "fmt"
)

func main() {
  nodes := []int{1, 2, 3, 4, 5, 6}
  tmp := make([]int, len(nodes))
  // copy(tmp, nodes[2:])
  copy(tmp, nodes[:2])
  fmt.Println(nodes)
  fmt.Println(tmp)
}
