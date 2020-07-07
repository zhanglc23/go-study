package chapter04

import (
	"fmt"
)

var g string

func IfCondition(grade int) {
	if grade > 90 && grade <= 100 {
		g = "优秀"
		fmt.Println("优秀")
	} else if grade > 80 && grade <= 90 {
		g = "良好"
		fmt.Println("良好")
	} else {
		g = "一般"
		fmt.Println("一般")
	}

	switch g {
	case "优秀":
		fmt.Println("优秀")
	case "良好":
		fmt.Println("良好")
		fallthrough
	case "一般":
		fmt.Println("一般")
		fallthrough
	case "ddd":
		fmt.Println("ddd")
		fallthrough
	case "fff":
		fmt.Println("fff")
		// fallthrough
	default:
		fmt.Println("差")
	}
}
