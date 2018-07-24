package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/thomasCM23/learnGoLang/DoublyLinkList"
	"github.com/thomasCM23/learnGoLang/Queue"
	"github.com/thomasCM23/learnGoLang/Stack"
)

func main() {
	stack := stack.NewStack(10, 34, "hello")
	fmt.Println(stack)
	stack.Push("Hello")
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
	fmt.Printf("%T\n", stack.Pop())
	queue := queue.NewQueue(1, 3, 4, 6, "hello")
	fmt.Println(queue)
	queue.Enqueue(10)
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue)
	ddList := doublylinklist.NewDLinkList("hello", "world")
	ddList.Insert(3)
	ddList.InsertMulti("Thomas", "Cha")
	fmt.Println(ddList)
	ddList.RemoveAt(3)
	fmt.Println(ddList)
	ddList.RemoveAll()
	fmt.Println(ddList)
}

//Concorrency Patters
// Function the returns a channel

func boring(msg string) <-chan string { // returns receive-only channel of strings
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // return the channel to the caller
}

// Multiplexing
func fanIn(inputs ...<-chan string) <-chan string {
	c := make(chan string)
	for _, input := range inputs {
		go func() {
			for {
				c <- <-input
			}
		}()
	}
	return c
}
