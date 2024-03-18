package main

import "fmt"

func PrintSort(param [10]int) {
	for k, v := range param {
		fmt.Printf("%d  ", v)
		_ = k
	}
	fmt.Println()
}

func MiddleInsertSort(a [10]int) [10]int {
	length := len(a)
	var low, mid, high int
	for i := 1; i < length; i++ {
		low = 0
		high = i - 1
		for low <= high {
			mid = (low + high) / 2
			if a[mid] > a[i] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		fmt.Printf("i = %d low = %d\n", i, low)
		sign := a[i]
		for j := i; j > low; j-- {
			a[j] = a[j-1]
		}
		fmt.Println("----")
		a[low] = sign
		fmt.Println(a)
	}
	return a
}

func main() {
	var a = [10]int{1, 5, 7, 3, 2, 10, 9, 4, 6, 8}

	fmt.Print(" 插入排序前: ")
	PrintSort(a)
	a = MiddleInsertSort(a)
	fmt.Print(" 插入后排序: ")
	PrintSort(a)

}
