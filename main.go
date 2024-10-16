package main

import "fmt"

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func multiply(a, b int) int {
	return a * b
}

func greet(name string) string {
	return "Hello " + name
}

func main() {

	//---------------lesson1------------------------
	//	s := "Sergii"
	//	tim := time.Now().Format("2006-01-02")
	//	fmt.Printf("Hello and welcome, %s %s!", s, tim)

	//--------------lesson 4-------------------------------------
	//summ := 0
	//for i := 1; i <= 10; i++ {
	//	fmt.Println(i)
	//	summ += i
	//}
	//fmt.Printf("summ:%d\n", summ)

	//-------------lesson5 functions---------------

	//rez := multiply(3, 4)
	//println(rez)
	//
	//println(greet("Serg"))
	//---------------lesson 6-------------------
	//mymass := [...]int{4, 5, 6, 2, 9}
	var mymass [5]int
	mymass[0] = 3
	mymass[1] = 2
	mymass[2] = 1

	fmt.Println(mymass)

	//sli := []int{1, 2, 3, 4, 5}
	sli := mymass[0:5]
	sli = append(sli, 5, 9)
	fmt.Printf("Slice:%d len-%d", sli, len(sli))

}
