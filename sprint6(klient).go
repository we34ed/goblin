package main

import (
	"fmt"
	"net/url"
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

// // TODO: чтобы регулярно опрашивать сервер и получать текущее время,помещаем код с запросом в бесконечный цикл:
//
//	func main() {
//		for {
//			response, err := http.Get("http://localhost:8080/time")
//			if err != nil {
//				fmt.Println("Ошибка", err)
//				break
//			}
//			body, err := io.ReadAll(response.Body)
//			response.Body.Close()
//			if err != nil {
//				fmt.Println("Ошибка", err)
//				break
//			}
//			fmt.Println(string(body))
//			fmt.Println(response.Status)
//			time.Sleep(5 * time.Second)
//		}
//	} //чтобы завершить работу,нажимаем "control C"
//
//	func main() {
//		response, err := http.Get("http://localhost:8080/time")
//		if err != nil {
//			fmt.Println("Ошибка", err)
//			return
//		}
//		body, err := io.ReadAll(response.Body)
//		response.Body.Close()
//		if err != nil {
//			fmt.Println("Ошибка", err)
//		}
//		fmt.Println(string(body))
//		fmt.Println(response.Status)
//	}
//func main() {
//
//	response, err := http.Get("http://localhost:8080/time/") //функция возвращает указатель на структуру типа "http.Response"(ответ сервера)
//	if err != nil {
//		fmt.Println("Ошибка", err)
//		return
//	}
//	body, err := io.ReadAll(response.Body) //читаем данные благодаря "io.ReadAll"
//	response.Body.Close()                  //завершаем работы с данными, освобождая память
//	if err != nil {
//		fmt.Println("Ошибка", err)
//		return
//	}
//	fmt.Println(string(body))
//	fmt.Println(response.Status)
//}

//func printAnswer(res *http.Response) {
//	body, err := io.ReadAll(res.Body)
//	res.Body.Close()
//	if err != nil {
//		fmt.Println("Ошибка", err)
//		return
//	}
//	fmt.Println(string(body))
//}
//func main() {
//	resp, err := http.Get("http://localhost:8080/")
//	if err != nil {
//		fmt.Println("Ошибка", err)
//		return
//	}
//	printAnswer(resp)
//
//	resp, err = http.PostForm("http://localhost:8080/", url.Values{
//		"email": {"my@my-best-site.ru"},
//		"name":  {"Василий"},
//	})
//	if err != nil {
//		fmt.Println("Ошибка", err)
//		return
//	}
//	printAnswer(resp)
//}

// она вначале отправляет GET запрос,а затем POST и в результате должна вывести:
// Отправьте POST-запрос с параметрами email и name
// Email:my@my-best-site.ru
// Name:Василий
//
//	func main() {
//		response, err := http.Get("http://localhost:8080/time")
//		if err != nil {
//			fmt.Println("Ошибка при отправлении запроса", err)
//			return
//		}
//		defer response.Body.Close()
//		body, err := io.ReadAll(response.Body)
//		if err != nil {
//			fmt.Println("Ошибка при получении ответа", err)
//			return
//		}
//		fmt.Println(string(body))
//		fmt.Println(response.Status)
//	}
//
//	func main() {
//		response, err := http.Post("http://localhost:8080", "text/plain", nil)
//		if err != nil {
//			fmt.Println("Ошибка при отправке запроса", err)
//			return
//		}
//		defer response.Body.Close()
//		body, err := io.ReadAll(response.Body)
//		if err != nil {
//			fmt.Println("Ошибка при получении ответа", err)
//			return
//		}
//		fmt.Println(string(body))
//		fmt.Println(response.Status)
//	}
//func main() {
//	client := &http.Client{
//		Timeout: 1 * time.Second,
//	}
//	request, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
//	if err != nil {
//		fmt.Println("Ошибка при форматировании запроса", err)
//		return
//	}
//	request.Header.Set("Custom-Header", "StasYao")
//	request.Header.Add("Custom-Header", "1234")
//
//	response, err := client.Do(request)
//	if err != nil {
//		fmt.Println("Ошибка при отправке запроса", err)
//		return
//	}
//	defer response.Body.Close()
//	body, err := io.ReadAll(response.Body)
//	if err != nil {
//		fmt.Println("Ошибка чтения ответа", err)
//		return
//	}
//	fmt.Println(string(body))
//	fmt.Println(response.Status)
//}

//TODO: Отправка формы

//func main(){
//	values:=url.Values{}
//	values.Set("nickname","Student")
//	values.Set("feedbacl","Все отлично!")
//	response,err:=http.PostForm("http://localhost:8080/",values)
//}

//квиз

//func formValues(achievement string) url.Values {
//	values := url.Values{}
//	values.Set("name", "Stasya")
//	values.Set("nickname", "Legend")
//	values.Set("achieves", "Master")
//	values.Add("achieves", "Rich")
//	values.Add("achieves", achievement)
//	return values
//}
//func main() {
//	v := formValues("HoodRich")
//	fmt.Println(v)
//	fmt.Println(v["name"])
//	fmt.Println(v["nickname"])
//	fmt.Println(v["achieves"])
//}

func main() {
	values := url.Values{}
	values.Set("name", "Stasya")
	values.Set("nickname", "HoodRich")
	values.Set("level", "Master")
	values.Add("level", "Boss")
	values.Add("level", "ProntoDeGusta")
	fmt.Println(values["name"])
	fmt.Println(values["nickname"])
	fmt.Println(values["level"])
}
