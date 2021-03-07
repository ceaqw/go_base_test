package doublelink

import (
	"fmt"
)

type MyNode struct {
	No int
	Info string
	Pre *MyNode
	Next *MyNode
}

//print by pre and next
func Print(head *MyNode)  {
	tmp := head.Next
	flag_ := head
	if tmp == nil {
		fmt.Println("当前链表为空")
		return
	}
	for tmp != nil {
		if tmp.Next == nil {
			flag_ = tmp
		}
		fmt.Printf("No: %d, Info: %s, Pre: %v, Next: %v\n", tmp.No, tmp.Info, tmp.Pre, tmp.Next)
		tmp = tmp.Next
	}
	for flag_ != nil {
		fmt.Printf("No: %d, Info: %s\n", flag_.No, flag_.Info)
		flag_ = flag_.Pre
	}
}

//create
func Insert(head *MyNode, newNode *MyNode)  {
	tmp := head
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = newNode
	newNode.Pre = tmp
}

//update
func Update(head *MyNode, no int, info string)  {
	tmp := head.Next
	for tmp != nil {
		if tmp.No == no {
			tmp.Info = info
			return
		}
	}
	panic("未找到该元素")
}

//delete
func Delete(head *MyNode, no int)  {
	tmp := head
	for tmp.Next != nil {
		if tmp.Next.No == no {
			tmp.Next = tmp.Next.Next
			if tmp.Next != nil {
				tmp.Next.Pre = tmp
			}
			return
		}
		tmp = tmp.Next
	}
	panic("未找到该元素")
}

//retrieve
func Retrieve(head *MyNode, no int) *MyNode {
	tmp := head.Next
	for tmp != nil {
		if tmp.No == no {
			return tmp
		}
		tmp = tmp.Next
	}
	panic("未找到该元素")
}

//func main() {
//	Head := &MyNode{Pre: nil, Next: nil}
//	A := MyNode{No: 1, Info: "cc", Pre: nil, Next: nil}
//	B := MyNode{No: 2, Info: "cb", Pre: nil, Next: nil}
//	Insert(Head, &A)
//	Insert(Head, &B)
//	Print(Head)
//	Update(Head, 1, "cv")
//	Print(Head)
//	Delete(Head, 2)
//	Print(Head)
//	fmt.Println(Retrieve(Head, 1))
//}
