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
	QQuser int
	password string
}

const N = 84331445

func QuickSortQQ(arr []QQ) []QQ {
	length := len(arr)
	if length < 1 {
		return arr
	} else {
		splitData := arr[0].QQuser //第一个数字
		low := make([]QQ, 0, 0)
		high := make([]QQ, 0, 0)
		mid := make([]QQ, 0, 0)
		mid = append(mid, arr[0]) // 保存分离的数据
		for i := 1; i < length; i++ {
			if arr[i].QQuser < splitData {
				low = append(low, arr[i])
			} else if arr[i].QQuser > splitData {
				high = append(high, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
		low, high = QuickSortQQ(low), QuickSortQQ(high) //递归循环
		myArr := append(append(low, mid...), high...)
		return myArr
	}
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

//1486790834----wufeifei521521
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
		if len(lineArr) == 4 {
			allData[i].QQuser,_ = strconv.Atoi(lineArr[0])
			allData[i].password = lineArr[1]
		}
		i++

	}
	fmt.Println(i, "内存载入完成")
	allData = QuickSortQQ(allData)
	for ;;{
		fmt.Println("请输入要查询的用户名")
		var inputStr int
		fmt.Scanf("%d",&inputStr)
		startTime := time.Now()
		for i:=0;i<N;i++{
			if allData[i].QQuser ==inputStr{
				fmt.Println(allData[i])
			}
		}
		fmt.Println("耗时：",time.Since(startTime))
	}
}
