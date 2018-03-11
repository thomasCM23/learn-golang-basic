package DLinkedList

type dListNode struct {
	TElement interface{}
	Prev     *dListNode
	Next     *dListNode
	Index    int64
}

// DoublyLinkList Create a doubly lint list
type DoublyLinkList struct {
	firstNode *dListNode
	lastNode  *dListNode
	size      int64
}

// Insert Item into list
func (d *DoublyLinkList) Insert(item interface{}) {
	newNode := new(dListNode)
	d.size++
	if d.firstNode.Next == nil && d.lastNode.Prev == nil {
		// first element to be inserted
		newNode.TElement = item
		newNode.Index = d.size - 1
	}
}

// NewDLinkList instantiates d link list
func NewDLinkList(items ...interface{}) *DoublyLinkList {

	dLList := new(DoublyLinkList)

	return dLList
}
