package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	sum := 0
	for sum < 3 { // это аналог while, тк while нет.
		fmt.Println(sum) // Нельзя указывать i, он инициировался внутри первого for
		sum++
	}
}
