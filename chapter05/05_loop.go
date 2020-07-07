package chapter05

import (
	"fmt"
)

func Loop() {

	for i := 0; i < 3; i++ {
		fmt.Print(i, " ")
	}
	fmt.Print("\n")

	numbers := [6]int{1, 2, 3, 5}
	for i, x := range numbers {
		if i == 1 {
			continue
		}
		if i == 2 {
			goto LABLE
		}
		if i == 3 {
			break
		}
		fmt.Print(i, "-", x, " ")
	}
	fmt.Print("\n")
LABLE:
	fmt.Println("这是goto执行了")

	// for true {
	// }
	// for {
	// }
}
