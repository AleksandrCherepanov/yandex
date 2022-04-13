package main

import (
	"fmt"

	"github.com/AleksandrCherepanov/yandex/algo_training_1/lesson_1"
)

func main() {
	var a, b, c int

	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)

	fmt.Println(lesson_1.Equation(a, b, c))
}
