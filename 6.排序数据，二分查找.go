package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type QQ struct {
	QQUser int
	QQpass string
}

const N = 84331445
func QuickSortStruct(arr []QQ) []QQ {
	length :=len(arr)
	if length <1{
		return arr
	}else {
		splitData := arr[0].QQUser //第一个数字
		low := make([]QQ, 0, 0)
		high := make([]QQ, 0, 0)
		mid := make([]QQ, 0, 0)
		mid =append(mid, arr[0])  // 保存分离的数据
		for i := 1; i < length; i++ {
			if arr[i].QQUser < splitData {
				low = append(low, arr[i])
			} else if arr[i].QQUser > splitData {
				high = append(high, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
		low, high = QuickSortStruct(low), QuickSortStruct(high) //递归循环
		myArr := append(append(low, mid...), high...)
		return myArr
	}
}

func binSearchStruct(arr[]QQ, data int)int  {
	left := 0
	right := len(arr) -1 //最下面最上面
	for left<right{
		mid := (left+right)/2
		if arr[mid].QQUser > data{
			right = mid-1
		}else if arr[mid].QQUser < data{
			left = mid +1
		}else {
			return mid
		}
	}
	return -1

}


//1486790834----wufeifei521521
func main() {
	allStrs := make([]QQ, N, N)

	path, _ := os.Open("/Users/natee/go-data/yincheng/day2晚上/QQ.txt")
	defer path.Close()
	br := bufio.NewReader(path)
	i := 0
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		//fmt.Println(string(line))
		lineArray := strings.Split(string(line), "----")
		if len(lineArray) < 2{
			fmt.Println(lineArray)
		}else {
			allStrs[i].QQpass = lineArray[1]
			allStrs[i].QQUser, _ = strconv.Atoi(lineArray[0])
			i++
			if i == 1000000{
				break
			}
		}

	}
	fmt.Println("数据载入内存")
	time.Sleep(time.Second * 5)
	for ; ; {
		fmt.Println("请输入要查询的数据")
		var QQ int
		fmt.Scanf("%d",&QQ)

		startTime := time.Now()
		// 顺序，二分查找
		fmt.Println("开始排序")
		QQArr := QuickSortStruct(allStrs)
		fmt.Println("排序结束")
		fmt.Println(binSearchStruct(QQArr,QQ))
		fmt.Println(time.Since(startTime))
	}
}


