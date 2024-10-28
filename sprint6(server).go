//TODO: Создание простого сервера

package main

import (
	"fmt"
	"net/http"
	"time"
)

//func main(){
//	fmt.Println("Запускаем сервер")
//	err:=http.ListenAndServe("localhost:8080",nil)
//	if err !=nil{
//		panic(err)
//	}
//	fmt.Println("Завершаем работу")
//}

//todo: listenAndServe(addr:)
//параметр addr:IP-адрес компьютера(:) и номер порта(8080)(чтобы сервера/сервисы не мешали друг другу)
//каждый порт слушает только один сервер,указывая порт можно однозначно связаться с нужным сервисом
//"addr" записывается в формате "IP-адрес:порт",например 127.0.0.1:(IP-адрес)8080(порт)
//вместе цифр можно указать "localhost",по умолчанию направляет на этот IP-адрес(127.0.0.1)
//если у компьютера несколько IP-адресов, можно написать ":порт"

//todo: listenAndServe(handler:)
//пока скипаем

//todo: Проверка работы
//пишем в браузере "http://localhost:8080"
//прекратить работу сервера можно командой в консоли: "Ctrl-C"

//todo: Обработчик HTTP-Запросов(handler)
//для добавления handler к созданному серверу,нужно создать функцию, которая будет обрабатывать входящие запросы
//http.HandleFunc(pattern string(`/`),handler func(имя http.ResponseWriter,имя *http.Request)(mainHandle)

//func handle(res http.ResponseWriter, req *http.Request) {
//	fmt.Println("Обработка запроса")
//}
//func main() {
//	fmt.Println("Запускаем сервер")
//	http.HandleFunc(`/`, handle)
//	err := http.ListenAndServe("localhost:8080", nil)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("Работа окончена")
//}
//todo: http.ResponseWriter(интерфейс потоковой записи,содержащий три функции

//type ResposeWriter interface{
//	Header()Header
//	Write([]byte)(int,error)
//	WriteHeader(statusCode int)
//}
//методы Header() и WriteHeader() используются для работы с HTTP-заголовками, а метод Write() выводит тело ответа

//выводим на страницу текущую дату и время методом Write()(вывод тела ответа)
//func header(res http.ResponseWriter,req http.Request){
//	s:=time.Now().Format("02,01,2006 15:04:05")
//	res.Write([]byte(s))
//}

//таким образом,мы создали простейший веб-сервер,показывающий текущее время,имеющий один обработчик,обрабатывающий любой URL-путь в браузере
//также в обработчике можно получить имя хоста с URL-путем и выполнять различные действия,в зависимости от запроса
//для этого обращаемся к полям структуры "http.Request"

//todo: http.Request

//Host-содержит имя хоста
//URL.Path-содержит URL-путь
//Method-имя HTTP-метода запроса

//если переделать обработчик так:
//func handler(res http.ResponseWriter, req *http.Request){
//	s:=fmt.Sprintf("Host: %s\nPath: %s/n,Method: %s",req.Host,req.URL.Path,req.Method)
//	res.Write([]byte(s))
//}
//то при открытии в браузере "http://localhost:8080/api" выведет:
//Host: localhost:8080
//Path: /api
//Method: Get

//func handler(res http.ResponseWriter, req *http.Request) {
//	a := fmt.Sprintf("Host:%s\nPath:%s\nMethod:%s\n", req.Host, req.URL.Path, req.Method)
//	res.Write([]byte(a))
//}
//func main() {
//	fmt.Println("Запускаем сервер")
//	http.HandleFunc(`/`, handler)
//	err := http.ListenAndServe("localhost:8080", nil)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("Работа окончена")
//}

//TODO: Объединяем два хендлера: текущее время и информацию о хосте
//объединяем таким образом, чтобы при запросе /time или /time/ показывалось текущее время
//для всех остальнвх запросов - информация о хосте и URL-пути

//	func handler(res http.ResponseWriter, req *http.Request) {
//		var out string
//		if req.URL.Path == `/time` || req.URL.Path == `/time/` {
//			out = time.Now().Format("02.01.2006 15:04:05")
//		} else {
//			out = fmt.Sprintf("Host:%s\nPath:%s\nMethod:%s\n", req.Host, req.URL.Path, req.Method)
//		}
//		res.Write([]byte(out))
//	}
//
//	func main() {
//		fmt.Println("Запускаем сервер")
//		http.HandleFunc(`/`, handler)
//		err := http.ListenAndServe("localhost:8080",nil) //http.ListenAndServe-запуск сервера
//		if err != nil {
//			panic(err)
//		}
//		fmt.Println("Работа окончена")
//	}
//
//	func handler(res http.ResponseWriter, req *http.Request) {
//		var out string
//		if req.URL.Path == `/time` || req.URL.Path == `/time/` {
//			out = time.Now().Format("02.01.2006 15:04:05")
//		} else {
//			out = fmt.Sprintf("Host:%s\nPath:%s\nMethod:%s", req.Host, req.URL.Path, req.Method)
//		}
//		res.Write([]byte(out))
//	}
//
//	func main() {
//		fmt.Println("Запускаем сервер")
//		http.HandleFunc(`/`, handler)
//		err := http.ListenAndServe("localhost:8080", nil) //http.ListenAndServe-запуск сервера
//		if err != nil {
//			panic(err)
//		}
//		fmt.Println("Работа окончена")
//	}
func handler(res http.ResponseWriter, req *http.Request) {
	var out string
	if req.URL.Path == `/time` || req.URL.Path == `/time/` {
		out = time.Now().Format("02.01.2006 15:04:05")
	} else {
		out = fmt.Sprintf("Host:%s\nPath:%s\nMethod:%s\n", req.Host, req.URL.Path, req.Method)
	}
	res.Write([]byte(out))
}
func main() {
	fmt.Println("Запускаем сервер")
	http.HandleFunc(`/`, handler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Работа окончена")
}
