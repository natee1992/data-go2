package main

import "fmt"

/*只能用于有序*/

func QuickSort(arr []int) []int {
	length := len(arr)
	if length < 1 {
		return arr
	} else {
		splitData := arr[0] //第一个数字
		low := make([]int, 0, 0)  //存储比我小的
		high := make([]int, 0, 0) //存储比我大的
		mid := make([]int, 0, 0)  //存储与我相等的
		mid = append(mid, splitData) //加入第一个相等的
		for i := 1; i < length; i++ {
			if arr[i] < splitData {
				low = append(low, arr[i])
			} else if arr[i] > splitData {
				high = append(high, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
		low, high = QuickSort(low), QuickSort(high) //递归循环
		myArr := append(append(low, mid...), high...)
		return myArr
	}
}

func binSearch(arr []int, data int) int {
	low := 0
	high := len(arr) - 1
	for low < high { //循环的终止条件
		mid := (low + high) / 2
		if arr[mid] > data {
			high = mid - 1

		} else if arr[mid] < data {
			low = mid + 1
		}else {
			return mid
		}

	}
	return -1
}

func main() {
	arr := []int{3, 8, 84, 784, 22, 255}
	fmt.Println(QuickSort(arr))
	index := binSearch(QuickSort(arr),22)
	fmt.Println(index)
}
