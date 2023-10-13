// 입력(fmt 패키지의 Scanln())된 수의 소수 판정 프로그램 v0.8

package main

import (
	"fmt"
	"log"
	"os"
)

// 소수 : 1과 자기자신 외에는 나누어 떨어지지 않는 수 (0과 1은 제외)
func main() {
	var number int
	fmt.Print("정수 입력 : ")
	_, err := fmt.Scanln(&number) // 에러나면 돌지않음

	// n, err := fmt.Scanln(&number)
	// fmt.Println(n)  // n 은 개수여서 출력됨

	if err != nil {
		log.Fatal(err)
	}
	for number < 2 {
		fmt.Println(number, "는(은) 소수가 아닙니다")
		os.Exit(0) // 프로그램 종료

	}

	isPrime := true               // 소수가 true
	for i := 2; i < number; i++ { // 1과 number 일 때 loop 돌지 않음
		if number%i == 0 {
			isPrime = false
			break
		}
	}

	if isPrime == true {
		fmt.Println(number, "는(은) 소수입니다!")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다!")
	}

}

// 난수 추출된 수의 소수 판정 프로그램 v0.7

// package main

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// 	// "math/rand"
// 	"bufio"
// 	"log"
// 	"os"
// 	// "time" // seed 생성용 패키지
// )

// // 소수 : 1과 자기자신 외에는 나누어 떨어지지 않는 수 (0과 1은 제외)
// func main() {

// 	fmt.Print("정수 입력 : ")
// 	rd := bufio.NewReader(os.Stdin) // 읽는거
// 	in, err := rd.ReadString('\n')  // 입력된 문자를 in 에 저장, nil 이 나오면 에러 없다는 뜻

// 	if err != nil { // 에러 발생하면
// 		log.Fatal(err)

// 	}
// 	// TrimSpace() 써서 양 옆 공백 제거
// 	in = strings.TrimSpace(in) // 위에서 := 로 선언해서 = 씀
// 	// dan, err := strconv.ParseInt(in, 10, 32) , 10은 10진수로 하겠단 뜻
// 	number, err := strconv.Atoi(in) // 위에 거를 따로 안 해도 괜찮게
// 	// Atoi 함수를 써서 정수로 바꿈
// 	if err != nil {
// 		log.Fatal(err)

// 	}

// 	isPrime := true               // 소수가 true
// 	for i := 2; i < number; i++ { // 1과 number 일 때 loop 돌지 않음
// 		if number%i == 0 {
// 			isPrime = false
// 			break
// 		}
// 	}

// 	if isPrime == true {
// 		fmt.Println(number, "는(은) 소수입니다!")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다!")
// 	}

// }
