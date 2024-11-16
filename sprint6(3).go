package main

//TODO: Работа с ответами
//в стандартной библиотеке есть функции, которые принимают переменные типа "io.Writer"
//в качетсве параметра и записывают туда данные
//поэтому можно выводить тело ответа и другими способами

//func handle(w http.ResponseWriter(Это интерфейс, который используется для записи ответа, который мы отправим обратно клиенту), req *http.Request(объект запроса, который содержит всю информацию о запросе (например, URL, заголовки, данные запроса)){
//	io.WriteString(w,"Гофер.")                           //функция удобна для вывода строк
//	fmt.Fprint(w,"Еще один гофер!")                     //подобно "fmt.Printf()",записывает в "w" любые типы данных
//	fmt.Fprintf(w,"Где же третий? %s?","гофер")  //записывает данные с форматированием как "fmt.Printf"
//}

//TODO: Получение Get параметров

//полный URL запроса запроса хранится в поле "URL" структуры "http.Request",
//указатель на которую передается в обработчик вторым параметром
//т.к. Get параметры передаются в URL, то для их получения используется метод: "Query()url.Values"
//который возвращает значение типа "type Values map[string][]string"
//так можно перебрать все Get параметры:

//func handleQuery(res http.ResponseWriter,req *http.Request){
//	body:="GET параметры ===============\n"
//	for k,v:=range req.URL.Query(){  //это метод, который возвращает все GET-параметры запроса в виде объекта url.Values
//		fmt.Fprintf(res,"%s:%v\n",k,v)
//	}
//}

//если нужно получить значение конкретного параметра, стоит воспользоваться методом "(v Values) Get(key string)string
//данный метод возвращает первое значение указанного параметра.
//если запрашиваемого параметра не существует, то "Get()" вернет пустую строку

//name:req.URL.Query().Get("name")

//TODO: Получение HTTP-Заголовков
//HTTP-Заголовки можно получить из структуры "http.Request"
//заголовки запроса находятся в поле "Header" в виде мапы строковых слайсов "map[string][]string"
//как и Get параметры, заголовок может содержать несколько значений, поэтому и исплльзуется мапа слайсов,а не строк
//для получения заголовков существует два метода:
//(h Header)Values(key string)[]string - возвращает слайс значений указанного заголовка
//(h Header)Get(key string)string - возвращает первое значение заголовка

//func handleHeader(res http.ResponseWriter,req *http.Request){
//	agent:=req.Header.Get("User-Agent")
//	res.Write([]byte(agent))
//}

//func mainHandler(res http.ResponseWriter, req *http.Request) {
//	var answer string
//	name := req.URL.Query().Get("name")
//	if len(name) == 0 {
//		answer = "Укажите имя заголовка в параметре: http://localhost:8080/?name=User..."
//	} else if v := req.Header.Get(name); len(v) > 0 {
//		answer = fmt.Sprintf("%s:%s", name, v)
//	} else {
//		answer = fmt.Sprintf("Заголовок %s не определен", name)
//	}
//	io.WriteString(res, answer)
//}
//func main() {
//	http.HandleFunc(`/`, mainHandler)
//	err := http.ListenAndServe("localhost:8080", nil)
//	if err != nil {
//		panic(err)
//	}
//}
//если мы запустим этот сервер в браузере,то веб-страница будет содержать примерно такой текст: Accepr-Language: ru,en;q=0.9

//TODO: Получение POST-Параметров
//если вы заполняете в интернете какую-либо форму,данные этой формы отправляются на сервер методом "POST"
//получение таких параметров не сложнее получения параметров,указанных в URL

//метод "(r *Request)FormValue(key string)string" возвращает первое значение параметра с указанным именем
//при этом может быть возвращен как "POST", так и GET-параметр, но поиск начинается с POST параметров

//метод "(r *Request)PostFormValue(ket string)string" отличается от предыдущего метода тем,
//что ищет только среди параметров POST-запроса
//он также возвращает первое значение параметра с уазанным именем

//если параметр отсутствует, оба метода возвращают пустую строку

//сервер,который обрабатывает POST-запросы:

