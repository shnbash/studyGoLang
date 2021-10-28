package main

import (
	"fmt"
)

type MyInt int

func (nMyInt MyInt) print() {
	fmt.Println(nMyInt)
}

type printer_i interface {
	print()
}
type hello_i interface {
}

func main() {
	var ifHello hello_i
	fmt.Println(ifHello)

	var nMyInt MyInt = 10

	var ifPrinter printer_i

	ifPrinter = nMyInt

	ifPrinter.print()
}
