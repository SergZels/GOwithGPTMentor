package main

import (
	"fmt"
)

func multiply(a, b int) int {
	return a * b
}

func greet(name string) string {
	return "Hello " + name
}

func addToMap(m map[string]int, name string, age int) map[string]int {
	m[name] = age
	return m
}

type Person struct {
	name string
	age  int
	sity string
}

func (p Person) Greeting() {
	fmt.Printf("Name: %s Age: %d Sity: %s", p.name, p.age, p.sity)
}

func (p Person) Set(name string, age int, s string) {
	p.name = name
	p.age = age
	p.sity = s
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
	//var mymass [5]int
	//mymass[0] = 3
	//mymass[1] = 2
	//mymass[2] = 1
	//
	//fmt.Println(mymass)
	//
	////sli := []int{1, 2, 3, 4, 5}
	//sli := mymass[0:5]
	//sli = append(sli, 5, 9)
	//fmt.Printf("Slice:%d len-%d", sli, len(sli))
	//--------------lesson 7 map----------------------

	//piple := make(map[string]int)
	//piple["Ivan"] = 55
	//piple["Olga"] = 32

	//piple := map[string]int{
	//	"Olga": 25,
	//	"Dima": 17,
	//}
	//
	//for k, v := range piple {
	//
	//	fmt.Printf("%s %d\n", k, v)
	//}
	//var name string
	//for {
	//	fmt.Println("Enter your name: ")
	//	fmt.Scanln(&name)
	//	if name == "exit" {
	//		break
	//	}
	//	val, exes := piple[name]
	//	if exes {
	//		fmt.Printf("%s is %d", name, val)
	//	} else {
	//		fmt.Printf("%s not found enter age:", name)
	//		var age int
	//		fmt.Scanln(&age)
	//		piple = addToMap(piple, name, age)
	//		fmt.Println(piple)
	//
	//	}
	//}

	//--------------------lesson 8 ------------------------------

	Oleg := Person{name: "Oleg", age: 41, sity: "Tereb"}
	Oleg.Greeting()
	Oleg.Set("Ivan", 78, "")

}
