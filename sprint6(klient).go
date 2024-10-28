package main

import (
	"fmt"
	"io"
	"net/http"
)

//func main() {
//
//	//TODO: Клиент посылает запрос HTTP-серверу
//	//для этого достаточно функции "http.Get(url string)(resp *Response,err error)"
//	//она возвращает указатель на структуру типа "http.Response"(ответ сервера) и значение типа "error"(ошибку,если она появилась)
//
//	//TODO: Отправляем запрос нашему серверу(должен быть включен):
//
//	response, err := http.Get("http://localhost:8080/time") //отправляет GET запрос серверу и возвращает его ответ
//	if err != nil {
//		fmt.Println(err)
//		return
//	}                                                //тело ответа
//	body, err := io.ReadAll(response.Body) //функция "io.ReadAll(r Reader)([]byte,error)" она читает все данные и возвращает полученный слайс байтов
//	response.Body.Close() //функция,указывающая,что работа с данными закончена и можно освободить занимаемую память
//	if err != nil {
//		fmt.Println("Ошибка", err)
//		return
//	}
//	fmt.Println(string(body))
//	fmt.Println(response.Status) //получаем значение поля для проверки кода статуса с пояснением
//}                                //можно использоваать команду "response.StatusCode,для получения кода без пояснения

//// TODO: чтобы регулярно опрашивать сервер и получать текущее время,помещаем код с запросом в бесконечный цикл:
//func main() {
//	for {
//		response, err := http.Get("http://localhost:8080/time")
//		if err != nil {
//			fmt.Println("Ошибка", err)
//			break
//		}
//		body, err := io.ReadAll(response.Body)
//		response.Body.Close()
//		if err != nil {
//			fmt.Println("Ошибка", err)
//			break
//		}
//		fmt.Println(string(body))
//		fmt.Println(response.Status)
//		time.Sleep(5 * time.Second)
//	}
//} //чтобы завершить работу,нажимаем "control C"

func main() {
	response, err := http.Get("http://localhost:8080/time")
	if err != nil {
		fmt.Println("Ошибка", err)
		return
	}
	body, err := io.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		fmt.Println("Ошибка", err)
	}
	fmt.Println(string(body))
	fmt.Println(response.Status)
}
