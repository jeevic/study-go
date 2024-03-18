package main

import "fmt"

func PrintSort(param [10]int) {
	for k, v := range param {
		fmt.Printf("%d  ", v)
		_ = k
	}
	fmt.Println()
}

func InsertSort(a [10]int) [10]int {
	length := len(a)
	for i := 1; i < length; i++ {
		sign := a[i]
		k := i
		for j := i - 1; j >= 0; j-- {
			if a[j] < sign {
				break
			}
			k--
		}
		if k < i {
			for l := i; l > k; l-- {
				a[l] = a[l-1]
			}
			a[k] = sign
		}
	}
	return a
}

func main() {
	var a = [10]int{1, 5, 7, 3, 2, 10, 9, 4, 6, 8}

	fmt.Print(" 插入排序前: ")
	PrintSort(a)
	a = InsertSort(a)
	fmt.Print(" 插入后排序: ")
	PrintSort(a)

}
