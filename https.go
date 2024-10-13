package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//	//http-ответ записывается в переменную responce
	//	//метод http.Get нужен для открытия указанной страницы(отправляет запрос к веб-странице)
	//	responce, err := http.Get("http://info.cern.ch")
	//	if err != nil { //переменная err-отображает возможную ошибку
	//		fmt.Println(err)
	//		return
	//	}
	//	//тело http-ответа будет записано в переменную body,переменная err-отображает возможную ошибку
	//	//метод io.ReadAll нужен для чтения http-ответа
	//	body, err := io.ReadAll(responce.Body) //в свойстве .Body этого объекта(responce) хранится хранится HTML-код
	//	if err != nil {                        //переменная err-отображает возможную ошибку
	//		fmt.Println(err)
	//		return
	//	}
	//	result := string(body) //преобразовываем тело http- ответа из байтов в строку
	//	fmt.Println(result)
	////}
	//	responce, err := http.Get("https://yandex.ru/search/?text=что+такое+backend&lr=213")
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	fmt.Println(responce.Status) //метод .Status позволяет увидеть работает ли страница
	//}

	//baseURL := "https://yandex.ru/search/" //определили url,к которому будет обращение
	//params := url.Values{}                 //создали объект,в котором хранятся параметры
	//
	////заполняем параметры для url как пару из ключ-значение
	//params.Add("text", "что такое backend")
	//params.Add("lr", "213")

	//преобразовываем URL-строку baseUrl в URL-объект
	//urlInstance, err := url.Parse(baseURL)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	////задаем параметры поиска
	//urlInstance.RawQuery = params.Encode() //метод .RawQuery нужен для параметров поиска
	//
	////преобразовываем полученный URL с параметрами в строку
	//formattedUrl := urlInstance.String()
	//fmt.Println(formattedUrl)
	//
	////теперь выполняем запрос к url с заполненными параметрами и получаем ответ
	//responce, err := http.Get(formattedUrl)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(responce.Status)
	//}

	//	bazaURL := "http://wttr.in/" //определили url для обращения
	//	parametrs := url.Values{}    // определили параметры url
	//	parametrs.Add("u", "")       //для параметров задали только имена(ключи)
	//	parametrs.Add("T", "")
	//
	//	preobr, err := url.Parse(bazaURL) //преобразовали url-строку в url-объект, заполнили параметры
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	preobr.RawQuery = parametrs.Encode()
	//	responce, err := http.Get(preobr.String()) //выполняем запрос и обрабатываем ответ
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	fmt.Println(responce.Status)
	//	body, err := io.ReadAll(responce.Body)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	result := string(body)
	//	fmt.Println(result)
	//}
	//	responce, err := http.Get("https://ya.ru/white")
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	fmt.Println("Status code:", responce.StatusCode)                  //получили статус кода и записали его в свой код
	//	fmt.Println("Content type:", responce.Header.Get("Content-Type")) //прочитали заголовок и вывели на печать
	//}

	request, err := http.NewRequest(
		"GET", "https://habr.com/", nil) //строкой передаем URL платформы habr
	if err != nil {
		fmt.Println(err)
		return
	}
	request.Header.Add("Accept-language", "ru") //запрашиваем материал на русском

	client := &http.Client{} //создаем материал на русском

	responce, err := client.Do(request) //выполняем ранее созданый запрос,с указанным адресом и методом
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(responce.Status)
	body, err := io.ReadAll(responce.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	result := string(body)
	fmt.Println(result)
}

//todo:знать!
