// package main
//
// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// 	"time"
// )
//
// func main() {
// 	var now time.Time
//
// 	now = time.Now()
// 	// var year int = now.Year()
// 	var year = now.Year()
// 	month := now.Month()
// 	fmt.Println(year)
// 	fmt.Println(year, month, now.Hour(), now.Minute(), now.Second())
// }

// package main
//
// import (
//    "bufio"
//    "fmt"
//    "os"
//    "strconv"
//    "strings"
//
// )
//
// func main() {
//    fmt.Print("Input score : ")
//    reader := bufio.NewReader(os.Stdin)
//    inputScoreString, err := reader.ReadString('\n') // option 2
//    if err != nil{
// 		log.Fatal(err)
//    }
//    inputScoreString = strings.TrimSpace(inputScoreString) // remove space
//    inputScore = strconv.ParseFloat(inputScoreString, 32) // mismatched types string and untype
//
//    var grade string
//
//    if inputScore >= 90{
// 	grade := "A"
//    } else{
// 	grade := "under A"
//    }
// }
//

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Input score : ")
	reader := bufio.NewReader(os.Stdin)
	inputScoreString, err := reader.ReadString('\n') // option 2
	if err != nil {
		log.Fatal(err)
	}
	inputScoreString = strings.TrimSpace(inputScoreString)      // remove space
	inputScore, err := strconv.ParseFloat(inputScoreString, 32) // string to 32bit float
	//fmt.Println(inputScore)
	var grade string
	if inputScore >= 90 {
		grade = "A grade!"
		//fmt.Println("You got", grade)
	} else {
		grade = "under A grade..."
		//fmt.Println("You got", grade)
	}
	fmt.Println("You got", grade)
}


// 게임 만들기

package main

import(
	"log"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"math/rand"

)
// func main(){
// 	dice := rand.Intn(6) + 1
// 	fmt.Println(dice)
// }
// 

func main(){
	fmt.Println("Guess number game!")
	answer = rand.Intn(100) + 1  // 1 ~ 100
	fmt.Println(answer)
	
//  	inputNumberString, err :=
//  
//  	fmt.Print("Input number : ")
//  	reader := bufio.NewReader(os.Stdin)
//  	inputGuessString, err := reader.ReadString('\n')
//  	if err != nil {
//  		log.Fatal(err)
//  	}
//  	inputGuessString = strings.TrimSpace(inputScoreString)   
//  	inputGuess, err := strconv.Atoi(inputScoreString)
//  	if err != nil{
//  		log.Fatal(err)
//  	}
//  	if inputGuess < answer{
//  		fmt.Println("too Low")
//  	}else if inputGuess > answer{
//  		fmt.Println("too High")
//  	}
//  
		// for 문
	for i := 0; i < 10; i++{

		fmt.Print("Input number : ")
		reader := bufio.NewReader(os.Stdin)
		inputGuessString, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		inputGuessString = strings.TrimSpace(inputScoreString)   
		inputGuess, err := strconv.Atoi(inputScoreString)
		if err != nil{
			log.Fatal(err)
		}
		if inputGuess < answer{
			fmt.Println("too Low")
		}else if inputGuess > answer{
			fmt.Println("too High")
		}else if inputGuess == answer{
			fmt.Println("Great!")
			break
		}

	}

}