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
		}

	}
	fmt.Println("数据载入内存")
	time.Sleep(time.Second * 5)
	for ; ; {
		fmt.Println("请输入要查询的数据")
		var QQ int
		fmt.Scanf("%d",&QQ)

		startTime := time.Now()
		for j := 0; j < N-1; j++ {
			if allStrs[i].QQUser ==QQ{
				fmt.Println(i,allStrs[i].QQUser,allStrs)
			}
		}
		fmt.Println(time.Since(startTime))
	}
}
