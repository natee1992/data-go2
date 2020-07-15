package main

import (
	"fmt"
	"math/rand"
)

func sortForMerge(arr []int, left int, right int) {
	for i := left; i <= right; i++ {
		temp := arr[i] // 备份数据
		var j int
		for j = i; j > left && arr[j-1] > temp; j-- {
			arr[j] = arr[j-1] // 数据往后移动
		}
		arr[j] = temp // 插入
	}
}


func swap(arr [] int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// 递归快速排序
func OpQuickSortX(arr []int, left int, right int) {
	if right-left < 3{ // 数组剩下3个数字，直接插入排序
		sortForMerge(arr,left,right)
	} else {
		// 随机找一个数字,放在第一个位置
		swap(arr, left, rand.Int()%(right-left+1)+1)
		vdata := arr[left] //坐标数组，比我小，左边，比我大右边
		lt := left         // arr[left+1,lt] < vdata
		gt := right + 1    //arr[gt, r] > vdata
		i := left + 1      // arr[lt+1,...i] == vdata
		for i < gt {
			if arr[i] < vdata {
				swap(arr, i, lt+1) //移动到小于的地方
				lt++               //前进循环
				i++
			} else if arr[i] > vdata {
				swap(arr, i, gt-1) //移动到大于的地方
				gt--
			} else {
				i++
			}
		}
		swap(arr, left, lt)           //交换头部位置
		OpQuickSortX(arr, left, lt-1) //递归处理小于那一段
		OpQuickSortX(arr, gt, right)  // 递归处理大于那一段
	}
}

// 快速排序核心
func QuickSortPlus(arr []int) {
	OpQuickSortX(arr, 0, len(arr)-1)
}

func main() {
	arr := []int{1, 51, 2, 4, 2, 858, 24, 3, 515}
	fmt.Println("not sort", arr)
	QuickSortPlus(arr)
	fmt.Println("not sort", arr)
}
