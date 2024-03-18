package main

import "fmt"

/*func main() {
	var a int
	typeOfA := reflect.TypeOf(a)
	fmt.Println(typeOfA.Name(), typeOfA.Kind())
}*/

/*func main() {
	var wg sync.WaitGroup
	typ := reflect.TypeOf(&wg)
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		argv := make([]string, 0, method.Type.NumIn())
		returns := make([]string, 0, method.Type.NumOut())
		// j 从 1 开始，第 0 个入参是 wg 自己。
		for j := 1; j < method.Type.NumIn(); j++ {
			argv = append(argv, method.Type.In(j).Name())
		}
		for j := 0; j < method.Type.NumOut(); j++ {
			returns = append(returns, method.Type.Out(j).Name())
		}
		log.Printf("func (w *%s) %s(%s) %s",
			typ.Elem().Name(),
			method.Name,
			strings.Join(argv, ","),
			strings.Join(returns, ","))
	}
}*/

func main() {
	/*var x uint8 = 'x'
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())                            // uint8.
	fmt.Println("kind is uint8: ", v.Kind() == reflect.Uint8) // true.
	x = uint8(v.Uint())      */ // v.Uint returns a uint64.

	/*type MyInt int
	var y MyInt = 7
	yv := reflect.ValueOf(&y)
	fmt.Println("type:", yv.Type())                            // uint8.
	fmt.Println("y value is can set:",  yv.CanSet())
	yv.Elem().SetInt(8)
	fmt.Println("type:", yv)
	fmt.Println("y value is type:", yv.Type())
	fmt.Println("kind is uint8: ", yv.Kind() == reflect.Int) // true.
	fmt.Println("y ele value is can set:",  yv.Elem().CanSet())
	fmt.Println(y)*/

	/*var a *int
	fmt.Println("var a *int:", reflect.ValueOf(a).IsNil())
	// nil值
	fmt.Println("nil:", reflect.ValueOf(nil).IsValid())
	// *int类型的空指针
	fmt.Println("(*int)(nil):", reflect.ValueOf((*int)(nil)).Elem().IsValid())
	// 实例化一个结构体
	s := struct{}{}
	// 尝试从结构体中查找一个不存在的字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(s).FieldByName("").IsValid())
	// 尝试从结构体中查找一个不存在的方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(s).MethodByName("").IsValid())
	// 实例化一个map
	m := map[int]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("不存在的键：", reflect.ValueOf(m).MapIndex(reflect.ValueOf(3)).IsValid())*/

	type Name struct {
		Id    int
		Class string
	}

	a := new(Name)

	fmt.Printf("%#v", a)

}
