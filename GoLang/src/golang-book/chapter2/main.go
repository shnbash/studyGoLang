package main

import (
	"fmt"
)

type st_name struct {
	x       float64
	y       string
	z, a, b int
}

func func_args(args ...int) (int, int) {
	nTotal := 0
	for _, nValue := range args {
		nTotal += nValue
	}
	return nTotal, 100
}

func average(afValue []float64) float64 {
	fTotal := 0.0
	for _, fValue := range afValue {
		fTotal += fValue
	}
	return fTotal / float64(len(afValue))
}

func main() {
	//fmt.Println("07. Function")
	//var afValue11 [6]float64 { 89, 41, 99, 84, 66, 49 }
	afValue11 := []float64{89, 41, 99, 84, 66, 49}
	fmt.Println(average(afValue11))

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
	fmt.Println("CREATE SLICE 1")
	var afSlice1 []float64
	afSlice1 = make([]float64, 5)
	afSlice1[4] = 1
	fmt.Println(afSlice1)
	fmt.Println(len(afSlice1))
	fmt.Println("CREATE SLICE 2")
	afSlice2 := []float64{1, 2, 3, 4, 5}
	fmt.Println(afSlice2)
	fmt.Println(len(afSlice2))

	fmt.Println("CREATE SLICE 3")
	var afSlice3 []float64
	afSlice3 = make([]float64, 5, 10)
	fmt.Println(afSlice3)
	fmt.Println(len(afSlice3))

	fmt.Println("CREATE SLICE 4")
	afSlice4 := []float64{1, 2, 3, 4, 5}
	afSliceCur := afSlice4[1:3]
	fmt.Println(afSliceCur)
	fmt.Println(len(afSliceCur))

	fmt.Println("SLICE append")
	anSlice1 := []int{1, 2, 3, 4, 5}
	anSlice2 := append(anSlice1, 4, 5)
	fmt.Println(anSlice1)
	fmt.Println(len(anSlice1))
	fmt.Println(anSlice2)
	fmt.Println(len(anSlice2))
	fmt.Println("SLICE copy")
	anSlice3 := []int{1, 2, 3, 4, 5}
	var anSlice4 []int
	anSlice4 = make([]int, 5)
	copy(anSlice4, anSlice3)
	fmt.Println(anSlice3)
	fmt.Println(len(anSlice3))
	fmt.Println(anSlice4)
	fmt.Println(len(anSlice4))

	var mMap map[string]int
	mMap = make(map[string]int)
	mMap["key"] = 100
	fmt.Println(mMap["key"])

	mElements := make(map[string]string)
	mElements["A"] = "A10"
	mElements["B"] = "B10"
	mElements["C"] = "C10"
	mElements["D"] = "D10"
	mElements["E"] = "E10"
	fmt.Println(mElements)
	fmt.Println(mElements["E"])

	szValue, nRv := mElements["E"]
	fmt.Println(szValue, nRv)

	mElements1 := map[string]string{
		"H": "Hydrogen"}
	fmt.Println(mElements1["H"])

	//mElement2 := map[string]map[string]string{}

	anVal21 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(anVal21)
	var nRet1, nRet2 int
	nRet1, nRet2 = func_args(anVal21[0], anVal21[1], anVal21[2])
	fmt.Println(nRet1, nRet2)

	nVal21 := 0
	func_add := func(x string, y int) (int, string) {
		fmt.Println(x, y)
		y++
		return y, "haha"
	}
	fmt.Println(func_add("hoho", nVal21))

	fmt.Println("POINTER")
	//var pnVal1 *int
	//pnVal1 = new(int)
	var pnVal1 *int = new(int)
	*pnVal1 = 1
	fmt.Println(*pnVal1)

	var stName1 st_name
	stName2 := new(st_name)

	stName1.a = 1
	stName1.y = "dede"
	fmt.Println(stName1)

	stName2.b = 3
	stName2.y = "eee"
	fmt.Println(*stName2)
	fmt.Println(stName2.y)

	fmt.Println("TEST HP")

}
