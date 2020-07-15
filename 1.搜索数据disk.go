package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main()  {
	startTime := time.Now()
	path,_ := os.Open("/Users/natee/go-data/yincheng/day2晚上/163.txt")
	defer path.Close()
	br := bufio.NewReader(path)
	i := 0
	for {
		line,_,err := br.ReadLine()
		if err == io.EOF{
			break
		}
		//fmt.Println(string(line))
		lineStr := string(line)
		if strings.Contains(lineStr,"zhangbin2005"){
			fmt.Println(lineStr)
		}
		i++
	}
	fmt.Println(time.Since(startTime))
	fmt.Println(i)
}

//
//func main()  {
//	startTime := time.Now()
//	path,_ := os.Open("/Users/natee/go-data/yincheng/day2晚上/QQ.txt")
//	defer path.Close()
//	scanner := bufio.NewScanner(path)
//	for scanner.Scan(){
//		line := scanner.Text()
//		if strings.Contains(line,"zhangbin2005"){
//			fmt.Println(line)
//		}
//	}
//	fmt.Println(time.Since(startTime))
//}