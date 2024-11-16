package main

//TODO: Отправка запросов
//метод "Post" отправляет на сервер дополнительные данные
//для этого можно использовать функцию "http.Post(url, contentType string, body io.Reader)(resp *Response,err error)"
//функция для метода "Post" имеет такие же возвращаемые значения,что и "Get",однако имеет два дополнительных параметра:

//1)contentType-содержит тип передаваемых данных,он будет записан в заголовок Content-Type
//2)body-переменная интерфейсного типа "io.Reader",из которой будут прочитаны передаваемые данные

//todo: передаем серверу запрос с данными в формате JSON(текстовый формат для описания структурируемых данных):

//func main(){
//	data:=`{"cmd": "gettime", "town": "Санкт-Петербург"}`
//	resp,err:=http.Post("http://localhost:8080/","application/json; charset=utf-8",strings.NewReader(data))
//}
//серверу придет POST-запрос с "Content-Type: application/json; charset=utf-8" и телом запроса в виде JSON-строки
//функция "strings.NewReader()" преобразует строку в переменную типа io.Reader

//TODO: отправляем запрос нашему серверу методом "Post"
//func main(){
//	response,err:=http.Post("http://localhost:8080","text/plain",nil)
//}
//при таком запросе сервер возратит:
//Host: localhost:8080
//Path: /
//Method: POST

//TODO:тип Client и функция NewRequest
//тип Client и его методы отвечают за отправку запросов в пакете "net/http"
//функции "http.Get" и "http.Host"-обертки над вызовом этих методов у глобальной переменной "http.DefaultClient"
//можно определить свою переменную "*http.Client" и использовать аналогичные методы "Get()" и "Post()"
//func main(){
//	client:=&http.Client{}
//	response,err:=client.Get("http://localhost:8080/")
//}

//клиент позволяет настроить несколько параметров
//в поле "Timeout""time.Duration" можно указать максимальную задержку ответа,после которой клиент отменяет запрос
//func main(){
//	client:=&http.Client{
//		Timeout: 1*time.Second,
//	}
//}

// если методы "client.Get()" и "client.Post()" не удовлетворяют всем потребностям,то можно создать запрос "*http.Request"
// и отправить его методом "(c *Client)Do(req *Request)(*Response,error)"
// для создания переменной "*http.Request" с запросом, следует использовать функцию
// "NewRequest(method,url string, body io.Reader)(*Request, error)",где:
// "method"-имя метода запроса
// "url"-URL адрес,по которому будет отправлен запрос
// "body"-переменная интерфейсного типа "io.Reader", которая должна содержать тело запроса,если тело запроса не нужно отправлять,пишем "nil"

//func main() {
//	client := &http.Client{
//		Timeout: 1 * time.Second, //данное поле служит для установки времени максимального ожидания ответа сервера
//	}
//	request, err := http.NewRequest(http.MethodGet, "http://localhost:8080/time", nil)
//	if err != nil {
//		fmt.Println("Ошибка форматирования запроса:", err)
//		return
//	}
//	response, err := client.Do(request) //этим методом можно отправлять запросы любого типа
//	if err != nil {
//		fmt.Println("Ошибка отправки запроса:", err)
//		return
//	}
//	defer response.Body.Close() //вызов после "defer" произойдет при выходе из функции,
//	//"Close()" с "defer" можно указать здесь, так как "Respose.Body.Close()" нужно вызывать в любом случае
//	body, err := io.ReadAll(response.Body)
//	if err != nil {
//		fmt.Println("Ошикба чтения ответа:", err)
//		return
//	}
//	fmt.Println(string(body))
//}

//вместо строк "GET" и "POST" можно использовать соответствующие константы пакета "http.MethodGet" и "http.MethodPost"

//func main() {
//	client := &http.Client{
//		Timeout: 1 * time.Second,
//	}
//	request, err := http.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
//	//используем функцию,если методы client.gET и client.Post не удовлетворяют всем потребностям
//	if err != nil {
//		fmt.Println("Ошибка форматирования запроса", err)
//		return
//	}
//	response, err := client.Do(request)
//	if err != nil {
//		fmt.Println("Ошибка отправки запроса", err)
//		return
//	}
//	defer response.Body.Close()
//	body, err := io.ReadAll(response.Body)
//	if err != nil {
//		fmt.Println("Ошибка чтения ответа", err)
//		return
//	}
//	fmt.Println(string(body))
//}

//func main() {
//	client := &http.Client{
//		Timeout: 1 * time.Second,
//	}
//	request, err := http.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
//	if err != nil {
//		fmt.Println("Ошибка при форматированиие запроса", err)
//		return
//	}
//	response, err := client.Do(request)
//	if err != nil {
//		fmt.Println("Ошибка при отправке запроса", err)
//		return
//	}
//	defer response.Body.Close()
//	body, err := io.ReadAll(response.Body)
//	if err != nil {
//		fmt.Println("Ошибка при получении ответа", err)
//		return
//	}
//	fmt.Println(string(body))
//	fmt.Println(response.Status)
//}

