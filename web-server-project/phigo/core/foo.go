package core

import (
	"fmt"
	"phigo/core/dto"
)

// Hello - my stub fun
func Hello() string {
	return "asd"
}

func asd() string {
	x := dto.MyData{Name: "asd", Val: 123}
	return fmt.Sprint(x)
}