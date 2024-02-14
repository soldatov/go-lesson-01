package main

import "fmt"

var _ summer = &numbersTwo{} // Это подсказка. Если будет изменения сигнатура интерфейса, тут будет видно ошику.
type numbersTwo struct {
	n1, n2 int
}

func (ns *numbersTwo) getSum() int {
	return ns.n1 + ns.n2
}

var _ summer = &numbersTree{}

type numbersTree struct {
	n1, n2, n3 int
}

func (ns *numbersTree) getSum() int {
	return ns.n1 + ns.n2 + ns.n3
}

type summer interface {
	getSum() int
}

func main() {
	var a2, a3 summer
	ns2 := numbersTwo{1, 2}
	ns3 := numbersTree{1, 2, 3}

	a2 = &ns2
	a3 = &ns3

	fmt.Println(a2.getSum()) // 3
	fmt.Println(a3.getSum()) // 6

	/* Еще вариант создания */
	var a33 summer = &numbersTree{2, 2, 2}
	fmt.Println(a33.getSum()) // 6
}
