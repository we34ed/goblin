package main

//база!
//func main() {
//	rand.Seed(time.Now().UnixNano()) //генератор случайных чисел
//	a := rand.Intn(3)
//	fmt.Println(1 / a)
//}

//var ErrZeroDivider = errors.New("делитель не должен быть равен нулю") //создаем ошибку статически
//
//func simpleDivider(divisible, divider int) (int, error) {
//	result := 0
//	if divider == 0 { // проверяем равен ли делитель нулю
//		return result, ErrZeroDivider //возвращаем ошибку
//	}
//	result = divisible / divider
//	return result,nil
//}
//func main() {
//	rand.Seed(time.Now().UnixNano())
//	a := rand.Intn(3) //случайное число от 0 до 3
//	divResult, err := simpleDivider(2, a)
//	if err != nil { //проверяем,нет ли ошибки в ходе выполнения функции
//		err := fmt.Errorf("ошибка в ходе выполнения программы: %v", err) //формируем сообщение с проблемой
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(divResult)
//}

//func main() {
//	rand.Seed(time.Now().UnixNano())
//	a := rand.Intn(3)
//	if a == 0 {
//		fmt.Println("стоп")
//		log.Fatal(a)
//	}
//	fmt.Println(0 / a)
//}
