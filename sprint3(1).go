package main

import "fmt"

//func main() {
//int и float

//	x := 1000 % 3 //оператор модуло % (подкапотно делит число,однако возвращает только остаток)
//	fmt.Println(x, 100%3, 101%2)
//}
//x := 10.5
//y := 2
//x += float64(y)//присваеваем флоат к инту
//fmt.Println(x)
//}

//a := 2.345
//b := 2.//сокращенный синтаксис
//c := -2.0
//d := .07//сокращенный синтаксис
//e := 0.07
//fmt.Println(a, "!", b, "!", c, "!", d, "!", e)
//}

//	x := 0.1
//	y := 0.2
//	z := x*10 + y*10 //пример,как избавиться от погрешности(сначала умножаем на определенное число)
//	z /= 10          //(потом на это же число делим)
//	fmt.Printf("%f", z)
//
//	fmt.Println(z)
//}

//квиз
//weekDistanceCount := 0.0
//weekKm := []float64{2.32, 1.56, 7.56, 8.91, 3.65, 4.33, 6.57}
//for _, sumWeek := range weekKm {
//	weekDistanceCount += sumWeek
//}
//fmt.Printf("%.2f", weekDistanceCount)
//}

//условные операторы
//	rating := 0.0
//	if rating > 0 {
//		if rating > 4.7 {
//			fmt.Println("фильм крут")
//		} else if rating > 3.5 && rating < 4.7 {
//			fmt.Println("смотреть можно")
//		} else {
//			fmt.Println("так себе фильмец")
//		}
//		fmt.Println("проверка окончена")
//	}
//}

//bridge1 := "сведен"
//bridge2 := "разведен"
//if bridge1 == "сведен" && bridge2 != "сведен" {
//	fmt.Println("программист доедет до оффера")
//} else {
//	fmt.Println("программист продолжает собеситься")
//}

//numbers := []int{2, -4, 5, 0, 16, 44, 0, -1, 32, 4, 0}
//for _, num := range numbers {
//	if num > 0 {
//		fmt.Println("число положительное")
//	} else if num == 0 {
//		fmt.Println("число равняетсся нулю")
//	} else {
//		fmt.Println("число отрицательное")
//	}
//}

//конструкция switch case
//answer := "не знаю"
//switch answer {
//case "да":
//	fmt.Println("правильно")
//case "нет":
//	fmt.Println("не правильно")
//case "не знаю":
//	fmt.Println("думай")
//}

//		rating := 0.0
//		switch {
//		case rating > 4.7:
//			fmt.Println("фильм крут")
//		case rating> 3.5 && rating < 4.7:
//			fmt.Println("смотреть можно")
//		default:
//			fmt.Println("так себе фильмец")
//
//			fmt.Println("проверка окончена")
//		}
//	}
//func printCounter(count int) {
//	switch count {
//	case 0:
//		fmt.Println("0")
//		fallthrough //эта функция позволяет спуститься вниз по case, не важно выполняется ли код
//	case 1:
//		fmt.Println("1")
//		fallthrough
//	case 2:
//		fmt.Println("2")
//		fallthrough
//	case 3:
//		fmt.Println("3")
//	}
//	fmt.Println("")
//}
//func main() {
//	for i := 0; i < 4; i++ {
//		printCounter(i)
//	}slice:=
//}

//	func printWind(beafort int) {
//		var comment string
//		switch beafort {
//		case 0:
//			comment = "штиль"
//		case 1, 2, 3:
//			comment = "слабый ветер"
//		case 4, 5:
//			comment = "умеренный ветер"
//		case 6, 7, 8:
//			comment = "сильный ветер"
//		case 9, 10, 11:
//			comment = "шторм"
//		case 12:
//			comment = "ураган"
//		}
//		fmt.Println(beafort, comment)
//	}
//
//	func main() {
//		for i := 0; i < 12; i++ {
//			printWind(i)
//		}
//	}
//func main() {
//	september := []int{-1, 0, 1, 1, -1, -1, 0, 0, 0, -1, 1, 1, 1, 0, 0, 0, 1, 1, -1, 0, -1, 1, 0, 0, 1, 1, -1, -1, 0, 0, 0}
//	sun := 0
//	rain := 0
//	clouds := 0
//	for _, s := range september {
//		switch s {
//		case 1:
//			sun++ //счетчик дней
//		case 0:
//			clouds++ //счетчик дней
//		case -1:
//			rain++ //счетчик дней
//		}
//	}
//	fmt.Printf("в Сентябре было %d солнечных, %d облачных и %d дождливых деньков", sun, clouds, rain)
//}

