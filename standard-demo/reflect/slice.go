package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*s1 := make([]int, 3, 6)
	// 添加数据，未超出底层数组容量限制。
	s2 := append(s1, 1, 2, 3)
	// append 不会调整原 slice 属性
	// s1 == [0 0 0] len:3 cap:6
	fmt.Println(s1, len(s1), cap(s1))
	// 注意 append 是追加，也就是说在 s1 尾部添加。
	// s2 == [0 0 0 1 2 3] len:6 cap:6
	fmt.Println(s2, len(s2), cap(s2))
	// 追加的数据未超出底层数组容量限制。
	// 通过调整 s1，我们可以看到依然使⽤用的是原数组。
	// s1 == [0 0 0 1 2 3] len:6 cap:6
	s1 = s1[:cap(s1)]
	fmt.Println(s1, len(s1), cap(s1))

	s1[4] = 5
	fmt.Println("----------------------------")
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))*/

	var a interface{} = 22

	/*switch v := a.(type) {
	case int:
		fmt.Println("int", a.(int))
		default:
			fmt.Println(v)
	}*/

	t := reflect.TypeOf(a)

	v := reflect.ValueOf(a)
	fmt.Println(t.Name())
	fmt.Println(t.Kind())
	fmt.Println(t, v)

}
