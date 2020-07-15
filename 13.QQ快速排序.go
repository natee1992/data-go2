package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const N = 84331445
type QQ struct {
	QQuser int
	password string
}
func sortForMerge(arr []QQ, left int, right int) {
	for i := left; i <= right; i++ {
		temp := arr[i] // 备份数据
		var j int
		for j = i; j > left && arr[j-1].QQuser > temp.QQuser; j-- {
			arr[j] = arr[j-1] // 数据往后移动
		}
		arr[j] = temp // 插入
	}
}


func swap(arr [] QQ, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// 递归快速排序
func OpQuickSortX(arr []QQ, left int, right int) {
	if right-left < 15{ // 数组剩下3个数字，直接插入排序
		sortForMerge(arr,left,right)
	} else {
		// 随机找一个数字,放在第一个位置
		swap(arr, left, rand.Int()%(right-left+1)+1)
		vdata := arr[left].QQuser //坐标数组，比我小，左边，比我大右边
		lt := left         // arr[left+1,lt] < vdata
		gt := right + 1    //arr[gt, r] > vdata
		i := left + 1      // arr[lt+1,...i] == vdata
		for i < gt {
			if arr[i].QQuser < vdata {
				swap(arr, i, lt+1) //移动到小于的地方
				lt++               //前进循环
				i++
			} else if arr[i].QQuser > vdata {
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
func QuickSortPlus(arr []QQ) {
	OpQuickSortX(arr, 0, len(arr)-1)
}

func binSearchQQ(arr []QQ, data int) int {
	left := 0
	right := len(arr) - 1 //最下面最上面
	for left < right {
		mid := (left + right) / 2
		if arr[mid].QQuser > data {
			right = mid - 1
		} else if arr[mid].QQuser < data {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1

}

func main() {
	allData := make([]QQ, N, N)

	path := "/Users/natee/go-data/yincheng/day3/QQ.txt"
	sqlFile, _ := os.Open(path)    //打开文件
	i := 0                         //统计行数
	br := bufio.NewReader(sqlFile) //读取文件对象
	for {
		line, _, end := br.ReadLine()
		if end == io.EOF {
			break
		}
		linestr := string(line)
		lineArr := strings.Split(linestr, "----")
		if len(lineArr) == 2 {
			allData[i].QQuser,_ = strconv.Atoi(lineArr[0])
			allData[i].password = lineArr[1]
		}
		i++

	}
	fmt.Println(i, "内存载入完成")
	QuickSortPlus(allData)

	for ;;{
		fmt.Println("请输入要查询的用户名")
		var inputStr int
		fmt.Scanf("%d",&inputStr)
		startTime := time.Now()
		index := binSearchQQ(allData,inputStr)
		if index != -1{
			fmt.Println("no value")
		}else {
			fmt.Println(allData[index])
		}
		fmt.Println("耗时：",time.Since(startTime))
	}
}
