package main

//TODO:Указатели

//TODO: '&'-ссылается на адресс переменной
//TODO: '*'-ссылается на значение переменной

//todo:указатель "&" -указывает на адрес переменной
//todo:указатель - переменная в которой хранится адрес объекта
//func main() {
//	var i int64
//	var j int64
//	var n int32
//	b := true
//	pi := 3.1415
//	fmt.Println(i, &j, &n, &b, &pi)
//}
//чтобы получить указатели нужно сoхранить адреса

//TODO: Способы объявления указателя в коде
//todo:1) сохранение в переменной

//func main() {
//	i := 10
//	result := &i
//	fmt.Println(result)
//}

//todo: 2)объявить указатель через "var" и присвоить нужный адрес
//todo: нужно указать('*'),на какой тип будет ссылаться указатель(разыменовыватель указателя)
//
//func main() {
//	var a *int
//	var b int
//
//	a = &b
//	fmt.Println(&a)
//}

// todo: 3)создать указатель через встроенную фуннкцию "new(тип)"
// todo: указатели могут ссылаться на переменные любых типов
//
//	func main() {
//		a := new(map[string]int)
//		b := new([]int)
//		c := new(string)
//		d := new(MyType(struct))
//	}

// todo:нулевое значение(nil)
// нулевое значение любого указателя - nil
//
//	func main() {
//		var ptr *int
//		if ptr != nil {
//	    //указатель ptr не пустой,с ним можно работать)
//		}
//	}
//todo:любому указателю можно присвоить его нулевое значение(nil)

// TODO:Указатели на структуры
//type Hero struct {
//	name   string
//	health int
//	age    int
//}
//
//func (h *Hero) damage(dmg int) int {
//	h.health -= dmg
//	return h.health
//}
//func (h *Hero) ChangeHealth(hp int) {
//	if h.health < 0 {
//		h.health -= hp
//	} else {
//		h.health += hp
//	}
//}
//
//func main() {
//	hero := Hero{
//		"Stas",
//		100,
//		20,
//	}
//	hero.ChangeHealth(10)
//	fmt.Println(hero)
//}

// TODO: Разыменование указателей
//func main() {
//	a := 123
//	b := &a
//	*b = 234
//	fmt.Println(*b)
//}

//type Queue struct {
//	first *QueueItem
//	last  *QueueItem
//}
//type QueueItem struct {
//	value string
//	next  *QueueItem
//}
//
//func (q *Queue) Push(value string) {
//	item := &QueueItem{value: value}
//	if q.first == nil {
//		q.first = item
//		q.last = item
//		return
//	}
//	q.last.next = item
//	q.last = item
//}
//
//
//type Character struct {
//	Level  int
//	Health int
//}
//
//func main() {
//	person := Character{
//		Level:  80,
//		Health: 100,
//	}
//	pointer := &person
//	fmt.Println(person.Level, person.Health)
//	*pointer = Character{100, 150}
//	fmt.Println(person.Level, person.Health)
//}

// todo:квиз
//func main() {
//	numbers := [5]int{1, 3, 8, 19, 42}
//	pointers := make([]*int, len(numbers))
//	for i := range numbers {
//		pointers[i] = &numbers[i]
//	}
//	for i := range pointers {
//		*pointers[i] *= 3
//	}
//	fmt.Println(numbers)
//}

// TODO:Передача параметров по ссылке
//
//	func myfunc(s string, i int) {
//		s += "*"
//		i = i * 3 / 2
//		fmt.Println("Функция myfunc:", s, i)
//	}
//
//	func main() {
//		ss := "Тест"
//		ii := 20
//		myfunc(ss, ii)
//		fmt.Println("Функция main:", ss, ii)
//	}
//
// todo:передача значений работает для всех типов,кроме мапов и слайсов,в них данные передаются по ссылке
// todo:их можно изменять внутри функции
//func myfunc(s []int, m map[string]int) {
//	s[0] = 101
//	m["one"] = 1
//	delete(m, "zero")
//}
//func main() {
//	slice := []int{1, 2, 3, 4}
//	numbers := map[string]int{
//		"zero": 0,
//	}
//	myfunc(slice, numbers)      //передаем слайс и мапу в параметре
//	fmt.Println(slice, numbers) //т.к. они передаются по ссылке, функция изменила их элементы
//}

// todo:как изменить и не копировать(использовать указатели)

