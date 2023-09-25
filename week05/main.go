package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time

	now = time.Now()
	// var year int = now.Year()
	var year = now.Year()
	month := now.Month()
	fmt.Println(year)
	fmt.Println(year, month, now.Hour(), now.Minute(), now.Second())
}

package main

import (
   "bufio"
   "fmt"
   "os"
)

func main() {
   fmt.Print("Input score : ")
   reader := bufio.NewReader(os.Stdin)
   inputScore := reader.ReadString('\n') // 1 variable but reader.ReadString returns 2 values
   fmt.Println(inputScore)
}
