package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
) // 함수 종류들?!

func main() {

	// 변수 쓸 때 숫자 먼저 X, 알파벳 대문자 시작
	//  ex) var c7 (o) , var 7c (x)

	var c string  // 빈 문자열
	var d int     // 0
	var e bool    // false
	var f float64 // 0
	var G = 99

	fmt.Println(c, d, e, f, G)

	// koreanzombie := "정찬성"  // 문법적으로는 문제X, camel 케이스 사용 관례 따르기
	koreanZombie := "정찬성"
	fmt.Println(koreanZombie)

	var a int // declaration 선언       _1
	a = 7     // assign value 초기 값
	fmt.Println(a)

	// var a int = 7   declaration & assign value   _2

	// var a = 7	 declaration & assign value		_3

	// a := 7							_4

	fmt.Println(a, reflect.TypeOf(a)) // 7 int

	// b := 8.34    float64
	var b float32 = 8.34 // float 32

	// fmt.Println(a * b)		정수형, 실수형 끼리는 X
	// fmt.Println(a > b) 		float 64와 float32 도 연산 X

	fmt.Println(a * int(b))     // 56 , 형변환 해주면 가능 , 잠깐 바꿔준 것
	fmt.Println(float32(a) > b) // false

	fmt.Println(b, reflect.TypeOf(b)) //8.34

	fmt.Println(reflect.TypeOf('Z'), // int32
		reflect.TypeOf(2),     // int
		reflect.TypeOf("Hi"),  // string
		reflect.TypeOf(4.99),  // float64
		reflect.TypeOf(false)) // bool

	fmt.Println('Z', '2', '\n', '김') // 90 50 10 44608 유니코드로 계산
	fmt.Println(math.Floor(2.17), math.Ceil(2.17))
	strings.Title("오픈소스 프로그래밍")
	fmt.Println("Open Source")

}
