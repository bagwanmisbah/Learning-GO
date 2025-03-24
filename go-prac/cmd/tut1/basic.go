package main

import "fmt"

type gas_engine struct {
	mpg     uint8
	gallons uint8
}

type elec_eng struct {
	mpkwh uint8
	kwh   uint8
}

func (e gas_engine) miles_left() uint8 {
	return e.mpg * e.gallons
}
func (e elec_eng) miles_left() uint8 {
	return e.mpkwh * e.kwh
}

func can_make_it(e gas_engine, miles uint8) {
	if miles <= e.miles_left() {
		fmt.Println("You can make it till there with the existing amount of fuel")
	} else {
		fmt.Println("You need to fuel up first in order to reach the destination")
	}

}

func main() {
	var my_engine1 gas_engine
	fmt.Println(my_engine1.mpg, my_engine1.gallons)
	var my_engine2 gas_engine = gas_engine{mpg: 25, gallons: 10}
	fmt.Println(my_engine2.mpg, my_engine2.gallons)
	var e gas_engine = gas_engine{2, 15}
	fmt.Println("Miles left : ", e.miles_left())
	var miles_needed uint8 = 80
	can_make_it(e, miles_needed)

}
