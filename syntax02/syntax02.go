package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {

	tests := "G@ M@ney~~"
	fmt.Println(tests)
	r := strings.NewReplacer("@", "o")
	newTexts := r.Replace(tests) // @ -> o
	fmt.Println(newTexts)

	// 변수명은 영문자로 시작
	// 영문 대문자의 경우 다른 패키지에서 접근할 수 있다.
	// 소문자로 시작하는 변수는 동일 패키지에서만 접근가능
	var a int     // 0
	var d bool    // 0
	var b float64 // 0
	var c rune    // false
	var e string  //
	// a := 7

	// zero value (제로 값)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)

	fmt.Println(reflect.TypeOf(d))
	fmt.Println(reflect.TypeOf(e))
}
