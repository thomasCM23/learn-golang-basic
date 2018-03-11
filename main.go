package main

import (
	"fmt"

	"github.com/thomasCM23/learnGoLang/DLinkedList"
	"github.com/thomasCM23/learnGoLang/QueueNode"
	"github.com/thomasCM23/learnGoLang/StackNode"
)

func main() {
	stack := stackNode.NewStack(10, 34, "hello")
	fmt.Println(stack)
	/*stack.Push("Hello")
	stack.Push([]int{1, 2, 3, 4, 5})
	stack.Push(1)
	stack.Push(3.3234534)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Top())
	fmt.Printf("%T\n", stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	stack.Push(3.3234534)
	fmt.Printf("%T\n", stack.Pop())*/
	queue := queueNode.NewQueue(1, 3, 4, 6, "hello")
	fmt.Println(queue)
	/*queue.Enqueue(10)
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue)*/
	ddList := DLinkedList.NewDLinkList("hello", "world")
	ddList.Insert(3)
	ddList.InsertMulti("Thomas", "Cha")
	fmt.Println(ddList)
	ddList.RemoveAt(3)
	fmt.Println(ddList)
	ddList.RemoveAll()
	fmt.Println(ddList)
}
