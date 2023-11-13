package main

import "fmt"

func main() {
	var s []int
	s = make([]int, 5) // 값이 없으니까 0으로만 나옴

	for _, value := range s {
		fmt.Println(value)
	}
}
