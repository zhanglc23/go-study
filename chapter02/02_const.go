package chapter02

import (
	"fmt"
)

const PI = 3.14

const (
	X = 12
	Y = 12
)

const (
	a1 = iota
	b1 = iota
	c1 = iota
)

const (
	a = iota //0
	b        //1
	c        //2
	d = "ha" //独立值，iota += 1
	e        //"ha"   iota += 1
	f = 100  //iota +=1
	g        //100  iota +=1
	h = iota //7,恢复计数
	i        //8
)

func Constant() {
	fmt.Println(PI, X, Y, a1, b1, c1, a, b, c, d, e, f, g, h, i)
}
