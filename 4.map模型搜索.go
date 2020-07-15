package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

/*

Map 高效哈希表
*/

const N = 84331445

//1486790834----wufeifei521521
func main() {
	allStrs := make(map[string]string)

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
		if len(lineArray) < 2 {
			fmt.Println(lineArray)
		} else {
			//fmt.Println(lineArray[0])
			allStrs[lineArray[0]] = lineArray[1]
			i++
		}

	}
	fmt.Println("数据载入内存")
	time.Sleep(time.Second * 5)
	for ; ; {
		fmt.Println("请输入要查询的数据")
		var QQ string
		fmt.Scanln(&QQ)

		startTime := time.Now()
		v, ok := allStrs[QQ]
		if ok {
			fmt.Println(v)
		}
		fmt.Println(time.Since(startTime))
	}
}
