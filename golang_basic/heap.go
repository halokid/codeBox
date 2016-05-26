package main 

import (
  "container/heap"
  "fmt"
)

//an IntHeap is a min-heap of ints
type IntHeap []int 

func (h IntHeap) Len() int                    { return len(h)}
func (h IntHeap) Less(i, j int) bool          { return h[i] < h[j]}
func (h IntHeap) Swap(i, j int)               { h[i], h[j] = h[j], h[i]}


func (h *IntHeap) Push(x interface{}) {
  //push and pop use pointer receviers because they modify the slice's length,
  //not just its contents.
  *h = append(*h, x.(int))
} 


func (h *IntHeap) Pop() interface{} {
  old := *h 
  n := len(old)
  
  x := old[n-1]
  *h = old[0 : n-1]
  return x
}

func main() {
  h := &IntHeap{ 100, 16, 4, 8, 70, 2, 36, 22, 5, 12}
  
  fmt.Println("Heap:")
  heap.Init(h)
  
  fmt.Println("Push(h, 3),然后输出堆看看")
  
  for i := 0; i < 40; i++ {
    heap.Push(h , i)
    heap.Pop(h)
  }
  
  for h.Len() > 0 {
    fmt.Printf("%d ", heap.Pop(h))
  }
  
}