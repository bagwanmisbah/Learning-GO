package main

import (
	"fmt"
)

func main() {
	var f1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("Address of f1 : %p", &f1)
	var res = square(&f1)
	fmt.Println(res)
	// fmt.Printf("Address of res : %p", &res)

}

func square(f2 *[5]float64) [5]float64 {
	fmt.Printf("Address of f2 : %p", &f2)
	for i := range f2 {
		f2[i] = f2[i] * f2[i]
	}

	return *f2
}
