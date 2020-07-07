package chapter01

import (
	"fmt"
)

// c_a := 123 此种声明不能用作全局变量
var c_b int = 12
var c_c, c_d = false, "1234"
var x, y int64
var (
	cc_a int    = 12
	cc_b string = "asdfgh"
)

func VarDef() {
	//1、自推导，多个变量
	a, a1 := true, 34
	//2、直接定义
	var b int = 13
	//3、 b := 45 错误 必须有新的变量声明
	//4、自推导
	var c = 4.456
	//5、先声明，再赋值
	var s string
	s = "Hello "
	fmt.Println("a = ", a, " a1= ", a1, "b = ", b, "c = ", c, "s = ", s)

	fmt.Println(c_b , c_c, c_d, x, y, cc_a, cc_b)
}
