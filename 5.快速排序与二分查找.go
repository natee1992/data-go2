package main

import "fmt"

func QuickSort(arr []int) []int {
	length :=len(arr)
	if length <1{
		return arr
	}else {
		splitData := arr[0] //第一个数字
		low := make([]int, 0, 0)
		high := make([]int, 0, 0)
		mid := make([]int, 0, 0)
		mid =append(mid, splitData)
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

func bin_search(arr[]int, data int)int  {
	left := 0
	right := len(arr) -1 //最下面最上面
	for left<right{
		mid := (left+right)/2
		if arr[mid] > data{
			right = mid-1
		}else if arr[mid] < data{
			left = mid +1
		}else {
			return mid
		}
	}
	return -1

}

func main()  {
	arr :=[]int{1,19,4,8,3,5,4,6,19,0}
	fmt.Println(QuickSort(arr))
	fmt.Println(bin_search(arr,55))
}
