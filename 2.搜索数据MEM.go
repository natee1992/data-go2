package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

const N = 84331445

func main() {
	allStrs := make([]string, N, N)

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
		allStrs[i] = string(line)
		i++
	}
	fmt.Println("数据载入内存")
	time.Sleep(time.Second*5)
	for ;;{
		fmt.Println("请输入要查询的数据")
		var inputStr string
		fmt.Scanln(&inputStr)

		startTime := time.Now()
		for j := 0; j < N-1; j++ {
			if strings.Contains(allStrs[j], inputStr) {
				fmt.Println(allStrs[j])
			}
		}
		fmt.Println(time.Since(startTime))
	}

}

/*
1486790834----wufeifei521521
1350690572----wufeifei520
820523474----wufeifei132.
450010707----wufeifei515
1262592058----wufeifei1224
*/
