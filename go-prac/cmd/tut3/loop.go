package main

import "fmt"

func main() {

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	var i int32 = 50
	var p *int32 = &i
	fmt.Println("Address stored in p: ", p)
	fmt.Println("Value that p points to: ", *p)
	fmt.Println("Adress of Value that p points to: ", &(*p))
	fmt.Println("Adress of p : ", (&p))
	// star * symbol is used to dereference the pointer
	//modifying the value that p points to
	*p = (*p) * 2
	//is that same as i=i*2
	fmt.Println("Value that p points to: ", *p)
	i = i * 3
	fmt.Println("Value that p points to: ", *p)

}
