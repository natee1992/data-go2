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

func getLines() {
	path := "/Users/natee/go-data/yincheng/day3/uuu9.com.sql"
	sqlFile, _ := os.Open(path)    //打开文件
	i := 0                         //统计行数
	br := bufio.NewReader(sqlFile) //读取文件对象
	for {
		_, _, end := br.ReadLine()
		if end == io.EOF {
			break
		}
		i++

	}
	fmt.Println(i)
}

func loadFile() {

}

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
