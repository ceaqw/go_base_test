package bubble_sort

func swap(arr []int, index int)  {
	tmp := arr[index]
	arr[index] = arr[index+1]
	arr[index+1] = tmp
}

func BubbleSort(arr []int, reverse bool) {
	arrLen := len(arr)
	for i:=1; i<arrLen-1; i++ {
		for j:=0; j<arrLen-i; j++ {
			if reverse {
				if arr[j] < arr[j+1] {
					swap(arr, j)
				}
			} else {
				if arr[j] > arr[j+1] {
					swap(arr, j)
				}
			}
		}
	}
}
//
//func main() {
//	Arr := make([]int, 10)
//	for i:=0; i<10; i++ {
//		Arr[i] = rand.Intn(20)
//	}
//	fmt.Println(Arr)
//	BubbleSort(Arr, true)
//	fmt.Println(Arr)
//}