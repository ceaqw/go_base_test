package single_link

// binary_tree's package

import (
	"fmt"
)

type MyNode struct {
	No int
	Info string
	NextNode *MyNode
}

// print node info
func PrintLink (head *MyNode)  {
	tmp := head.NextNode

	if tmp == nil {
		fmt.Println("This is an empty single link!")
		return
	}

	fmt.Println("Print result:")
	for {
		fmt.Printf("Num: %d, Info: %s\n", tmp.No, tmp.Info)
		tmp = tmp.NextNode
		if tmp == nil {
			return
		}
	}
}

// create
func InsertByHead(head *MyNode, newNode *MyNode) {
	tmp := head
	newNode.NextNode = tmp.NextNode
	tmp.NextNode = newNode
}

// update
func Update(head *MyNode, no int, info string)  {
	tmp := head.NextNode
	for tmp != nil {
		if tmp.No == no {
			tmp.Info = info
			return
		}
		tmp = tmp.NextNode
	}
	panic("NoExisted!")
}

// delete
func Delete(head *MyNode, delNo int)  {
	tmp := head
	// first node del
	if tmp.No == delNo {
		head = head.NextNode
		return
	}

	for tmp.NextNode != nil {
		if tmp.NextNode.No == delNo {
			fmt.Println(111)
			tmp.NextNode = tmp.NextNode.NextNode
			fmt.Println("deleted!")
			return
		}
		tmp = tmp.NextNode
	}
	panic("NoExisted!")
}

// retrieve
func Retrieve(head *MyNode, no int) *MyNode {
	tmp := head.NextNode
	for tmp != nil {
		if tmp.No == no {
			return tmp
		}
		tmp = tmp.NextNode
	}
	return nil
}

//func main() {
//	Head := &MyNode{}
//	A := MyNode{No: 1, Info: "cc", NextNode: nil}
//	B := MyNode{No: 2, Info: "cv", NextNode: nil}
//	C := MyNode{No: 3, Info: "ct", NextNode: nil}
//	InsertByHead(Head, &A)
//	InsertByHead(Head, &B)
//	InsertByHead(Head, &C)
//	PrintLink(Head)
//	Update(Head, 2, "cm")
//	PrintLink(Head)
//	Delete(Head, 3)
//	PrintLink(Head)
//
//	fmt.Println(Retrieve(Head, 2).Info)
//}
