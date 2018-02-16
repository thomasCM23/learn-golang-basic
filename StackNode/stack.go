package stackNode

import "fmt"

type nodeStack struct {
	TElement interface{}
	Previous *nodeStack
}

// Stack create new stack
type Stack struct {
	node nodeStack
	size int
}

// Push item onto stack
func (s *Stack) Push(item interface{}) {
	oldNode := s.node
	newNode := nodeStack{TElement: item, Previous: &oldNode}
	s.node = newNode
	s.size++
}

// Pop item off of stack
func (s *Stack) Pop() interface{} {
	if s.size > 0 {
		i := s.node.TElement
		if s.node.Previous != nil {
			s.node = *s.node.Previous
		}
		s.size--
		return i
	}
	return "Empty Stack Error"
}

// Top item in stack
func (s *Stack) Top() interface{} {
	return s.node.TElement
}

// Size get size of stack
func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) String() string {
	var items string
	current := s.node
	for i := 0; i < s.Size(); i++ {
		items += fmt.Sprintf("%v", current.TElement)
		if i < s.Size()-1 {
			items += ", "
		}
		current = *current.Previous
	}
	return "[" + items + "]"
}

// NewStack instantiates stack with items
// returns pointer to stack
func NewStack(items ...interface{}) *Stack {
	stack := new(Stack)
	for _, item := range items {
		stack.Push(item)
	}
	return stack
}