//todo: *x,*y=*y,*x(swap)-одним присваиванием изменили несколько переменных

//func swap(x, y *int) {
//a := *x
//*x = *y
//*y = a
//*x, *y = *y, *x
//}
//func main() {
//	i := 25
//	j := 36
//	swap(&i, &j)
//	fmt.Println(i, j)
//}
//
//func CheckYourself(ptr *int,x,val int)int{
//	ret:=ptr
//	*ptr+=val
//	x+=val
//	return 2* *ret
//}
//func main(){
//	i:=6
//	j:=11
//	fmt.Println(CheckYourself(&i,j,4),i,j)
//}
//TODO: передача параметров по ссылке/по значению

//todo: по ссылке:в стек попадают указатели на переменные
//1)функция должна менять значения параметра
//2)экономия места на копировании данных

// todo: по значению:в стек попадают копии значений переменных
//func div(x, y int) (int, int) {
//	return x / 2, y / 2
//
//}
//func main() {
//	x := 10
//	y := 20
//	fmt.Println(div(x, y))
//}

//todo:квиз

//	func array(arr *[5][5]int) {
//		for i, v := range arr {
//			for idx, e := range v {
//				if e == 1 {
//					(*arr)[i][idx] = 0
//				} else {
//					(*arr)[i][idx] = 1
//				}
//			}
//		}
//	}
//
//	func main() {
//		arrray := [5][5]int{
//			{0, 1, 1, 0, 1},
//			{1, 0, 1, 1, 0},
//			{0, 1, 0, 0, 1},
//			{1, 0, 1, 1, 0},
//			{0, 1, 0, 0, 1},
//		}
//		array(&arrray)
//		fmt.Println(arrray)
//	}
//func array(arr *[2][2]int) {
//	for i, v := range arr {
//		for ind, e := range v {
//			if e == 1 {
//				(*arr)[i][ind] = 0
//			} else {
//				(*arr)[i][ind] = 1
//			}
//		}
//	}
//}
//func main() {
//	arrayy := [2][2]int{
//		{1, 0},
//		{0, 1},
//	}
//	array(&arrayy)
//	fmt.Println(arrayy)
//}

//TODO:Использование указателей со структурами

//type Hero struct {
//	name string
//}
//
//func (h Hero) Name() string {
//	return h.name
//}
//func (h *Hero) SetName(name string) {
//	h.name = name
//}
//func main() {
//	hero := Hero{
//		name: "Stas",
//	}
//	hero.SetName("Stanislav")
//	fmt.Println(hero.name)
//
//}

//	type Thing struct {
//		name string
//	}
//
//	type Hero struct {
//		name  string
//		thing []*Thing
//	}
//
//	func (h Hero) Doing() {
//		for i, v := range h.thing {
//			h.thing[i].name = strings.ToUpper(v.name)
//		}
//	}
//
//	func main() {
//		hero := Hero{
//			name:  "Тор",
//			thing: []*Thing{{name: "Молот"}, {name: "Доспехи"}},
//		}
//		hero.Doing()
//		for _, v := range hero.thing {
//			fmt.Println(v.name)
//		}
//	}
//type Hero struct {
//	name   string
//	level  int
//	health int
//}
//
//func (h *Hero) LevelUp() {
//	h.level += 1
//}
//func (h *Hero) HealthUp() {
//	h.health += 1
//}
//func (h *Hero) Lvl(pointer int) {
//	h.level += pointer
//}
//func (h *Hero) Hp(pointer int) {
//	h.health += pointer
//}
//func main() {
//	hero := Hero{
//		name:   "Stas",
//		level:  10,
//		health: 100,
//	}
//	hero.LevelUp()
//	fmt.Println(hero.level)
//	hero.HealthUp()
//	fmt.Println(hero.health)
//	hero.Lvl(10)
//	fmt.Println(hero.level)
//	hero.Hp(100)
//	fmt.Println(hero.health)
//	fmt.Println(hero)
//}

//TODO:Указатели на структуры можно передавать и в параметре функций
//
//type Config struct{
//	host string
//	port int
//}
//func SetConfig(cfg *Config)*Config{
//	cfg.host="LocalHost"
//	cfg.port=123
//	return cfg
//}

// TODO:Указатели при эмбеддинге
//type Person struct {
//	Name   string
//	Health int
//}
//type Warrior struct {
//	Person
//	Attack int
//}
//
//func (p *Person) SetName(name string) {
//	p.Name = "Stas"
//}
//func main() {
//	warrior := Warrior{Person:&Person,Attack:100}
//	warrior.SetName("Конан")
//	fmt.Println(warrior)
//}

