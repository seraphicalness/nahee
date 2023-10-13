package main

import (
	"bufio"
	"fmt"
	"log" // Fatal function
	"os"
	"strconv" // TrimSpace 쓰기위해
	"strings" // ParseInt 쓰기위해
)

func main() {
	fmt.Print("단 입력 : ")
	rd := bufio.NewReader(os.Stdin) // 읽는거
	in, err := rd.ReadString('\n')  // 입력된 문자를 in 에 저장, nil 이 나오면 에러 없다는 뜻

	if err != nil { // 에러 발생하면
		log.Fatal(err)

	}

	in = strings.TrimSpace(in) // 위에서 := 로 선언해서 = 씀

	// dan, err := strconv.ParseInt(in, 10, 32) , 10은 10진수로 하겠단 뜻
	dan, err := strconv.Atoi(in) // 위에 거를 따로 안 해도 괜찮게

	if err != nil {
		log.Fatal(err)

	}

	for i := 1; i < 10; i++ {
		fmt.Println(dan, " * ", i, " = ", (int(dan) * i))
		// fmt.Printf("%d * %d = %d\n", dan, i, (dan * i)) 이렇게 해도 괜찮
	}

	// 다른언어의 while 문 구현
	// i := 1
	// for i < 10 {
	// 	fmt.Printf("%d * %d = %d\n", dan, i, (dan * i))

	// }

	fmt.Println(dan * 2)

}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	fmt.Print("숫자 입력 : ")
// 	rd := bufio.NewReader(os.Stdin) // 읽는거
// 	in, _ := rd.ReadString('\n')    // 입력된 문자를 in 에 저장
// 	fmt.Println(in)
// }
