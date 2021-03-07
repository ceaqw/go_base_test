package josephu

import (
	"fmt"
)

//错解
//func JosePhu(arr [41]int, arrLen int, m int) {
//	//fmt.Println(arr)
//	lastNum := m
//	if m > arrLen {
//		lastNum = m%arrLen
//	}
//	if lastNum == 0 {
//		lastNum = arrLen
//	}
//	fmt.Println(arr[lastNum-1])
//	if arrLen == 1 {
//		return
//	}
//	for i:=lastNum-1; i<arrLen-1; i++ {
//		arr[i] = arr[i+1]
//	}
//	JosePhu(arr, arrLen-1, m)
//}

type JoseNode struct {
	Num int
	Next *JoseNode
}

func MakeCircleLink(head *JoseNode, len int) {
	tmp := head
	for i:=2; i<len+2; i++ {
		tmp.Next = &JoseNode{Num: i, Next: nil}
		if i == len+1 {
			tmp.Next = head
			break
		}
		tmp = tmp.Next
	}
}

func JosePhu(head *JoseNode, m int, arrLen int)  {
	tmp := head
	flag_ := head
	for i:=1; i<m; i++ {
		flag_ = tmp
		tmp = tmp.Next
	}
	fmt.Println(tmp.Num)
	if arrLen == 1{
		return
	}
	flag_.Next = tmp.Next
	JosePhu(flag_.Next, m, arrLen-1)
}

func Print_(head *JoseNode)  {
	tmp := head
	for {
		fmt.Println(tmp)
		tmp = tmp.Next
		if tmp == head {
			return
		}
	}
}

//func main() {
//	//arr := [41]int{}
//	//for i:=0; i<41; i++{
//	//	arr[i] = i+1
//	//}
//	//JosePhu(arr, 41, 3)
//	Head := &JoseNode{Num: 1, Next: nil}
//	MakeCircleLink(Head, 41)
//	JosePhu(Head, 3, 41)
//	//fmt.Print(Head.Next)
//	//Print_(Head)
//}