//TODO: Добавление HTTP-Заголовков
//если отправлять запросы с использованием "http.NewRequest", можно дополнительно установить HTTP-Заголовки
//структура "http.Request" имеет поле "Header" типа "http.Header"
//для добавления заголовка существует два метода:
//(h Header)Set(key,value string) - устанавливает значение заголовка
//(h Header)Add(key,value string) - добавляет значение заголовка к уже сузествующим,заголовок может иметь неск.значений

//func main() {
//	client := &http.Client{
//		Timeout: 1 * time.Second,
//	}
//	request, err := http.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
//	if err != nil {
//		fmt.Println("Ошибка при форматировании запроса", err)
//		return
//	}
//	request.Header.Set("Custom-Header", "John Doe") // установили значение заголовка
//	request.Header.Add("Custom-Header", "1234")     // добавили значение заголовка к существующему
//
//	response, err := client.Do(request)
//	if err != nil {
//		fmt.Println("Ошибка при отправке запроса", err)
//		return
//	}
//	defer response.Body.Close()
//	body, err := io.ReadAll(response.Body)
//	if err != nil {
//		fmt.Println("Ошибка при получении ответа", err)
//		return
//	}
//	fmt.Println(string(body))
//}

//TODO:Отправка формы
//если пользователь заполняет и отправляет с сайта какую-либо форму,то по умолчанию данные отправляются на сервер методом "POST"
//чтобы эмулировать отправку формы со стороны клиента, используется функция "http.Post()" или "client.Post()"
//при этом необходимо указать корректный "Content-Type" и перекодировать отправляемые данные

//TODO: путь короче:
//специально для отправки формы существует метод "(c *Client)PostForm(url string, data url.Values)(resp *Response,err error)
//т.к. тип "url.Values" - это map[string][]string",то можно инициализировать все параметры прямо при отправке формы:
//func main() {
//	response, err := http.PostForm("http://localhost:8080/",
//		url.Values{
//			"nickname": {"student"},
//			"feedback": {"Все отлично!"}})
//}

//в качестве альтернативы можно определять значения "url.Values" с помощью методов:
//(v Values)Add(key,Value string) - добавить значение, если значение не существует,то оно будет создано
//(v Values)Set(key,value string) - установить значение

//func main(){
//	values:=url.Values{}
//	values.Set("Name","Stas")
//	values.Add("Username","Weedle")
//	request,err:=http.PostForm("http://localhost:8080/",values)
//}

//func main() {
//	client := &http.Client{
//		Timeout: 1 * time.Second,
//	}
//	//values := url.Values{}
//	//values.Set("Name", "Stas")
//	//values.Add("UserName", "Weedle")
//	//request, err := http.PostForm("http://localhost:8080/", values)
//

// квиз
//func FormValues(achivement string) url.Values {
//	values := url.Values{}
//	values.Set("name", "Vasya")
//	values.Set("nickname", "superstar")
//	values.Set("achieves", "cool")
//	values.Add("achieves", "best")
//	values.Add("achieves", achivement)
//	return values
//}
//func main() {
//	v := FormValues("80 Level")
//	fmt.Println(v["name"], v["nickname"], v["achieves"])
//}

//	func main(){
//		client:=&http.Client{
//			Timeout:1*time.Second,
//		}
//		request,err:=http.NewRequest(http.MethodGet,"http://localhost:8080/",nil)
//		if err!=nil{
//			fmt.Println("Ошибка при форматировании",err)
//			return
//		}
//		request.Header.Set("Custom-Header","StasYao")
//		request.Header.Add("Custom-Header","1234")
//		response,err:=client.Do(request)
//		if err!=nil{
//			fmt.Println("Ошибка при отправлении запроса",err)
//			return
//		}
//		defer response.Body.Close()
//		body,err:=io.ReadAll(response.Body)
//		if err !=nil{
//			fmt.Println("Ошибка при получении ответа",err)
//			return
//		}
//		fmt.Println(string(body))
//		fmt.Println(response.Status)
//	}
//func main() {
//	client := &http.Client{
//		Timeout: 1 * time.Second,
//	}
//	request, err := http.NewRequest(http.MethodGet, "http://localhost:8080/time", nil)
//	if err != nil {
//		fmt.Println("Ошибкв при форматировании запроса", err)
//		return
//	}
//	request.Header.Set("Custom-Header", "Stasyao")
//	request.Header.Add("Custom-Header", "1234")
//	response, err := client.Do(request)
//	if err != nil {
//		fmt.Println("Ошибка при отправке запроса", err)
//		return
//	}
//	defer response.Body.Close()
//	body, err := io.ReadAll(response.Body)
//	if err != nil {
//		fmt.Println("Ошибка при получении ответа", err)
//		return
//	}
//	fmt.Println(string(body))
//	fmt.Println(response.Status)
//}