//func mainHandle(res http.ResponseWriter, req *http.Request) {
//	if req.Method == http.MethodPost {
//		fmt.Fprintf(res, "Email:%s\nName:%s", req.PostFormValue("email"), req.PostFormValue("name"))
//		return
//	}
//	io.WriteString(res, "Отправьте POST-запрос с параметрами email и name")
//}
//func main() {
//	http.HandleFunc(`/`, mainHandle)
//	err := http.ListenAndServe("localhost:8080", nil)
//	if err != nil {
//		panic(err)
//	}
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

//она вначале отправляет GET запрос,а затем POST и в результате должна вывести:
//Отправьте POST-запрос с параметрами email и name
//Email:my@my-best-site.ru
//Name:Василий

//TODO: Формирование ответа
//мы уже знакомы с интерфейсным типом "http.ResponseWriter", методы которого используются для форматировния и отправки ответа клиенту:

//type ResponseWriter interface{
//	Header() Header
//	Write([]byte)(int,error)
//	WriteHeader(statusCode int)
//}

//по умолчанию сервер добавляет в ответ некоторые HTTP-заголовки и код статуса,
//но можно возвращать дополнительные заголовки и указывать нужный код статуса

//метод "Header()" возвращает объект типа "http.Header"
//он содержит имена и значения заголовков.Для записи нужных заголовков можно использовать такие методы:
//"(h Header)Set(key string,value string)" - установить заголовок
//"(h Header)Add(key string,value string)" - добавить к существующему заголовку значение
//"(h Header)Del(key string)" - удалить заголовок

//метод "WriteHeader()" записывает в ответ сервера текущие установленные HTTP заголовки и код статуса.
//после вызова этой функции изменения заголовков не будут влиять на ответ сервера
//если вызов "WriteHeader()" отсутствует,то заголовки запишутся автоматически с кодом статуса 200

//в пакете "net/http" для кодов статуса определены соответствующие константы. Вот. некоторые из них:
//StatusOk = 200
//StatusBadRequest = 400
//StatusUnauthorized = 401
//StatusForbidden = 403
//StatusNotFound = 404
//StatusMethodNowAllowed = 405
//StatusInternalServerError = 500

//Вот пример обработчика, который возвращает данные в формате HTML, но принимает только GET-запросы:
//const pattern = `<!DOCTYPE html>
// <html lang="ru"><head>
// <meta charset="utf-8" />
// <title>Тестовый сервер</title>
// </head>
//<body>%s</body></html>`
//
//func mainHandle(w http.ResponseWriter, req *http.Request) {
//	var answer string
//
//	// устанавливаем заголовок Content-Type
//	w.Header().Set("Content-Type", "text/html; charset=utf-8")
//	if req.Method != http.MethodGet {
//		w.WriteHeader(http.StatusMethodNotAllowed)
//		fmt.Fprintf(w, pattern, "Сервер поддерживает только GET-запросы")
//		return
//	}
//	fmt.Fprintf(w, pattern, "Получен GET-запрос")
//}

//измененный хэндлер в предыдущей версии сервера
//если открыть http://localhost:8080/ в браузере, то отобразится "Получен GET-запрос"
//после этого запускаем такую версию клиента:

//func main() {
//	resp, err := http.PostForm("http://localhost:8080/", url.Values{})
//	if err != nil {
//		fmt.Println("Ошибка", err)
//		return
//	}
//	fmt.Println("Код статуса:", resp.StatusCode)
//	// читаем тело ответа
//	body, err := io.ReadAll(resp.Body)
//	resp.Body.Close()
//	if err != nil {
//		fmt.Println("Ошибка", err)
//		return
//	}
//	fmt.Println(string(body))
//}

//программа выведет:
//Код статуса: 405
//<!DOCTYPE html>
//<html lang="ru"><head>
//<meta charset="utf-8" />
//<title>Тестовый сервер</title>
//</head>
//<body>Сервер не поддерживает POST-запросы</body></html>

