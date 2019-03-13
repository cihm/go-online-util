package oop

import (
	"fmt"
)

func Flow() {
	var a Factory
	value := a.FacMethod(20, 3, "-")
	fmt.Println(value)
}
