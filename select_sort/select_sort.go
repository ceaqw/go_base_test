package main

import (
	"fmt"
	"math/rand"
)

func SelectSort(arr []int, reverse bool) {
	arrLen := len(arr)
	for i:=0; i<arrLen-1; i++ {
		tmp := arr[i]
		index := i
		for j:=i+1; j<arrLen; j++ {
			if reverse {
				if tmp < arr[j] {
					tmp = arr[j]
					index = j
				}
			} else {
				if tmp > arr[j] {
					tmp = arr[j]
					index = j
				}
			}
		}
		arr[index] = arr[i]
		arr[i] = tmp
	}
}

func main() {
	Arr := make([]int, 10)
	for i:=0; i<10; i++ {
		Arr[i] = rand.Intn(20)
	}
	fmt.Println(Arr)
	SelectSort(Arr, false)
	fmt.Println(Arr)
}
