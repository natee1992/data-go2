package main

import "fmt"

// 找到第一个等于3
// 找到最后一个等于3
// 找到第一个大于等于2
// 找到最后一个小于7的数据

func main() {
	arr := []int{1, 2, 3, 3, 3, 3, 3, 3, 3, 4, 5, 5, 5, 5, 6, 6, 6, 6, 7, 9, 9, 10}

	fmt.Println(binserachC(arr, 3))
}

func binseracha(arr []int, data int) int {
	low := 0
	high := len(arr) - 1
	index := -1

	for low < high {
		mid := (low + high) / 2

		if arr[mid] > data {
			high = mid - 1
		} else if arr[mid] < data {
			low = mid + 1
		} else {
			if mid == 0 || arr[mid-1] != data {
				index = mid
				//fmt.Println(mid)
				break
			} else {
				high = mid - 1
			}
		}
	}
	return index
}

// 找到第一个大于等于3
func binserachC(arr []int, data int) int {
	low := 0
	high := len(arr) - 1
	index := -1

	for low < high {
		mid := (low + high) / 2

		if arr[mid] > data {
			high = mid - 1
		} else if arr[mid] < data {
			low = mid + 1
		} else {
			if mid == 0 || arr[mid-1] != data {
				index = mid
				//fmt.Println(mid)
				break
			} else {
				high = mid - 1
			}
		}
	}
	return index
}