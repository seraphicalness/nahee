package main

import "fmt"

func main() {
	a := 10 // var a int = 10
	var pa *int
	pa = &a

	fmt.Println(a, *pa)                         // 10 10
	fmt.Println(&a, pa)                         // 주소값
	fmt.Printf("%T %T %T %T\n", a, *pa, &a, pa) // int int *int *int
	fmt.Println(&pa)                            //  위에 &a , p 와는 다른 주소값
}
