package main

import (
	"fmt"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	//TIP Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined or highlighted text
	// to see how GoLand suggests fixing it.
	var age int

	fmt.Print("Введіть ваш вік:")
	fmt.Scanln(&age)

	if age < 18 {
		fmt.Println("Ви ще не повнолітній")
	} else if age > 18 && age < 65 {
		fmt.Println("Вітаю Ви повнолітній")
	} else {
		fmt.Println("You are a senior citizen.")
	}

}
