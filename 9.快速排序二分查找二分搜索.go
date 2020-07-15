package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type uuu9 struct {
	user  string
	md5   string
	email string
	pwd   string
}

const N = 18333811

func QuickSortU(arr []uuu9) []uuu9 {
	length := len(arr)
	if length < 1 {
		return arr
	} else {
		splitData := arr[0].user //第一个数字
		low := make([]uuu9, 0, 0)
		high := make([]uuu9, 0, 0)
		mid := make([]uuu9, 0, 0)
		mid = append(mid, arr[0]) // 保存分离的数据
		for i := 1; i < length; i++ {
			if arr[i].user < splitData {
				low = append(low, arr[i])
			} else if arr[i].user > splitData {
				high = append(high, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
		low, high = QuickSortU(low), QuickSortU(high) //递归循环
		myArr := append(append(low, mid...), high...)
		return myArr
	}
}

func binSearchU(arr []uuu9, data string) int {
	left := 0
	right := len(arr) - 1 //最下面最上面
	for left < right {
		mid := (left + right) / 2
		if arr[mid].user > data {
			right = mid - 1
		} else if arr[mid].user < data {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1

}

//1486790834----wufeifei521521
func main() {
	allData := make([]uuu9, N, N)

	path := "/Users/natee/go-data/yincheng/day3/uuu9.com.sql"
	sqlFile, _ := os.Open(path)    //打开文件
	i := 0                         //统计行数
	br := bufio.NewReader(sqlFile) //读取文件对象
	for {
		line, _, end := br.ReadLine()
		if end == io.EOF {
			break
		}
		linestr := string(line)
		lineArr := strings.Split(linestr, " | ")
		if len(lineArr) == 4 {
			allData[i].user = lineArr[0]
			allData[i].md5 = lineArr[1]
			allData[i].email = lineArr[2]
			allData[i].pwd = lineArr[3]
		}
		i++

	}
	fmt.Println(i, "内存载入完成")
	allData = QuickSortU(allData)
	for ;;{
		fmt.Println("请输入要查询的用户名")
		var inputStr string
		fmt.Scanln(&inputStr)
		startTime := time.Now()
		for i:=0;i<N;i++{
			if allData[i].user ==inputStr{
				fmt.Println(allData[i])
			}
		}
		fmt.Println("耗时：",time.Since(startTime))
	}
}
