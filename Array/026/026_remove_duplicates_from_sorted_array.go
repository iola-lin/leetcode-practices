package main

import "fmt"

type sortedArr []int

func main() {
	// nums0 := sortedArr {}
	n1 := sortedArr{1, 1, 2}
	n2 := sortedArr{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	var l1, l2 int

	// nums0.removeDuplicate()
	// fmt.Println("%v\n", num0)
	n1, l1 = removeDuplicate(n1)
	n2, l2 = removeDuplicate(n2)
	fmt.Println(n1)
	fmt.Println("length: ", l1)
	fmt.Println(n2)
	fmt.Println("length: ", l2)
}

func removeDuplicate(a sortedArr) (b sortedArr, l int) {
	if len(a) == 0 {
		return
	}
	for i, num := range a {
		if i >= len(a) {
			break
		}
		if i == 0 {
			l++
			continue
		}
		if num != a[l-1] {
			a = append(a[:l], a[i:]...)
			l++
		}
		fmt.Println(a)
	}
	b = a
	return
}
