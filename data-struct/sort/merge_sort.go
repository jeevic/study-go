package main

import "fmt"

func Merge(sr *[]int, tr *[]int, i int, m int, n int) {

	j := m + 1
	k := i
	for i <= m && j <= n {
		if (*sr)[i] < (*sr)[j] {
			(*tr)[k] = (*sr)[i]
			i++
		} else {
			(*tr)[k] = (*sr)[j]
			j++
		}
		k++
	}

	for i <= m {
		(*tr)[k] = (*sr)[i]
		k++
		i++
	}

	for j <= n {
		(*tr)[k] = (*sr)[j]
		k++
		j++
	}
}

func MSort(sr *[]int, tr1 *[]int, s int, t int) {
	tr2 := make([]int, 10)

	if s == t {
		(*tr1)[s] = (*sr)[s]
	} else {
		m := (s + t) / 2
		MSort(sr, &tr2, s, m)
		MSort(sr, &tr2, m+1, t)
		Merge(&tr2, tr1, s, m, t)

	}
}

func MergeSort(sr *[]int) {
	MSort(sr, sr, 0, len((*sr))-1)
}

func main() {
	L := []int{1, 9, 5, 8, 7, 2, 4, 3, 6, 10}
	fmt.Println("L before sort:", L)
	MergeSort(&L)
	fmt.Println("L after sort:", L)
}