// TODO:Квиз
//
//	type Queue struct{
//		first *QueueItem
//		last *QueueItem
//	}
//
//	type QueueItem struct{
//		person *Character
//		next *QueueItem
//	}
//
//	type Character struct{
//		name string
//		level int
//	}
//
//	func (q *Queue)Pop()(*Character,bool){
//		if q.first==nil{
//			return nil,false
//		}
//		item:=q.first
//		q.first=item.next
//
//		if q.first==nil{
//			q.last=nil
//		}
//		return item.person,true
//	}
//
// todo: Разыменование указателей при обращении к полям в методе происходит автоматически
//
// todo: Если указатель на структуру равен nil,то обращение по полям вызовет панику
//
// TODO: Интерфейсы
// todo:интерфейс-тип данных,описывающий методы,которые реализуются в структурном типе
// todo:так описывается интерфейс:
//
//	type intefaceName interface{
//		method1()
//		method2()
//	}
//
// todo:это тоже самое что и структуры,только вместо "struct" пишем "interface", а вместо полей пише методы
//
//	type Loot interface{
//		apply() //применить
//		magic()int //получить
//		sell(cost int) //продать
//	}

//TODO:интерфейс String из библиотеки Go из пакета fmt
//todo:при использовании методов с этим типом функции из пакета fmt будут выводить данные в соответствии с методом String()
//type StarStr string
//
//func (str StarStr) String() string {
//	return `*** ` + string(str) + ` ***`
//}
//func main() {
//	var stars StarStr = `Интерфейсы - это просто!`
//	fmt.Println(stars)
//}
//todo:квиз
//
//type Character struct {
//	name   string
//	health int
//}
//type Warrior struct {
//	Character
//	power int
//}
//type Mage struct {
//	Character
//	magic int
//}
//
//func (w Warrior) String() string {
//	return `* ` + string(w.name)
//}
//func (m Mage) String() string {
//	return `* ` + string(m.name)
//}
//
//func main() {
//	fmt.Println(Warrior{Character: Character{name: "Крестоносец"}})
//	fmt.Println(Warrior{Character: Character{name: "Коммандос"}})
//	fmt.Println(Mage{Character: Character{name: "Шаман"}})
//	fmt.Println(Mage{Character: Character{name: "Друид"}})
//}
//TODO:Реализация интерфейса
//type AttackDefence interface{
//	attack() int
//	defence() int
//}
//type Character struct {
//	name   string
//	health int
//}
//type Warrior struct {
//	Character
//	power int
//}
//type Mage struct {
//	Character
//	magic int
//}
//func(w Warrior)attack()int{
//	return 10
//}
//func (w Warrior)defence()int{
//	return 6
//}
//func (m Mage)attack()int{
//	return 7
//}
//func (m Mage)defence()int{
//	return 9
//}
//func Fight(attacker AttackDefence, defender AttackDefence){
//	power:=attacker.attack()
//	protection:=defender.defence()
//}
//func Engine(){
//	var warrior Warrior
//	var mage Mage
//	Fight(warrior,mage)
//	Fight(mage,warrior)
//}

// TODO:Квиз
//type Sword struct {
//	power int
//}
//type Scroll struct {
//	magic int
//}
//type Loot interface {
//	apply()
//}
//
//func (sw Sword) apply() {
//	fmt.Println("Меч", sw.power)
//}
//func (sc Scroll) apply() {
//	fmt.Println("Свиток", sc.magic)
//}
//func main() {
//	loot := []Loot{
//		Sword{50},
//		Scroll{20},
//		Scroll{70},
//	}
//	for _, element := range loot {
//		element.apply()
//	}
//}

// TODO:Реализация интерфейсов
// todo:проверка ошибки на стадии компиляции (var _ Loot = new(Sword) или Loot = (*Sword)(nil)
//type Loot interface {
//	apply()
//}
//type Sword struct {
//	power int
//}
//
//func (s Sword) apply() int {
//	return 0
//}
//func main() {
//	var _ Loot = new(Sword)
//}
//
//todo:проверка ошибки при выполнении loot:=variable.(Loot) или loot,ok:=variable.(loot),где ок-"false"

