package main

import (
	"fmt"
)


type newType float64

func (a newType) IsEqual (b newType) bool {
	var r = a - b;
	if r == 0.0 {
		return true
	} else if r < 0.0 {
		return r > -0.0001
	}
	return r < 0.0001

}

func IsEqual(a, b  float64) bool {
	var r = a - b
	if r == 0.0 {
		return true
	} else if r < 0.0 {
		return r > -0.0001
	}
	return r < 0.0001
}


func main() {
	var a newType = 1.999999
	var b newType  = 1.9999998

	fmt.Println(a.IsEqual(b))
	fmt.Println(a.IsEqual(3))

	fmt.Println( IsEqual((float64)(a), (float64)(b) ) )
}







