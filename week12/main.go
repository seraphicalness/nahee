package main

import "fmt"

func main() {
	a := make([]string, 4, 5) // type length, capacity
	a[0] = "a"
	a[2] = "c"
	a[3] = "d"
	as := a[0:2]
	as[1] = "Z"
	// c := append(a, "y")
	c := append(a, "y", "x")

	fmt.Println(a, len(a), cap(a)) // [a Z c d] 4 5
	fmt.Println(c, len(c), cap(c)) // [a Z c d y] 5 5
	fmt.Printf("%x %x %x\n", &a[0], &as[0], &c[0])

}

// func main() {
// 	a := []string{"a", "b", "c", "d"}
// 	as := a[:2]
// 	as[1] = "Z"
// 	fmt.Println(a, as)

// 	b := [4]int{4, 3, 2, 1}
// 	bs := b[1:3]
// 	fmt.Println(b, bs)
// }

// func main() {
// 	// var s []int
// 	// s = make([]int, 5) // 값이 없으니까 0으로만 나옴

// 	// s := make([]int, 5) // 단축변수 사용, 변수 선언과 동시에 메모리 할당, 원소들은 제로 값으로 초기화

// 	s := []int{0, 0, 0, 0, 0} // 단축 연산자 및 슬라이스 리터럴 사용, 변수 선언과 동시에 메모리 할당, 원소 초기화

// 	s[4] = 100
// 	s[0] = 7
// 	s[1] = 91

// 	for _, value := range s {
// 		fmt.Println(value) // 7 91 0 0 100
// 	}

// 	copyS := s[1:4]
// 	for _, value := range copyS { // 91 0 0
// 		fmt.Println(value)
// 	}

// 	test := [3]string{"inha", "go", "student"}
// 	testS := test[:2]
// 	testS2 := test[1:]

// 	testS2[0] = "python"
// 	// 역방향 인덱싱은 X, [-2] -> X

// 	// 전부 다 바뀜
// 	// 	[inha python student] 3
// 	// [inha python] 2
// 	// [python student] 2

// 	fmt.Println(test, len(test))     // [inha go student] 3
// 	fmt.Println(testS, len(testS))   // [inha go] 2
// 	fmt.Println(testS2, len(testS2)) // [go student] 2

// }
