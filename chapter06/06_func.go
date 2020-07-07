package chapter06

import (
	"fmt"
)

func Fun() {
	print(add(10, 10))
	minus(10, 10)

	_, d := swape(1, 2)
	print(d)

}

func print(a ...interface{}) {
	fmt.Println(a)
}
func add(a, b int) int {
	return a + b
}
func minus(a, b int) {
	print(a - b)
}
func swape(a, b int) (c, d int) {
	return b, a
}
