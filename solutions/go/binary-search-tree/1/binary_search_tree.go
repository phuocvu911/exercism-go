package binarysearchtree

type BinarySearchTree struct {
	left  *BinarySearchTree
	data  int
	right *BinarySearchTree
}

// NewBst creates and returns a new BinarySearchTree.
func NewBst(i int) *BinarySearchTree {
	return &BinarySearchTree{data: i}
}

func (bst *BinarySearchTree) Insert(i int) {
	if i<= bst.data{
        if bst.left == nil{ //check if we have empty slot to write i
            bst.left = NewBst(i)
        } else{
            bst.left.Insert(i) //if we dont, we go down 1 level, now bst.left is become head
        }       
    } else{
        if bst.right == nil{
            bst.right = NewBst(i)
        } else{
            bst.right.Insert(i)
        }
    }
}

// SortedData returns the ordered contents of BinarySearchTree as an []int.
// The values are in increasing order starting with the lowest int value.
// A BinarySearchTree that has the numbers [1,3,7,5] added will return the
// []int [1,3,5,7].
func (bst *BinarySearchTree) SortedData() []int {
	res := []int{}
    var order func(node *BinarySearchTree)
    order = func(node *BinarySearchTree){
        if node == nil{
            return
        }
        order(node.left)
        res= append(res, node.data)
        order(node.right)
    }
    order(bst)
    return res
}
