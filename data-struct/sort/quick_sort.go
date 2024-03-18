package main

import "fmt"

/**
 * 分区排序
 */
func partition(L *[]int, low int, high int) int {
	sign := (*L)[low]
	for low < high {
		for (high > low) && (*L)[high] >= sign {
			high--
		}
		(*L)[low] = (*L)[high]
		for (low < high) && (*L)[low] < sign {
			low++
		}
		(*L)[high] = (*L)[low]
	}
	(*L)[low] = sign
	return low
}

func qsort(L *[]int, low int, high int) {
	if low < high {
		mark := partition(L, low, high)
		qsort(L, low, mark-1)
		qsort(L, mark+1, high)
	}
}

func Qsort(L *[]int) {
	qsort(L, 0, len((*L))-1)
}

func main() {
	L := []int{1, 9, 5, 8, 7, 2, 4, 3, 6, 10}
	fmt.Println("L before sort:", L)
	Qsort(&L)
	fmt.Print("after sort:", L)
}