//type Loot interface {
//	apply()
//}
//
//type Sword struct {
//	power int
//}
//
//func (s Sword) apply() {
//	fmt.Println("меч", s.power)
//}
//
//type Shoes struct {
//	model string
//}
//
//func (s Shoes) wear() {
//	fmt.Println("обувь", s.model)
//}
//
//// todo:"any"-алиас от interface{}
//func apply(object []any) {
//	for _, variable := range object {
//		loot, ok := variable.(Loot)
//		if ok {
//			loot.apply()
//		}
//	}
//}
//func main() {
//	apply([]any{
//		Sword{35},
//		Shoes{"Скороходы"},
//	})
//}

//TODO:План реализации интерфейса

//	type Property struct {
//		name  string
//		price int
//	}
//
//	type Loot interface {
//		apply()
//		properties() Property
//		sell() int
//	}
//
//	type Scroll struct {
//		Property
//	}
//
//	type Sword struct {
//		Property
//	}
//
// func (s Scroll) apply() {
// }
//
//	func (s Scroll) properties() Property {
//		return s.Property
//	}
//
//	func (s Scroll) sell() int {
//		return s.price
//	}
//
// func (s Sword) apply() {
// }
//
//	func (s Sword) properties() Property {
//		return s.Property
//	}
//
//	func (s Sword) sell() int {
//		return s.price
//	}
//
//	func main() {
//		loot := []Loot{
//			Scroll{Property{"Свиток огня", 1000}},
//			Sword{Property{"Меч короля", 2000}},
//		}
//		for _, item := range loot {
//			item.apply()
//			fmt.Println(item.properties())
//			fmt.Println(item.sell())
//			fmt.Println("Успех!")
//		}
//	}
//
//	type Shape interface {
//		Area() float64
//	}
//
//	type Circle struct {
//		Radius float64
//	}
//
//	func (c Circle) Area() float64 {
//		return math.Pi * c.Radius * c.Radius
//	}
//
//	type Rectangle struct {
//		Width, Height float64
//	}
//
//	func (r Rectangle) Area() float64 {
//		return r.Width * r.Height
//	}
//
//	func area(s Shape) {
//		fmt.Printf("Площадь равна:%.2f\n", s.Area())
//	}
//
//	func main() {
//		a := Circle{
//			5,
//		}
//		b := Rectangle{
//			10,
//			5,
//		}
//		area(a)
//		area(b)
//	}
//
//	type Vehicle interface {
//		Drive() string
//		FuelType() string
//		MaxSpeed() int
//	}
//
//	type car struct {
//		brand string
//		fuel  string
//		speed int
//	}
//
//	type moto struct {
//		car
//	}
//
//	func (c car) Drive() string {
//		return fmt.Sprintf("фирма: %s", c.brand)
//	}
//
//	func (c car) FuelType() string {
//		return fmt.Sprintf("жрет %s бензин", c.fuel)
//	}
//
//	func (c car) MaxSpeed() int {
//		return c.speed
//	}
//
//	func (m moto) Drive() string {
//		return fmt.Sprintf("фирма: %s", m.brand)
//	}
//
//	func (m moto) FuelType() string {
//		return fmt.Sprintf("жрет %s бензин", m.fuel)
//	}
//
//	func (m moto) MaxSpeed() int {
//		return m.speed
//	}
//
//	func vehicleInfo(v Vehicle) {
//		fmt.Println("Транспорт:", v.Drive())
//		fmt.Println("Расход:", v.FuelType())
//		fmt.Println("Максимальная скорость:", v.MaxSpeed(), "км.ч\n")
//	}
//
//	func main() {
//		a := car{
//			"Porsche",
//			"сотый",
//			300,
//		}
//		vehicleInfo(a)
//		b := moto{
//			car{
//				"Yamaha",
//				"сотый",
//				325,
//			},
//		}
//		vehicleInfo(b)
//	}
//
//	type Item interface {
//		String() string
//	}
//
//	type Sword struct {
//		name string
//	}
//
//	type Scroll struct {
//		name string
//	}
//
//	func (s Sword) String() string {
//		return `*** ` + string(s.name) + ` ***`
//	}
//
//	func (s Scroll) String() string {
//		return `*** ` + string(s.name) + ` ***`
//	}
//
//	func main() {
//		a := Sword{
//			"Меч",
//		}
//		b := Scroll{
//			"Свиток",
//		}
//		fmt.Println(a.String())
//		fmt.Println(b.String())
//	}
