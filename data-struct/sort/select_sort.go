package main

import "fmt"

func PrintSort(param [10]int) {
	for k, v := range param {
		fmt.Printf("%d  ", v)
		_ = k
	}
	fmt.Println()
}

func SelectSort(param [10]int) [10]int {
	length := len(param)
	for i := 0; i < length; i++ {
		k := i
		for j := i + 1; j < length; j++ {
			if param[k] > param[j] {
				k = j
			}
		}

		if k != i {
			param[i], param[k] = param[k], param[i]
		}
	}
	return param
}

func main() {
	var a = [10]int{1, 5, 7, 3, 2, 10, 9, 4, 6, 8}

	fmt.Print(" 选择排序前: ")
	PrintSort(a)
	a = SelectSort(a)
	fmt.Print(" 选择后排序: ")
	PrintSort(a)

}
