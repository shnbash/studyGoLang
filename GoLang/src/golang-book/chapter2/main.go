package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("1.0 + 1.0 =", 1.0+1.0)
	fmt.Println(len("Hello World!"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World" + "!!")
	fmt.Println("Test Boolean")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println("Test Variable")
	var szVal string = "Hello World~!"
	fmt.Println("szVal: ", szVal)
	var szVal1 string
	szVal1 = "first"
	fmt.Println(szVal1)
	szVal1 = "second"
	fmt.Println(szVal1)
	var szVal2 string
	szVal2 = "First "
	fmt.Println(szVal2)
	szVal2 = szVal2 + "Second"
	fmt.Println(szVal2)
	var szVal3 string = "first"
	var szVal4 string = "first"
	fmt.Println(szVal3 == szVal4)
	//var nVal1 int32 = 5
	nVal1 := 5
	nVal2 := 10
	fmt.Println("nVal1: ", nVal1, "nVal2:", nVal2)
}
