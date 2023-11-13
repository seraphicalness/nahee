package main

import "fmt"

func main() {

	// 1
	// var primes [3]int
	// primes[0] = 2
	// primes[1] = 3
	// primes[2] = 5
	// fmt.Println(primes, primes[1]) // [2 3 5] 3

	// 2
	// var primes [3]int = [3]int{2, 3, 5} // 배열 리터럴로 primes 배열 초기화
	// fmt.Println(primes, primes[1])

	// 3
	primes := [3]int{2, 3, 5}
	fmt.Println(primes, primes[1])

}