//циклы

//func main() {
//	movies := map[string]float64{
//		"Терминатор": 8.0,
//		"Матрица":    8.5,
//		"Сеть":       7.1,
//		"Хакеры":     6.9,
//	}
//	sum := 0.0
//	for _, m := range movies {
//		sum += m / 4
//	}
//	fmt.Printf("средний рейтинг фильмов: %.1f", sum)
//}

// }
//
//	func main() {
//		movie := "Матрица"
//		for _, m := range movie {
//			fmt.Printf("| %c\n", m)
//			fmt.Printf("--\n")
//		}
//	}
//
// func main() { //пример,как можно применять цикл
//
//		recomendMovies := []string{"Хатико", "23", "Достучаться до небес", "Хакеры", "Трон", "1408"}
//		hackerMovies := []string{"Трон", "Военные игры", "Тихушники", "Хакеры", "Нирвана", "Сеть"}
//		for _, r := range recomendMovies {
//			for _, h := range hackerMovies {
//				if r == h {
//					fmt.Printf("Разработчикам рекомендуем посмотреть фильм: %s\n", r)
//				}
//			}
//		}
//	}
//
//	func main() {        //просто цикл
//		count := 0
//		for i := 0; i < 1000; i++ {
//			if i%3 == 0 {
//				count++
//			}
//		}
//		fmt.Println(count)
//	}
//func main() {         //цикл задом наперед
//	quiz := []string{"фазан", "сидит", "где", "знать", "желает", "охотник", "каждый"}
//	for i := len(quiz) - 1; i >= 0; i-- {
//		fmt.Println(quiz[i])
//	}
//}

//func main() {     //пример break
//	top:
//	for i := 0; i < 100; i++ {
//		for j := 0; j < i; j++ {
//			if i*j > 500 {
//				break top
//			}
//		}
//	}
//}

//	func main() {
//		friends := []string{"sasha", "slava", " ", "vanya", "danya", "stasya"}
//		for _, i := range friends {
//			if i == " " {   //continue продолжает перебор элементов,однако забирает выделенный элемент
//				continue
//			}
//			if i == "danya" {
//				break      //break завершает перебор элементов,забирая выделенный элемент
//			}
//			fmt.Println(i)
//		}
//	}
//
//	func main() {
//		count := 0.0
//		slice := []float64{2.1, 4.3, 3.2, 5.5, 5.1, 4.4, 3.5}
//		for _, s := range slice {
//			count += s / 7
//		}
//		fmt.Printf("средний балл студента: %.1f", count)
//	}
//
//	func main() {
//		bestFriends := []string{"Sasha", "Vanya", "Danya", "Slavik", "Artem", "Egor"}
//		friendsDoters := []string{"Sasha", "Vanya", "Slavik"}
//		for _, b := range bestFriends {
//			for _, f := range friendsDoters {
//				if b == f {
//					fmt.Printf("лучшие друзья играющие в доту: %s\n", b)
//				}
//			}
//		}
//func studsGrades(stud []float64) float64 {
//	count := 0.0
//	for _, grades := range stud {
//		count += grades / 4
//
//	}
//	return count
//}
//func main() {
//	slice := []float64{4.5, 5.3, 3.1, 4.4}
//	fmt.Println(studsGrades(slice))
//}

//func sredball(grades []int) int {
//	count := 0
//	for _, grade := range grades {
//		count += grade
//	}
//	return count / len(grades)
//}
//func main() {
//	studGrades := []int{2, 5, 4, 3, 4, 5}
//	result := sredball(studGrades)
//	fmt.Println(result)
//	fmt.Printf("Средний балл: %d", result)"
//}

// строки - они имутабельны(неизменяемые)
//func main() {              //юникодовые значения:
//	eng := "hello world"    //для ASCII всегда меньше 128
//	ru := "привет мир"    //для кириллицы всегда больше 1000
//	fmt.Println(len(eng))
//	fmt.Println(len(ru))
//	fmt.Println([]byte(eng))
//	fmt.Println([]byte(ru))
//
//}

