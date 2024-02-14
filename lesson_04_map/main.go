package main

import "fmt"
import "github.com/mitchellh/mapstructure"

type Point struct {
	X int
	Y int
}

func (p Point) doPrint() {
	fmt.Println("Hi")
}

func main() {
	/*
		Фигурные скобки в конце означают что мы этот map инициизировали.
		Аналог make(map[string]Point).
		Объявление без иницииализации var pointMap map[string]Point
	*/
	pointMap := map[string]Point{}
	pointMap["a"] = Point{
		X: 2,
		Y: 2,
	}
	pointMap["b"] = Point{4, 4}
	fmt.Println(pointMap)

	/* Еще один вариант создания карты */

	pointMap2 := map[string]Point{
		"a": {
			X: 2,
			Y: 2,
		},
		"b": {4, 4},
	}
	fmt.Println(pointMap2)

	pointMap2["a"].doPrint()
	pointMap2["b"].doPrint()

	/* Как проверить наличие элемента по ключу */

	value, ok := pointMap2["a"]
	if ok {
		fmt.Println(fmt.Sprintf("Value X: %d; Value Y: %d", value.X, value.Y))
	}

	/* Конвертирование map-ов */

	pointMap3 := map[string]int{
		"x": 1,
		"y": 2,
	}

	p3 := Point{}
	mapstructure.Decode(pointMap3, &p3)
	fmt.Println(p3)

	pointMap4 := map[string]int{
		"xxx": 1,
		"yyy": 2,
	}

	p4 := Point{}
	mapstructure.Decode(pointMap4, &p4)
	fmt.Println(p4)

	/* Итерирование map */

	for k, v := range pointMap2 {
		fmt.Println(fmt.Sprintf("Key: %s; Value X: %d; Value Y: %d", k, v.X, v.Y))
	}
}
