package main

import (
	"fmt"
	"sync"

	"math/rand"
	"time"
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

type NyInt interface {
	Foo()
}

type MyStruct struct {
	name string
}

func (s MyStruct) Foo() {
	fmt.Println(s.name)
}

func someSleep(i int) {
	random := rand.Intn(5000)                            // Генеруємо випадкову кількість мілісекунд
	time.Sleep(time.Duration(random) * time.Millisecond) // Конвертуємо в мілісекунди
	fmt.Printf("Hi from %d\n", i)
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

	//Oleg := Person{name: "Oleg", age: 41, sity: "Tereb"}
	//Oleg.Greeting()
	//Oleg.Set("Ivan", 78, "")
	//----------------------les 9-----------------------
	//var i NyInt
	//
	//i = MyStruct{
	//	name: "HHH",
	//}
	//i.Foo()
	//-----------------------les 10-----------------

	//for i := 0; i < 10; i++ {
	//	go someSleep(i)
	//}
	//time.Sleep(time.Second * 10)
	//
	//ch := make(chan string)
	//
	//go sendDataCh(ch, "Hi from go rutine")
	//valFromCh := <-ch
	//
	//fmt.Printf("ValFromCh: %s\n", valFromCh)

	//---------------11--------------------------
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // додаємо горутину до групи
		go func(i int) {
			defer wg.Done() // повідомляємо про завершення горутини
			time.Sleep(time.Second * 2)
			fmt.Printf("Goroutine %d done\n", i)
		}(i)
	}

	wg.Wait() // чекаємо на завершення всіх горутин
	fmt.Println("All goroutines finished")

	for i := 0; i < 1000; i++ {

		go increment2()
	}

	fmt.Println("Final counter value2:", counter2)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter)
}

func sendDataCh(ch chan string, data string) {
	ch <- data
}

var mu sync.Mutex
var counter int
var counter2 int

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock() // блокуємо доступ до змінної counter
	counter++
	mu.Unlock() // знімаємо блокування
}
func increment2() {
	counter2++
}
