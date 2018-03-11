package DLinkedList

import "fmt"

type dListNode struct {
	TElement interface{}
	Prev     *dListNode
	Next     *dListNode
	Index    int
}

// DoublyLinkList Create a doubly lint list
type DoublyLinkList struct {
	firstNode *dListNode
	lastNode  *dListNode
	size      int
}

// Insert Item into list
func (d *DoublyLinkList) Insert(item interface{}) {
	d.size++
	newNode := new(dListNode)
	newNode.TElement = item
	newNode.Index = d.size - 1

	oldLastReal := d.lastNode.Prev
	newNode.Next = d.lastNode
	newNode.Prev = oldLastReal
	oldLastReal.Next = newNode
	d.lastNode.Prev = newNode
}

// InsertMulti Item into list
func (d *DoublyLinkList) InsertMulti(items ...interface{}) {
	for _, item := range items {
		d.Insert(item)
	}
}

// GetFirst element in list
func (d *DoublyLinkList) GetFirst() interface{} {
	return d.firstNode.Next.TElement
}

// GetLast element in list
func (d *DoublyLinkList) GetLast() interface{} {
	return d.lastNode.Prev.TElement
}

// Count number of itme in list
func (d *DoublyLinkList) Count() int {
	return d.size
}

// RemoveAt index specified
func (d *DoublyLinkList) RemoveAt(index int) interface{} {
	retVal := d.firstNode.Next
	if index > d.Count()-1 || index < 0 {
		return "Index Out Of Range Error"
	}
	for index != retVal.Index {
		retVal = retVal.Next
	}
	prev := retVal.Prev
	next := retVal.Next
	prev.Next = next
	next.Prev = prev
	retVal = nil
	d.size--
	return nil
}

// RemoveAll from list
func (d *DoublyLinkList) RemoveAll() {
	//TODO: fix, this will create dangling pointer
	d.firstNode.Next = d.lastNode
	d.lastNode.Prev = d.firstNode
	d.size = 0
}

// GetAtIndex return item at specified index
func (d *DoublyLinkList) GetAtIndex(index int) interface{} {
	retVal := d.firstNode.Next
	if index > d.size-1 || index < 0 {
		return "Index Out Of Range Error"
	}
	for index != retVal.Index {
		retVal = retVal.Next
	}
	return retVal.TElement
}

// NewDLinkList instantiates d link list
func NewDLinkList(items ...interface{}) *DoublyLinkList {

	dLList := new(DoublyLinkList)
	setUpHeadTail(dLList)
	for _, item := range items {
		dLList.Insert(item)
	}
	return dLList
}

func setUpHeadTail(d *DoublyLinkList) *DoublyLinkList {
	headerNode := new(dListNode)
	headerNode.Index = -1
	headerNode.TElement = nil
	tailNode := new(dListNode)
	tailNode.Index = -1
	tailNode.TElement = nil
	headerNode.Next = tailNode
	headerNode.Prev = nil
	tailNode.Prev = headerNode
	tailNode.Next = nil
	d.firstNode = headerNode
	d.lastNode = tailNode
	return d
}

func (d *DoublyLinkList) String() string {
	var items string
	current := d.firstNode.Next
	for i := 0; i < d.Count(); i++ {
		items += fmt.Sprintf("%v", current.TElement)
		if i < d.Count()-1 {
			items += ", "
		}
		current = current.Next
	}
	return "[" + items + "]"
}
