package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

func main() {

	fmt.Printf("%s", time.Now().Format("2006-01-02 15:04:06,000"))

	/*var f float64

	fmt.Printf("\nf var %#v", f)


	var path = "/test/name/go/"

	fmt.Println(len(strings.Split(path, "/")))


	data := time.Now().Format("2006-01-02 15:04:05")

	fmt.Println(data)


	str := "aaaaaa_bbbb_cccc_dddd"

	n := strings.Split(str, "_")

	fmt.Println(n[0:len(n) - 1])


	c := str[2]

	fmt.Println(string(c))


	var BaseChars = [...]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D",
		"E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V","W","X", "Y", "Z",
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v",
		"w", "x", "y", "z",
	}

	fmt.Println(len(BaseChars))*/

	/*s1 := strings.Repeat("0", 8)
	s2 := strings.Repeat("0", 24)

	fmt.Println(strings.Compare(s1, s2))*/

	/*a := [2]int{1,2}
	b := [...]int{1,2}
	fmt.Println(a ==b)*/

	// 设置元素数量为1000
	const elementCount = 1000
	// 预分配足够多的元素切片
	srcData := make([]int, elementCount)
	// 将切片赋值
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}

	/*// 引用切片数据
	refData := srcData
	// 预分配足够多的元素切片
	copyData := make([]int, elementCount)
	// 将数据复制到新的切片空间中
	copy(copyData, srcData)
	// 修改原始数据的第一个元素
	srcData[0] = 999
	// 打印引用切片的第一个元素
	fmt.Println(refData[0])
	// 打印复制切片的第一个和最后一个元素
	fmt.Println(copyData[0], copyData[elementCount-1])
	// 复制原始数据从4到6(不包含)
	copy(copyData, srcData[4:6])
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", copyData[i])
	}*/

	/*l := list.New()


	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}*/

	/*s := make([]int, 0)
	s = append(s, 22)
	fmt.Println(s)


	var i *int

	b := 2

	i = &b

	fmt.Println(*i)*/

	str := "aaaaaa"

	type d struct {
		A string `json:"name"`
	}

	var D d
	err := json.Unmarshal([]byte(str), &D)

	fmt.Println("aaaaa", err)

	reflect.TypeOf(str).Name()

	fmt.Println(time.Now().Format("2006-01-02 15:04:05,000"))
}