//func main() {
//	eng := "hello world"
//	eng1 := []rune(eng)
//	eng1[0] = 'H'
//	eng2 := string(eng1)
//	fmt.Println(eng, eng1, eng2)
//}

//func main() {//helloBytes=UTF8, helloRunes=Unicode
//	helloBytes := []byte{208, 159, 209, 128, 208, 184, 208, 178, 208, 181, 209, 130, 44, 32, 208, 188, 208, 184, 209, 128, 33}
//	helloRunes := []rune{1055, 1088, 1080, 1074, 1077, 1090, 44, 32, 1084, 1080, 1088, 33}
//	fmt.Println("Hello Bytes!", string(helloBytes))
//	fmt.Println("Hello Runes!", string(helloRunes))
//}

//func main() {
//	alert := "потребление каллорий вредно пиздец " + "выкинь нахуй свой бургер " + "урод"
//	fmt.Println(alert)
//
//}

//func main() {
//	s := "hello world!"
//	substr := s[:5]     //срез строки (:5 - выделяет кол-во символов для среза)
//	fmt.Println(substr)
//}

//func main() {
//	stroka := "стася чемпион"
//	fmt.Println(strings.ToUpper(stroka)) //функция ToUpper из пакета strings переводит символы в верхний регистр
//	stroka2 := "СТАСЯ ЛУЧШИЙ"            //функция ToLower из пакета strings переводит символы в нижний регистр
//	fmt.Println(strings.ToLower(stroka2))
//}

//func main() {
//	stroka := "Саня,привет!"
//	if strings.HasPrefix(stroka, "Саня") {     //HasPrefix-начало строки
//		fmt.Println("Запрос к Сане получен")
//	}
//	if strings.HasSuffix(stroka, "!") {        //HasSuffix-конец строки
//		fmt.Println("good")
//	}
//}

//	func main() {
//		stroka := "Саня, как найти работу?"
//		if strings.HasPrefix(stroka, "Саня") {
//			fmt.Println("обращение к Сане выполнено успешно")
//		}
//		if strings.HasSuffix(stroka, "?") {
//			fmt.Println("успешно!")
//		}
//	}
//func main() {
//	msg := "Сегодня вы сожгли 500 каллорий"                        //strings.ReplaceAll-функция заменяющая old на new
//	cats := "Котики важны.Котики улучшают эмоциональное здоровье." //strings.Index -функция возвращает позицию первого вхождения(если не находит возвращает -1),считает в байтах
//
//	workout := strings.ReplaceAll(msg, "вы", "мы")
//
//	fmt.Println(workout)
//	fmt.Println(strings.Index(cats, "важны"))
//	fmt.Println(strings.ReplaceAll(cats, "Котики", "Тренировки"))
//
//	fuckingAwesome := "Пацаны,кто сегодня зарплату получил?"
//	fmt.Println(strings.ReplaceAll(fuckingAwesome, "Пацаны", "Братики"))
//	fmt.Println(strings.Index(fuckingAwesome, "кто"))
//}

//func main() {
//	cats := "\r\n Алиса,го гулять?\n"
//	fmt.Println(strings.TrimSpace(cats)) //функция TrimSpace удаляет все ненужное из строк
//
//	cats2 := "Алиса, привет!"
//	fmt.Println(strings.TrimPrefix(cats2, "Алиса, ")) //удаляет указанную строку
//
//	cats3 := "Саня, пошли шмали надуемся!"
//	fmt.Println(strings.TrimPrefix(cats3, "Саня, ")) //удаляет указанную строку
//}

//func main() {
//	stroka := "привет, как дела?"
//	strokaSplit := strings.Split(stroka, " ")       //из строки создает слайс
//	fmt.Println(strokaSplit)
//	stroka2 := []string{"stroka1", "stroka2", "stroka3", "stroka4", "stroka5"}
//	strokaJoin := strings.Join(stroka2, ",")        //из слайса создает строку
//	fmt.Println(strokaJoin)
//}