//Для возвращения ошибок можно использовать функцию http.Error(w ResponseWriter, error string, code int)
//Она возвращает ответ в виде текста с cообщением об ошибке и указанным кодом статуса
//При этом использовать метод Write() в этом случае не требуется.
//if req.Method != http.MethodGet {
//http.Error(w, fmt.Sprintf("Сервер не поддерживает %s запросы", req.Method),
//http.StatusMethodNotAllowed)
//return
//}

//TODO: По-новой

//TODO: получение GET параметров

//func handle(res http.ResponseWriter,req *http.Request){
//	body:="GET параметры ===============\n"
//	for k,v:=range req.URL.Query(){              //метод "URL.Query()" нужен для доступа к GET параметрам,он возвращает значение типа "Values(map[string][]string)"
//		fmt.Fprintf(res, "%s: %v\n",k,v)
//	}
//}

//func handle(res http.ResponseWriter,req *http.Request){
//	name:=req.URL.Query().Get("name")
//	res.Write([]byte(name))
//}

//TODO: получение HTTP-заголовков
//получить заголовок можно из параметра "http.Request"
//заголовки запроса находятся в поле "Header" в виде "map[string][]string"
//как и Get параметр заголовок может содержать несколько значений
//есть два способа получитб HTTP-заголовки:
//(h Header)Values(key string)[]string - возвращает слайс значений
//(h Header)Get(key string)string - возвращает первое значение заголовка
//
//func handlerr(res http.ResponseWriter,req *http.Request){
//	agent:=req.Header.Get("User-Agent")
//	res.Write([]byte(agent))
//}

//реализация:

//func handlerrr(res http.ResponseWriter, req *http.Request) {
//	var answer string
//	name := req.URL.Query().Get("name")
//	if len(name) == 0 {
//		answer = "Укажите имя заголовка в параметра: http://localhost:8080/"
//	} else if v := req.Header.Get(name); len(v) > 0 {
//		answer = fmt.Sprintf("%s: %s", name, v)
//	} else {
//		answer = fmt.Sprintf("Заголовок %s не определен", name)
//	}
//	io.WriteString(res, answer)
//}
//func main() {
//	fmt.Println("Запускаем сервер")
//	http.HandleFunc(`/`, handlerrr)
//	err := http.ListenAndServe(":8080", nil)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("Работа окончена")
//
//}

//TODO: получение POST-параметров

//метод "(r *Request)FormValue(key string)string" - возвращает первое значение параметра с указанным именем
//при этом может быть возвращен как POST, так и GET параметр,но поиск начинается с POST параметров

//метод "(r *Request)PostFormValue(key string)string"- ищет только среди параметров POST-запросоа,возвращает пераое значение параметра с указанным именем

//если параметр отсутствует,оба метода возвращают пустую строку

//пример такого сервера:

//func hhandler(res http.ResponseWriter, req *http.Request) {
//	if req.Method == http.MethodPost {
//		fmt.Fprintf(res, "Email: %s\nName: %s", req.PostFormValue("email"), req.PostFormValue("name"))
//		return
//	}
//	io.WriteString(res, "Отправьте POST запрос с параметрами email и name")
//}
//func main() {
//	fmt.Println("Запускаем сервер")
//	http.HandleFunc(`/`, hhandler)
//	err := http.ListenAndServe(":8080", nil)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("Работа окончена")
//}

//клиент для проверки сервера:

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
//	response, err := http.Get("http://localhost:8080/")
//	if err != nil {
//		fmt.Println("Ошибка", err)
//		return
//	}
//	printAnswer(response)
//	response, err = http.PostForm("http://localhost:8080/", url.Values{
//		"email": {"my@my-best-site.ru"},
//		"name":  {"Василий"},
//	})
//	if err != nil {
//		fmt.Println("Ошибка", err)
//		return
//	}
//	printAnswer(response)
//}

//TODO: формирование ответа

//методы для форматирования и отправки ответов клиенту:
//type ResponseWriter interface{
//	Header() Header
//	Write([]byte)(int error)
//	WriteHeader(statusCode int)
//}

//метод Header() содержит имена и значение заголовков и возвращает объект типа http.Header
//для записи нужных заголовков можно использовать соедующие методы:
//(h Header)Set(key,value string) - установить заголовок
//(h Header)Add(key, value string) - добавить значение к существующему заголовку
//(h Header)Del(key string) - удалить заголовок

