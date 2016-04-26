package main 

import (
  "fmt"
)


// ================ struct start

type Node struct {
  Value int
}

//stack is a bisic LIFO stack that resizes as needed
type Stack struct {
  nodes []*Node
  count int
}

//queue is a basic FIFO queue based on a circular  list that resizes as needed
type Queue struct {
  nodes []*Node
  size int
  head int
  tail int
  count int
}

// ================= struct end 

//make the value to be string type
func (n *Node) String() string {
  return fmt.Sprint(n.Value)
}

//return a new Stack
func NewStack() *Stack {
  return &Stack{}
}


//push adds a node to the stack
func (s *Stack) Push(n *Node) {
  s.nodes = append(s.nodes[:s.count], n)
  s.count++
}

//pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() *Node {
  if s.count == 0 {
    return nil
  }
  s.count-- 
  return s.nodes[s.count]   //s.count is the last insert node`
}


//newQueue returns a new queue with the given initial size.
func NewQueue(size int) *Queue {
  return  &Queue{
            nodes: make([]*Node, size),
            size: size
  }
}

// Push adds a node to the queue.
/**
这是一个可扩展的环形结构，一旦添加进去的队列元素超过了第一个环，就会把第一个环的数据
保存到成整个slice的其中一段， 这一段的长度就等于是原来的那个满了的环的长度，然后又在
满了的环的下一个分元素开始, 位置 （环的长度+1）这个slice的key开始再开始新的环，然后再
往新的环形去添加队列的数据
**/
func (q *Queue) Push(n *Node) {
	if q.head == q.tail && q.count > 0 {
		nodes := make([]*Node, len(q.nodes)+q.size)
		copy(nodes, q.nodes[q.head:])
		copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
		q.head = 0
		q.tail = len(q.nodes)
		q.nodes = nodes
	}
	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes)
	q.count++
}


// Pop removes and returns a node from the queue in first to last order.
func (q *Queue) Pop() *Node {
	if q.count == 0 {
		return nil
	}
	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--          
	return node
}
 
func main() {
	s := NewStack()
	s.Push(&Node{1})
	s.Push(&Node{2})
	s.Push(&Node{3})
	fmt.Println(s.Pop(), s.Pop(), s.Pop())

	q := NewQueue(1)
	q.Push(&Node{4})
	q.Push(&Node{5})
	q.Push(&Node{6})
	fmt.Println(q.Pop(), q.Pop(), q.Pop())
}



