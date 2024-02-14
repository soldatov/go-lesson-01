package main

import "fmt"

func main() {
	arr := []string{"a", "b", "c"}
	for k, v := range arr {
		fmt.Println(k, v) // 0 a  1 b  2 c
	}

	/* Если нужно пропустить ключ (так же можно пропустить value) */
	for _, v := range arr {
		fmt.Println(v) // a b c
	}
}