//метод WriteHeader() записывает в ответ сервера текущие установленные HTTP-заголовки и код статуса
//после запуска данной функции изменения заголовков не будут влиять на ответ сервера
//если вызов WriteHeader() отсутствует,то заголовки запишутся автоматически с кодом статуса 200(OK)

//TODO: возвращение ошибок

//для возвращения ошибок можно использовать функцию http.Error(w ResponseWriter,error string,code int)
//данная функция возвращает ответ в виде текста с сообщением об ошибке c указанным кодом статуса
//в данном случае не потербуется вызывать метод Write([]byte)
//func handleer(res http.ResponseWriter,req *http.Request) {
//	if req.Method != http.MethodGet{
//		http.Error(res, fmt.Sprintf("Сервер не поддерживает %s запросы",req.Method),
//			http.StatusMethodNotAllowed)
//		return
//	}
//}

//TODO: маршрутизация

//маршрутизация - механизм, распределяющий запросы по нужным обработчикам
//запросы направляются к разным обработчикам в зависимости от указанного URL-пути
//в Go маршрутизатор по умолчанию - "http.DefaultServeMux"

//func handleTime(res http.ResponseWriter, req *http.Request) { //создали первый хендлер
//	a := time.Now().Format("02.01.2006 15:04:05")
//	res.Write([]byte(a))
//}
//func handleMain(res http.ResponseWriter, req *http.Request) { //создали второй хендлер
//	a := fmt.Sprintf("Host:%s\nPath:%s\nMethod:%s\n", req.Host, req.URL.Path, req.Method)
//	res.Write([]byte(a))
//}
//func main() {
//	http.HandleFunc(`/time`, handleTime) //обработали путь первого хендлера
//	http.HandleFunc(`/`, handleMain)     //обработали путь второго хендлера
//	err := http.ListenAndServe(":8080", nil)
//	if err != nil {
//		panic(err)
//	}
//}

//правила определения шаблонов
//http.HandleFunc можно вызывать в любом порядке
//но стоит учитывать,что

//если паттерн не заканчивается `/`, то требуется точное совпадение пути и шаблона

//если паттерн заканчивается `/`, то тогда хендлер будет обрабатывать все пути,начало которых совпадает с шаблоном(включая `/`)
//а так же путь без `/` (если для него нет обработчика)
//например, если указать шаблон "/time/", то обработчик сработает на "/time" (если нет обработчика для "/time"),"/time/","/time/get","/timer"
//не обработаются, так как нет совпадения с "/time/"

//так же учитываем, что если происходит регистрация разных хендлеров с одинаковым шаблоном
//то вызовется паника на этапе выполнения программы
//в шаблонах не получится использовать регулярные выражения и маски,в которых
//*-любая последовательность символов, а ?-один символ

//TODO: тип ServeMux

//"ServeMux" нужна для сложных проектов,дело в том,что если в коде встретится вызов "HandleFunc()" с каким-либо шаблоном и обработчиком, то это встроится в наш сервер
//тип "ServeMux" является собственным маршрутизатором
//создается функцией "NewServeMux()*ServeMux"
//после этого для переменной типа "*ServeMux" нужно вызвать методы "HandleFunc" с маршрутами и обработчиками
//правила остаются такими же,как и для маршрутизатора по умолчанию
//при вызове функции "ListenAndServe()" переменная-маршрутизатор передается во втором параметре

//func handlerTime(res http.ResponseWriter, req *http.Request) {
//	a := time.Now().Format("02.01.2006 15:04:05")
//	res.Write([]byte(a))
//}
//func handlerMain(res http.ResponseWriter, req *http.Request) {
//	a := fmt.Sprintf("Host:%s\nPath:%s\nMethod:%s\n", req.Host, req.URL.Path, req.Method)
//	res.Write([]byte(a))
//}
//func main() {
//	mux := http.NewServeMux()
//	mux.HandleFunc("/time/", handlerTime)
//	mux.HandleFunc("/", handlerMain)
//	err := http.ListenAndServe(":8080", mux)
//	if err != nil {
//		panic(err)
//	}
//}