// quizz
//func main() {
//	weight := 75.0
//	height := 175.0
//	steps := 8462.0
//	hours := 2.0
//	lenStep := 0.65
//
//	dist := (steps * lenStep) / 1000
//	meanSpeed := dist / hours
//	minutes := hours * 60
//	spentCallories := (0.035*weight + (meanSpeed*meanSpeed/height)*0.029*weight) * minutes
//	fmt.Printf("Сегодня вы прошли %.1f км. и затратили %.1f калл.", dist, spentCallories)
//}

//func main() {
//	weight := 75.0
//	height := 175.0
//	steps := 8462.0 //условия
//	hours := 2.0
//	lenStep := 0.65
//
//	dist := (steps * lenStep) / 1000 //вычислил дистанцию
//	var achive string                //создал пустую переменную типа string
//	if dist >= 6 {
//		achive = "Отличный результат!"
//	} else if dist >= 3.9 { //применил условие,отслеживающее насколько хорош результат
//		achive = "Хороший результат!"
//	} else if dist >= 2.0 {
//		achive = "Маловато,но правильно что начал!"
//	} else {
//		achive = "мдэуу"
//	}
//	meanSpeed := dist / hours //вычислил км/ч
//
//	minutes := hours * 60 //вычислил затраченные минуты
//
//	spentCallories := (0.035*weight + (meanSpeed*meanSpeed/height)*0.029*weight) * minutes //формула расчета каллорий
//
//	fmt.Printf("Сегодня вы прошли %.1f км. и затратили %.1f калл.\n %s", dist, spentCallories, achive)
//
//}

//func main() {
//	brandname := "nike"
//	fmt.Println("адрес переменной brandname:", &brandname, "значение", brandname) //амперсанд & указывает на адрес переменной(фактический адрес)
//	brandname = "cat"
//	fmt.Println("адрес переменной brandname:", &brandname, "значение", brandname)
//}//адрес переменной останется тем же,однако прошлое значение переменной будет удалено,так как на него никто не ссылается
//если на ячейку памяти есть хотя бы одна ссылка(переменная,то она удаляться не будет)

//	func main() {
//		brand := "Nike"
//		fmt.Println("Адрес переменной brand:", &brand)//благодаря & нашли адрес переменной brand
//
//		brandAdress := &brand
//		fmt.Println("Значение переменной brand:", *brandAdress)//благодаря* обратились через адрес к значению
//	}
//
//	func studgrades(studs []float64) float64 {
//		counter := 0.0
//		for _, value := range studs {
//			counter += value
//		}
//		return counter / float64(len(studs))
//	}
//
//	func main() {
//		studsgrades := []float64{2, 5, 5, 5, 4, 3, 2, 4, 5}
//		result := studgrades(studsgrades)
//		fmt.Printf("средний балл ученика:%.3f", result)
//	}
//
//	func main() {
//		responce, err := http.Get("https://yandex.ru")
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		body, err := io.ReadAll(responce.Body)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		result := string(body)
//		fmt.Println(result)
//	}
//
// func main() {
// name := "Stas"
// fmt.Println(&name)
// nameAdres := &name
// fmt.Println(*nameAdres)
// fmt.Println(strings.ReplaceAll(name, "Stas", "Stanislav"))
// stroka := "stas kuchma"
// strokaSplit := strings.Split(stroka, " ")
// fmt.Println(strokaSplit)
// slice := []string{"stas", "kuchma", "boss"}
// sliceJoin := strings.Join(slice, ",")
// fmt.Println(sliceJoin)
// gle := "stanislav kuchma"
// gleAmg := []rune(gle)
// gleAmg[0] = 'S'
// gleAmg[10] = 'K'
// gleAmg2 := string(gleAmg)
// fmt.Println(gleAmg2)
//
//	for i := 0; i < 1000; i++ {
//		if i%2 == 0 {
//			fmt.Println(i)
//		}
//	}
//
//	for i := 1000; i >= 0; i-- {
//		if i%2 == 0 {
//			fmt.Println(i)
//		}
//	}
//func main() {
//	mapp := map[string][]float64{
//		"stas":   {2, 5, 4, 3, 5},
//		"den":    {3, 5, 4, 5, 4},
//		"vanya":  {4, 3, 4, 5, 5}, //TODO:повторить
//		"sasha":  {3, 4, 5, 3, 2},
//		"slavik": {2, 3, 3, 4, 5},
//	}
//	for _, value := range mapp {
//		counter := 0.0
//		for _, element := range value {
//			counter += element
//		}
//		fmt.Println(counter / float64(len(mapp)))
//	}
//
//}

