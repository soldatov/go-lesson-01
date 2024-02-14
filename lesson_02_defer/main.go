package main

import "fmt"

func main() {
	/* Отложенный вызов работает по принципу стека. */
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("4")
}
