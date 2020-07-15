package main

import "fmt"

func QuicklySort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		// 取第一个当基准值
		splitData := arr[0]

		// 记录比基准值小的部分
		low := make([]int, 0, 0)
		// 记录比基准值大的部分
		high := make([]int, 0, 0)
		// 记录等于基准值的部分
		equl := make([]int, 0, 0)
		// 保存分离的数据
		equl = append(equl,splitData)

		// 从第二个开始遍历
		for i := 1; i < length; i++ {
			if arr[i] < splitData {
				low = append(low, arr[i])
			} else if arr[i] > splitData {
				high = append(high, arr[i])
			} else {
				equl = append(equl, arr[i])
			}
		}
		// 递归循环low和high部分
		low,high = QuicklySort(low),QuicklySort(high)
		myArr :=append(append(low,equl...),high...)
		return myArr
	}
}
func main() {
	arr := []int{1, 51, 2, 4, 2, 858, 24, 3, 515}
	fmt.Println("not sort", arr)
	fmt.Println(QuicklySort(arr))

}