//	func main() {
//		responce, err := http.Get("https://yandex.ru")
//		if err != nil {
//			fmt.Printf("вышла ошибка c соединением %v", err) //TODO:повторить %w
//			return
//		}
//		body, err := io.ReadAll(responce.Body)
//		if err != nil {
//			fmt.Printf("вышла ошибка с получением %v", err)
//			return
//		}
//		//defer close(responce)          //TODO:научиться закрывать тело ответа
//		result := string(body)
//		fmt.Println(result)
//	}
//func main() {
//	mapp := map[string][]float64{
//		"vanya":  {5, 4, 4},
//		"stasya": {4, 5, 3},
//		"danya":  {5, 5, 4},
//	}
//	for _, value := range mapp {
//		counter := 0.0
//		for _, element := range value {
//			counter += element
//		}
//		fmt.Println(counter / float64(len(mapp)))
//	}
//}

// todo:константы
//
//	func main() {
//		const pi = 3.1415          //константа- как и переменная,способ хранения данных,только ее нельзя изменять
//		                           //константа- неизменяемое значение
//	}
//func main() {
//	const (
//		a = 2
//		b              //если указать значение только одной из констант,все последующие присвояет это же значение
//		c
//		d
//		e = 3
//		f
//	)
//	fmt.Println(a, b, c, d, e, f)
//}

//func main(){
//	s:=5
//	const n=s        //константе нельзя присвоить переменную(нельзя инициализировать переменной)
//}

//func main() {
//	const s = "Стас" //нетипизированная(ей не присвоен явный тип) константа,так же есть типизированные константы
//	const t = s
//	fmt.Println(t)
//}

//func main() {
//	const (
//		one       = 1
//		oneFloat  = 1.000
//		oneE      = 1e3 - 99.0*10 - 9
//		oneBase16 = '\x01'
//	)
//	f := one
//	i := oneFloat
//	e := oneE
//	d := oneBase16
//	fmt.Println(f, i, e, d)
//}

//TODO:iota
//инструмент для создания нескольких констант,когда каждая последующая увеличивается на 1

// const (
//
//	stas = iota + 1           //если нужно изменить изначальное число,прибавьте загаданное числа к iota
//	den
//	vanya
//	slavik
//	sasha
//	egor
//	semen
//
// )
//
//	func main() {
//		fmt.Printf("братик %v\n", stas)
//		fmt.Printf("братик %v\n", den)
//		fmt.Printf("братик %v\n", vanya)
//		fmt.Printf("братик %v\n", slavik)
//		fmt.Printf("братик %v\n", sasha)
//		fmt.Printf("братик %v\n", egor)
//		fmt.Printf("братик %v\n", semen)
//	}
//
// квиз
// const (
//
//	ok = iota
//	canselled
//	fail
//
// )
//
//	func codeWars(code int) string {
//		if code == ok {
//			return "status ok"
//		} else if code == canselled {
//			return "status canselled"
//		} else if code == fail {
//			return "status fail"
//		} else {
//			return "unknown status"
//		}
//	}
//
//	func main() {
//		fmt.Println(codeWars(0))
//	}
//
// тот же квиз но с помощью switch
// const (
//
//	ok = iota
//	cancelled
//	fail
//
// )
//
//	func codeW(code int) string {
//		switch code {
//		case ok:
//			return "status ok"
//		case cancelled:
//			return "status cancelled"
//		case fail:
//			return "status fail"
//		default:
//
//			return "undefined status"
//		}
//	}
//
//	func main() {
//		fmt.Println(codeW(0))
//	}

