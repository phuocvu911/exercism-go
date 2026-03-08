package linkedlist
import "errors"
// Define List and Node types here.
type Node struct{
    Value any
    next *Node
    prev *Node
}
type List struct{
    head *Node
    tail *Node
    size int
}

func NewList(elements ...any) *List {
	l := &List{}
    for _,v:= range elements{
        l.Push(v)
    }
    return l
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v any) {
	node:= &Node{Value: v}
    if l.size == 0{
        l.head = node
        l.tail = node
    } else{
        node.next = l.head // make the current head become inbetweens first then set head to that node at the last
        l.head.prev = node
        l.head = node       
    }
    l.size++
}

func (l *List) Push(v any) {
    node := &Node{Value: v}
	if l.tail == nil{
        l.head = node
        l.tail = node
    } else {
        node.prev = l.tail
        l.tail.next = node
        l.tail = node
    }
    l.size++
}

func (l *List) Shift() (any, error) {
	if l.size == 0 {
        return nil, errors.New("list is empty")
    }
    res := l.head.Value
    if l.tail == l.head{//if list have only 1 node
        l.tail = nil
        l.head = nil
    } else{
        l.head = l.head.next
        l.head.prev = nil
    }   
    l.size--
    return res, nil
}

func (l *List) Pop() (any, error) {
	if l.size == 0{ //if list is empty
        return nil, errors.New("list is empty")
    }
    res:= l.tail.Value
    if l.tail == l.head{//if list have only 1 node
        l.tail = nil
        l.head = nil
    } else{
        l.tail = l.tail.prev // normal case
    	l.tail.next = nil
    }   
    l.size--
    return res, nil
}

func (l *List) Reverse() {
	if l.size == 0{
        return
    }
    node:= l.head
    for node!=nil{
        node.next, node.prev = node.prev, node.next
        node = node.prev
    }
    l.head, l.tail = l.tail, l.head
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

func (l *List) Count() int {
	return l.size
}

// Delete removes the first node in a list with a given value.
// Returns true if a node was removed.
func (ll *List) Delete(v any) bool {
	if ll.size == 0{ //if list is empty
        return false
    }
    current:= ll.head
    for current != nil{
        if current.Value == v{
            switch{
                case ll.head ==ll.tail: 
                	ll.Pop()
                case ll.head == current:
                	ll.Shift()
                case current == ll.tail:
                	ll.Pop()
                default:
                	current.prev.next = current.next
                	current.next.prev = current.prev
                	ll.size--
            }       
            return true
        }
        current = current.next
    }
    return false
}
