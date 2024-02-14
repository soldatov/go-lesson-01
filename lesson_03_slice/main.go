package main

import "fmt"

func main() {
	/* Пример слайсов (массивы без размерности) */

	letters := []string{"a", "b"}
	letters = append(letters, "c", "error")
	letters[3] = "d"
	fmt.Println(letters)

	letters2 := make([]string, 3) // Создает слайс с тремя эелементами с пустыми строками.
	letters2[0] = "A"
	letters2[1] = "B"
	// letters2[3] = "D" выведет панику, потому что мы объявили 3 а обращается к 4-му. Нужно через append
	letters2 = append(letters2, "C")
	fmt.Println(letters2) // вернет ["A" "B" "" "C"]

	fmt.Println(len(letters2)) // 4 всего элемтов
	fmt.Println(cap(letters2)) // 6 капасити, размер в памяти. Не всегда совпадает с количестовм.

	/* Замена элемента в массивах */

	animalArr := [4]string{
		"dog",
		"cat",
		"pig",
		"wolf",
	}

	aArr := animalArr[0:2]
	fmt.Println(aArr) // [dog cat]
	animalArr[0] = "rabbit"
	fmt.Println(aArr) // [rabbit cat], хотя в переменной aArr мы не меняли.

	/* Замена элемента в слайсах */

	animalSlice := []string{
		"dog",
		"cat",
		"pig",
		"wolf",
	}

	aSlice := animalSlice[0:2]
	fmt.Println(aSlice) // [dog cat]
	animalSlice[0] = "cow"
	fmt.Println(aSlice) // [cow cat], в слайте так же, заменился, хотя animalSlice не меняли.

	fmt.Println(animalSlice) // [cow cat pig wolf]
	aSlice[0] = "bee"
	fmt.Println(animalSlice) // [bee cat pig wolf]

	/* Сокращения */

	aSlice = animalSlice[:2] // C начала до 3-го
	aSlice = animalSlice[1:] // Cо второго до конца
	aSlice = animalSlice[:]  // С начала до конца

	/* Сравнение с nil */

	nilSlice := []string{}
	fmt.Println(nilSlice, len(nilSlice), cap(nilSlice)) // [] 0 0

	if nilSlice == nil {
		fmt.Println("Slice is nil") // не показывает
	}

	var nilSlice2 []string
	fmt.Println(nilSlice2, len(nilSlice2), cap(nilSlice2)) // [] 0 0

	if nilSlice2 == nil {
		fmt.Println("Slice is nil") // Показывает
	}

}
