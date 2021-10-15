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
	var (
		a = 10
		b = 20
		c = 30
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Print("Chanpter 05. Control")
	var nIdx int = 0
	fmt.Print("Chanpter 05-1. for")
	for nIdx = 0; nIdx < 10; nIdx++ {
		fmt.Println(nIdx)
		fmt.Print("Chanpter 05-2. if")
		if nIdx%2 == 0 {
			fmt.Println("짝수")
		} else {
			fmt.Println("홀수")
		}
		fmt.Print("Chanpter 05-4. switch")
		switch nIdx {
		case 1:
			fmt.Print("switch 1")
		case 9:
			fmt.Print("switch 9")
		}
	}

	fmt.Println("Chanpter 06. ARRAY, SLICE, MAP")
	fmt.Println("Chanpter 06-1. ARRAY")
	var anVal [5]int
	for nIdx = 0; nIdx < 5; nIdx++ {
		anVal[nIdx] = nIdx + 10
	}
	fmt.Println("anVal:", anVal)

	var fTotal float64
	afVal := [5]float64{98, 93, 77, 82, 83}

	//for nIdx = 0; nIdx < len(afVal); nIdx++ {
	for _, nValue := range afVal {
		fTotal += nValue
	}
	fmt.Println("Total:", fTotal/(float64)(len(afVal)))

	fmt.Println("Chanpter 06-2. SLICE")
	var afSlice []float64
	afSlice = make([]float64, 5)
	afSlice[4] = 1
	fmt.Println(afSlice)
	fmt.Println(len(afSlice))

}
