package main

import "fmt"

type Set struct {
	buf  []interface{}        //存储数据
	num  int                  //数量
	hash map[interface{}]bool //借助 map 实现映射
}

// 新建一个可以变长的set
func NewSet() *Set {
	return &Set{make([]interface{}, 0), 0, map[interface{}]bool{}}
}

func (this *Set) Add(value interface{}) bool {
	if this.isExist(value) {
		return false
	} else {
		this.buf = append(this.buf, value)
		this.hash[value] = true
		this.num++
		return true
	}
}

func (this *Set) isExist(value interface{}) bool {
	return this.hash[value]
}

func main()  {
	set := NewSet()
	set.Add(1)
	set.Add(2)
	set.Add("3")
	set.Add(3)
	fmt.Println(set)
}