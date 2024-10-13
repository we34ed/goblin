package main

//func main() {
//pi := 3.1415
//fmt.Printf("число пи является: %.4f\n", pi)
//bros := 10
//fmt.Printf("%d братиков хотят депать.\n", bros)
//oneHundred := 100
//rubles := "рублей"
//friends := "друзей"
//fmt.Printf("Не имей %d %s, а имей %d %s.\n", oneHundred, rubles, oneHundred, friends)

// fmt.Sprintf(нужна для форматироваия return)
// fmt.Printf(нужна для форматирования чего-либо)
//func oneHundredFunc(one int, two, three string) string {
//	return fmt.Sprintf("Не имей %d %s, а имей %d %s\n", one, two, one, three)
//}
//func main() {
//	oneHundred := 100
//	rubles := "рублей"
//	friends := "друзей"
//
//	result := oneHundredFunc(oneHundred, rubles, friends)
//	fmt.Println(result)
//	fmt.Printf("Не имей %d %s, а имей %d %s", oneHundred, rubles, oneHundred, friends)
//}

//	func showMeteo(temp int, weat string) {
//		fmt.Printf("Сейчас %s,на градуснике %d\n", weat, temp)
//	}
//
//	func main() {
//		temperature := 21
//		weather := "погода"
//		showMeteo(temperature, weather)
//	}
//func main() {
//использование форматирования с мапами
//favouriteSongs := map[string]string{
//	"Тополиный пух": "Иванушки",
//	"Город золотой": "Аквариум",
//	"Перемен":       "Кино",
//}
//gruppa1 := favouriteSongs["Тополиный пух"]
//gruppa2 := favouriteSongs["Город золотой"]
//gruppa3 := favouriteSongs["Перемен"]
//fmt.Printf("из групп %s,%s и %s,самая молодая - группа %s\n", gruppa1, gruppa2, gruppa3, gruppa1)

//пример,как использовать сложение с числами
//number1 := 100
//number2 := 200
//fmt.Printf("сумма чисел %d и %d = %d", number1, number2, number1+number2)

//пример, как использовать форматирование со строками
//slovo1 := "Stas"
//slovo2 := " Kuchma"
//fmt.Printf("Мое имя %s,а фамилия%s,полное имя - %s", slovo1, slovo2, slovo1+slovo2)

//библиотеки (math.Sqrt-функция находящая квадртаный корень)
//mat := math.Sqrt(16)
//fmt.Println("квадратный корень из 16 -", mat)

//num := 16.0
//sqrt := math.Sqrt(num)
//fmt.Printf("вычислим квадратный корень из %.2f\n", num)
//fmt.Printf("Sqrt(%.f) = %.f", num, sqrt)

//функция math.Pow(возводит X в степень Y), дополнительные переменные не требуются,вызов функции сразу подставляется в аргумент
//fmt.Printf("2 в степени 3 = %.2f\n", math.Pow(2, 3))
//fmt.Printf("3 в степени 2 = %.2f\n", math.Pow(3, 2))

// квиз
//
//	func weather(town string, temp int, davl int) {
//		fmt.Printf("Город %s, температура %d по цельсию, %d мм.рт.ст.\n", town, temp, davl)
//	}
//func main() {
//	town := "Нижневартовск"
//	temperature := 21
//	davlenie := 5
//	weather(town, temperature, davlenie)
//}

//функция и новый тип данных Time
//var vremya time.Time
//fmt.Println(vremya)

//переменная с определенной датой и временем(Date(year int, month Month,day, hour, main, sec, nsec int, loc *location))
//start := time.Date(1961, 4, 12, 6, 7, 0, 0, time.UTC)
//fmt.Println("Уже", start, "поехали!")

//чтобы узнать настоящее время,нужно использовать функцию tine.Now(),она возвращает переменную типа time.Time,которая содержит время
//timeNow := time.Now()
//fmt.Println(timeNow)

//интервал между двумя моментами времени - метод второй.Sub(первый),применяется между значениями типа time.Time
//start := time.Date(1961, 4, 12, 6, 7, 0, 0, time.UTC)
//landing := time.Date(1961, 4, 12, 7, 55, 0, 0, time.UTC)
//duration := landing.Sub(start)
//fmt.Println(duration)
//fmt.Printf("Полет Гагарина проходил %v минут", int(duration.Minutes()))

//квиз научить Алису,сообщать,сколько шел любимый сериал пользователя(int(duration.Hours())-указывает время в часах)
//firstEpisode := time.Date(2011, 4, 17, 0, 0, 0, 0, time.UTC)
//lastEpisode := time.Date(2019, 4, 15, 0, 0, 0, 0, time.UTC)
//duration := lastEpisode.Sub(firstEpisode)
//fmt.Println(duration)
//fmt.Printf("сериал пользователя длился %d часов", int(duration.Hours()))

//квиз 2
//v120 := time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC)
//v121 := time.Date(2023, 8, 8, 0, 0, 0, 0, time.UTC)
//duration := v121.Sub(v120)
//fmt.Println(duration)
//fmt.Printf("переход от версии v120 к версии v121, занял %d часов", int(duration.Hours()))

//найти время нынешнее(now), время по миру(utc)
//today := time.Now()
//timeUTC := today.UTC()
//fmt.Println("время по Москве",today)
//fmt.Println("время по Миру", timeUTC)

//time.Duration(хранит интервал в единицах времени:часы,минуты и т.д.)
//moscow := time.Duration(3*time.Hour + 27*time.Minute + 55*time.Second)
//piter := time.Duration(12*time.Hour + 16*time.Minute + 55*time.Second)
//fmt.Println(moscow, piter)

//пример: вычисление Московского времени
//today := time.Now().UTC()
//moscow := today.Add(3 * time.Hour)
//fmt.Println("в Москве сейчас:", moscow)

//еще один пример с Duration(промежуток и тд)
//valteryBottas := time.Duration(1*time.Minute + 25*time.Second + 580*time.Millisecond)
//louisHamilton := valteryBottas + time.Duration(477*time.Millisecond)
//fmt.Printf("самый быстрый круг Хэмиллтона - %d секунд", int(louisHamilton.Seconds()))

// Форматирование времени(02-день,01-месяц,2006-год,3/15-час,4-минута,5-секунда
//func main() {
//	today := time.Now()
//	fmt.Println("время сейчас:", today.Format("02.01.2006 15:04:05"))
//	timeNow := time.Now()
//	fmt.Println("Время сейчас:", timeNow.Format("02.01.2006. 15.04.05."))
//}
//todo:знать!
