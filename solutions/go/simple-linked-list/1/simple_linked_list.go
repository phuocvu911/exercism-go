package simplelinkedlist
import "errors"
// Define the List and Element types here.
type Element struct{
    data int
    next *Element
}

type List struct{
    head *Element
    tail *Element
    size int
}

func New(elements []int) *List {
	l:= &List{}
    for _, v:= range elements{
        l.Push(v)
    }
    return l
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(element int) {
	node:= &Element{data: element}
    if l.size == 0{
        l.head = node
        l.tail = node
    } else {
        l.tail.next = node
        l.tail = node
    }
    l.size++
}

func (l *List) Pop() (int, error) {
	if l.size == 0{
        return 0, errors.New("list is empty")
    }
    res:= l.tail.data
    if l.head == l.tail{
        l.head = nil
        l.tail = nil
    } else{
        cur:= l.head
        for cur!= nil{
            if cur.next == l.tail{
                l.tail = cur
                l.tail.next = nil
                break
            }
            cur = cur.next
        }
    }
    l.size--
    return res, nil
}

func (l *List) Array() []int {
    res:= []int{}
    cur := l.head
    for cur != nil{
        res = append(res, cur.data)
        cur = cur.next
    }
    return res
}

func (l *List) Reverse() *List {
	if l.size == 0 || l.size == 1{
        return l
    }
    slice:=l.Array()
    reverse:= []int{}
    for i:= l.size -1; i>=0; i--{
        reverse = append(reverse, slice[i])
    }
    return New(reverse)
}
