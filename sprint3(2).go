package main

// )
// //todo:типизация
//
//	func main() {
//		oneInt := 1
//		twoString := "2"
//		threeFloat := 3.0
//		stringVInt, err := strconv.Atoi(twoString)
//		if err != nil {
//			panic(err)
//		}
//		fmt.Printf("сумма чисел %d,%d и %.f равна %d", oneInt, stringVInt, threeFloat, oneInt+stringVInt+int(threeFloat))
//	}
//
// //todo:функции
//func hello(a string, b string) {                  //обрабатывают данные
//	fmt.Println(a + ",привет бро!" + b)           //возвращают значение(результат)
//}
//func main() {
//	name := "Стас"
//	backend := "как дела?"
//	hello(name, backend)
//}

// квиз(функция выводящая большее число
//func max(a, b int) {
//	if a > b {
//		fmt.Println("большее число:", a)
//	} else if b > a {
//		fmt.Println("большее число:", b)
//	}
//}
//func main() {
//	max(3, 2)
//}

// произвольное количество аргументов("..." после до присваивания типа)
//func hello(a ...int) {
//	for _, s := range a {
//		fmt.Println(s)
//	}
//}
//func main() {           //стоит понимать,что такой синтаксис позволяет работать с агументами только ОДНОГО типа
//	hello(1, 2, 3, 4, 5)
//	hello(111)
//	hello()
//}

// квиз
//
//	func maximumNum(numbers ...int) int {
//		if len(numbers) == 0 {
//			return 0
//		} else {
//			max := 0
//			for _, num := range numbers {
//				if num > max {
//					max = num
//				}
//			}
//			return max
//		}
//	}
//
//	func main() {
//		maximum := maximumNum(10, 11, -1, 15, 0)
//		fmt.Println(maximum)
//	}
//func maximumNum(num ...int) int {
//	if len(num) == 0 {
//		return 0
//	}
//	max := 0
//	for _, number := range num {
//		if number > max {
//			max = number
//		}
//	}
//	return max
//}
//func main() {
//	result := maximumNum(1, 2, 3, 10)
//	fmt.Println(result)
//}

//	func minimum(num ...int) int {
//		if len(num) == 0 {
//			return 0
//		}
//		min := 100
//		for _, m := range num {
//			if m <= min {
//				min = m
//			}
//		}
//		return min
//	}
//
//	func main() {
//		result := minimum(11, 10, 8, -3, -1000)
//		fmt.Println(result)
//	}
//
// todo:анонимная функция(ее можно создать в переменной и вызвать в функции)
//func main() {
//	result := func(a, b int) bool { return a < b }(10, 20)
//	fmt.Println(result)
//}

// так же ее можно присвоить переменной
//func main() {
//	f := func(a, b int) bool { return a < b }
//	result := f(10, 20)
//	fmt.Println(result)
//}

// так же анонимную функцию можно использовать в функции

//	func printOnTrue(f func(a int, b int) bool, s string) {
//		if f(10, 20) {
//			fmt.Println(s)
//		}
//	}
//
//	func main() {
//		f := func(a, b int) bool { return a < b }
//		printOnTrue(f, "и правда меньше")
//	}
//
// todo:конец

//	func main() {
//		timing1 := time.Date(1978, 05, 14, 00, 00, 00, 00, time.UTC)
//		timing2 := time.Date(2024, 10, 12, 9, 43, 00, 00, time.UTC)
//		duration := timing2.Sub(timing1)
//		fmt.Printf("я живу %.1f дней", duration.Hours()/24)
//	}

// функция
//
//	func codew(code map[string][]int) {
//		for key, value := range code {
//			counter := 0.0
//			for _, element := range value {
//				counter += float64(element)
//			}
//			fmt.Println(key, counter/float64(len(value)))
//		}
//	}
//
// // мэйн
//
//	func main() {
//		mapp := map[string][]int{
//			"stas":  {2, 5, 5, 5, 4},
//			"vanya": {3, 4, 4, 3, 5},
//			"danya": {4, 4, 4, 5, 5},
//			"sanya": {2, 5, 5, 5, 4},
//		}
//		codew(mapp)
//	}

//	func main() {
//		counter := 0
//		for i := 0; i < 1000; i++ {
//			if i%2 == 0 {
//				counter += 1
//			}
//		}
//		fmt.Println("Количество чисел, кратных 3 от 0 до 999:", counter)
//	}
//
//	func code(num []int) int {
//		minimum := len(num)
//		for _, element := range num {
//			if element <= minimum {
//				minimum = element
//			}
//		}
//		return minimum
//	}
//
//	func main() {
//		slice1 := []int{-1, 2, 3, 0, 5}
//		slice2 := []int{0, -13, 15, 1, 5}
//		fmt.Println(code(slice1))
//		fmt.Println(code(slice2))
//	}
