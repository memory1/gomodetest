package main

import (
	"fmt"
	"gomodetest"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello()) 
	eq := CompareLength(1,2)
	if eq {
		fmt.Println("equal")
	}else
	{
		fmt.Println("false")
	}
}