// TODO:последовательности
//
//	func main() {
//		numbers := []int{1, 2, 3, 4, 5}
//		fmt.Println("исходный слайс", numbers)
//		numbers = append(numbers, 6)
//		fmt.Println("в слайс добавлен элемент", numbers)
//		numbers = append(numbers, 7, 8, 9)
//		fmt.Println("в слайс добавлено несколько элементов", numbers)
//		fmt.Println("в слайс можно добавить элемент и сразу напечатать его", append(numbers, 10, 11, 12))
//	}
//
//	func main() {
//		movies := [4]string{"Матрица", "Хакеры", "Трон", "Джонни Мнемоник"}
//		fmt.Println(movies[1])
//		movies[2] = "Терминатор"
//		fmt.Println(movies)
//	}
//
//	func main() {
//		movie := []rune("Джони Мнемоник")
//		fmt.Printf("%c%c%c%c%c", movie[0], movie[2], movie[9], movie[12], movie[13])
//
// }
// квиз
//
//	func add(src []int, i int) []int {
//		return append(src, i)
//	}
//
//	func main() {
//		numbers := []int{77, 88, 99}
//		numbers = append(numbers, 100)
//		fmt.Println(add(numbers, 101))
//	}
//
// TODO:слияние слайсов(в append добавить "..."
//
//	func main() {
//		slice1 := []string{"пацаны,", "я", "пошел"}
//		slice2 := []string{"играть", "в", "доту"}
//		slice1 = append(slice1, slice2...)
//		fmt.Println(slice1)
//	}
//
// TODO:Срезы(обращаемся по индексу от:до,однако последний индекс не входит в счет)
//
//	func main() {
//		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//		numbers1 := numbers[0:10]
//		fmt.Println(numbers1)
//		numbers2 := numbers[:]
//		fmt.Println(numbers2)
//		numbers3 := numbers[:9]
//		fmt.Println(numbers3)
//		numbers4 := numbers[3:]
//		fmt.Println(numbers4)
//	}
//func main() {
//	slice1 := []int{1, 2, 3, 4, 5}
//	slice2 := slice1[0:3]
//	slice3 := make([]int, 5, 5)
//	slice4 := append(slice1, slice2...)
//	fmt.Println(slice1)
//	fmt.Println(slice2)
//	fmt.Println(slice3)
//	fmt.Println(slice4)
//}
//квиз

//	func main() {
//		words := []string{"Терры", "важны", "в", "Роботы", "стали", "период", "эмиграции", "с"}
//		sentence := make([]string, 0, 8)
//		sentence = append(sentence, words[3:5]...)
//		sentence = append(sentence, words[1:3]...)
//		sentence = append(sentence, words[5:8]...)
//		sentence = append(sentence, words[0])
//		fmt.Println(sentence)
//	}
//
// const (
//
//	HelloInEnglish = "Hello!"
//	HelloInRussian = "Привет!"
//	HelloInSpanish = "¡Hola!"
//
// )
//
//	func main() {
//		choice := 0
//
//		fmt.Println("Выберите язык:")
//		fmt.Println("1 - Английский")
//		fmt.Println("2 - Русский")
//		fmt.Println("3 - Испанский")
//		fmt.Scan(&choice)
//
//		switch choice {
//		case 1:
//			fmt.Println(HelloInEnglish)
//		case 2:
//			fmt.Println(HelloInRussian)
//		case 3:
//			fmt.Println(HelloInSpanish)
//		default:
//			fmt.Println("Неверный выбор")
//		}
//	}
//
// todo:конец
//
//	func slovo(code []int) int {
//		counter := 0
//		for _, element := range code {
//			counter += element
//
//		}
//		return counter
//	}
//
//	func main() {
//		slice := []int{1, 2, 3, 4, 5}
//		result := slovo(slice)
//		fmt.Println(result)
//	}
//
//	func main() {
//		slice := []int{1, 2, 15, 4, -3, 0}
//		maximum := 0
//		for _, element := range slice {
//			if element > maximum {
//				maximum = element
//			}
//		}
//		fmt.Println(maximum)
//	}
//
// //todo:update
func grade(grades map[string][]int) {
	for key, value := range grades {
		counter := 0
		for _, element := range value {
			counter += element
		}
		fmt.Println(key, float64(counter)/float64(len(value)))
	}

}
func main() {
	gradesStud := map[string][]int{
		"stas":  {5, 4, 4, 3, 5},
		"vanya": {4, 4, 4, 3, 5},
		"danya": {4, 5, 5, 5, 4},
	}
	grade(gradesStud)
}
