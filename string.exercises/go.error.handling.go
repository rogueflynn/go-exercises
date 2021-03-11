package main

import (
	"fmt"
)

type ErrNegativeSqrt struct {
	What string
}
func (e *ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("%s", e.What)
}

func Sqrt(x float64) (float64, error) {
	if(x < 0) {
		return 0, &ErrNegativeSqrt{
			fmt.Sprintf("cannot Sqrt negative number: %f", x),
		}
	}
	z := 1.0
	for i:= 0; i < 10; i++ {
		z -= (z*z-x)/2*z
		fmt.Println(z)
	}
	return z, nil
}


func main() {
	i, err := Sqrt(-2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}
}