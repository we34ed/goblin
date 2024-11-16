package main

import (
	"fmt"
	"strings"
)

// import "fmt"
//
//	func grades(code map[string][]int) {
//		for key, value := range code {
//			counter := 0.0
//			for _, element := range value {
//				counter += float64(element)
//			}
//			fmt.Println(key, counter/float64(len(value)))
//		}
//	}
//
//	func main() {
//		studGrades := map[string][]int{
//			"Stas":  {2, 5, 5, 5, 5},
//			"Danil": {3, 5, 5, 5, 4},
//			"Sasha": {3, 5, 4, 4, 5},
//			"Vanya": {5, 5, 5, 5, 4},
//		}
//		grades(studGrades)
//	}
//

//		type Item interface {
//			Name() string
//			Damage() string
//		}
//
//		func (s Sword) Name() string {
//			return fmt.Sprintln("Название предмета:", s.name)
//		}
//
//		func (s Sword) Damage() string {
//			return fmt.Sprintln("Урон предмета", s.damage)
//		}
//
//		func (s Scroll) Name() string {
//			return fmt.Sprintln("Название предмета:", s.name)
//		}
//
//		func (s Scroll) Damage() string {
//			return fmt.Sprintln("Урон предмета", s.damage)
//		}
//
//		func info(i Item) {
//			fmt.Println("Вещь:", i.Name())
//			fmt.Println("Дамаг:", i.Damage())
//		}
//
//		func main() {
//			a := Sword{
//				"Рапира",
//				100,
//			}
//			info(a)
//			b := Scroll{
//				Sword{
//					"Свиток",
//					99,
//				},
//			}
//			info(b)
//		}
//
//		func handler(res http.ResponseWriter, req *http.Request) {
//			var out string
//			if out == `/time` || out == `/time/` {
//				out = time.Now().Format("02.01.2006 15:04:05")
//			} else {
//				fmt.Sprintf("Host:%s\nPath:%s\nmMethod:%s\n", req.Host, req.URL.Path, req.Method)
//			}
//			res.Write([]byte(out))
//		}
//
//		func main() {
//			fmt.Println("Запускаем сервер")
//			http.HandleFunc(`/`, handler)
//			err := http.ListenAndServe("localhost:8080", nil)
//			if err != nil {
//				panic(err)
//			}
//			fmt.Println("Работа окончена")
//		}
//
//		func main() {
//			response, err := http.Get("http://localhost:8080")
//			if err != nil {
//				fmt.Println("Ошибка", err)
//				return
//			}
//			body, err := io.ReadAll(response.Body)
//			response.Body.Close()
//			if err != nil {
//				fmt.Println("Ошибка", err)
//			}
//			fmt.Println(string(body))
//			fmt.Println(response.Status)
//		}
//
//		func main() {
//			response, err := http.Get("http://localhost:8080/time")
//			if err != nil {
//				fmt.Println("Ошибка", err)
//			}
//			body, err := io.ReadAll(response.Body)
//			response.Body.Close()
//			if err != nil {
//				fmt.Println("Ошибка", err)
//			}
//			fmt.Println(string(body))
//			fmt.Println(response.Status)
//	}
//func handlerr(res http.ResponseWriter, req *http.Request) { //http.ResponseWriter-интерфейс
//	//http.Request-обработчик запросов
//	var out string
//	if out == `/time` || out == `/time/` {
//		out = time.Now().Format("02.01.2006 15:04:05")
//	} else {
//		out = fmt.Sprintf("Host:%s\nPath:%s\nmMethod:%s\n", req.Host, req.URL.Path, req.Method)
//	} //req.Host-имя хоста,req.URL.Path-юрл путь,Req.method-метод
//	res.Write([]byte(out))
//}
//
//func main() {
//	fmt.Println("Запускаем сервер")
//	http.HandleFunc(`/`, handlerr) //todo: `/` - обработка запросов по любому URL-пути
//	err := http.ListenAndServe("localhost:8080", nil)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("Работа окончена")
//}

//	func main() {
//		response, err := http.Get("http://localhost:8080/time/")
//		if err != nil {
//			fmt.Println("Ошибка", err)
//		}
//		body, err := io.ReadAll(response.Body)
//		response.Body.Close()
//		if err != nil {
//			fmt.Println("Ошибка", err)
//		}
//		fmt.Println(string(body))
//		fmt.Println(response.Status)
//	}
//
//	func calculateStats(slice []int) {
//		if len(slice) == 0 {
//			fmt.Println("Слайс пустой(")
//			return
//		}
//		average := 0
//		minimum := 0
//		maximum := 0
//
//		for _, e := range slice {
//			average += e
//			if e <= minimum {
//				minimum = e
//			} else if e >= maximum {
//				maximum = e
//			}
//		}
//		fmt.Printf("Среднее значение: %.2f\n", float64(average)/float64(len(slice)))
//		fmt.Println("Минимальное число:", minimum)
//		fmt.Println("Максимальное число:", maximum)
//	}
//
//	func main() {
//		slicee := []int{24, 21, 1, 2, 3, 4, 5, 67, 54, 11, 0}
//		calculateStats(slicee)
//	}
func main() {
	stroka := "Stas"
	stroka2 := strings.Split(stroka, "")
	fmt.Println(stroka2)
}
