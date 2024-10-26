package main

// import "fmt"
//
// //TODO:Структуры
//
// type People struct { //объявление структуры
//
//		name     string
//		username string
//		age      int
//		height   float64
//		man      bool
//	}
//
// //func main() {
// //	use := People{             //вызов структуры
// //		name:     "Stas",
// //		username: "we34ed",
// //		age:      21,
// //		height:   1.75,
// //		man:      false,
// //	}
// //	use.age += 1               //использование определенного фрагмента структуры через "."
// //	use.height += 0.2
// //	use.name += " Kuchma"
// //	use.man = true
// //	use.username += " boss"
// //	fmt.Println(use)
// //	fmt.Println(use.age)
// //	fmt.Println(use.name)
// //}
//
// //TODO:Методы
//
// //	func (p People) String() string {
// //		structure := fmt.Sprintf("имя:%s\nник:%s\nвозраст:%d\nрост:%.2f\n", people.name, people.username, people.age, people.height)
// //		return structure
// //	}
// //
// //	func main() {
// //		use := People{
// //			name:     "Stas",
// //			username: "weedle",
// //			age:      21,
// //			height:   1.75,
// //		}
// //		fmt.Println(use)
// //	}
// //
// //	func (p *People) ages() {
// //		p.age = p.age + 1
// //	}
// //
// //	func main() {
// //		use := People{
// //			name:     "Stas",
// //			username: "we34ed",
// //			age:      21,
// //			height:   1.75,
// //			man:      false,
// //		}
// //		use.ages()
// //		fmt.Println(use)
// //
// // }
// //
// //	func (s People) talks(slova string) {
// //		fmt.Println(s.name, slova)
// //	}
// //
// //	func main() {
// //		use := People{
// //			name:     "Stas",
// //			username: "we34ed",
// //			age:      21,
// //			height:   1.75,
// //			man:      false,
// //		}
// //		use.talks("Hello nigas")
// //	}
//
//	func (p People) Say() {
//		fmt.Println("Меня зовут:", p.name)
//	}
//
//	func main() {
//		use := People{
//			name:     "Stas",
//			username: "we34ed",
//			age:      21,
//			height:   1.75,
//			man:      false,
//		}
//		use.Say()
//	}
//
//	type person struct {
//		name    string
//		surname string
//	}
//
//	type work struct {
//		person
//		work string
//	}
//
//	func (w work) say() {
//		fmt.Println("my name is", w.person, "im workin at", w.work)
//	}
//
// func main() {
//
//		persona := person{
//			name:    "Stas",
//			surname: "Kuchma",
//		}
//		wwork := work{
//			person: person{name: "Slavik", surname: "Rudan"},
//			work:   "dota2",
//		}
//		fmt.Println("my name is", persona.name, "my surname is", persona.surname)
//		fmt.Println("my name is", wwork.person, "my work is", wwork.work)
//		wwork.say()
//	}
//
//	type person struct {
//		name     string
//		username string
//		stack    []string
//	}
//
//	type personality struct {
//		person
//		work string
//	}
//
//	func (p person) say() {
//		fmt.Printf("my name is %s,my username is %s and my stack is %s\n", p.name, p.username, p.stack)
//	}
//
//	func (p personality) talk() {
//		fmt.Printf("my name is %s,my username is %s and my stack is %s,i woring in %s sphere", p.name, p.username, p.stack, p.work)
//	}
//
//	func main() {
//		firstPerson := person{
//			name:     "Stas",
//			username: "we34ed",
//			stack:    []string{"golang", "git", "swag"},
//		}
//		secondPerson := personality{
//			person: person{"Slava", "Rudge", []string{"dota", "cs:go", "github"}},
//			work:   "gaming",
//		}
//		firstPerson.say()
//		secondPerson.talk()
//	}
//
// TODO:эмбеддинг структуры в структуры через "append"
//
//	type Thing struct {
//		name   string
//		weight int
//	}
//
//	type Room struct {
//		name   string
//		things []Thing
//	}
//
//	func main() {
//		thing := Thing{
//			name:   "карандаш",
//			weight: 50,
//		}
//		room := Room{}
//		room.name = "Библиотека"
//		room.things = append(room.things, thing)
//		fmt.Println(room)
//	}
//
//	type Thing struct {
//		name   string
//		weight int
//	}
//
//	type Room struct {
//		name   string
//		things []Thing
//	}
//
//	func main() {
//		thing := Thing{
//			name:   "Fridge",
//			weight: 25,
//		}
//		room := Room{}
//		room.name = "kitchen"
//		room.things = append(room.things, thing)
//		fmt.Printf("в комнате %s стоит %s,который весит %dкг", room.name, thing.name, thing.weight)
//	}
//
//	type Thing struct {
//		name   string
//		weight int
//	}
//
//	type Room struct {
//		name   string
//		things Thing
//	}
//
//	func main() {
//		thing := Thing{
//			name:   "шкаф",
//			weight: 40,
//		}
//		room := Room{
//			name:   "кухня",
//			things: Thing{thing.name, thing.weight},
//		}
//		fmt.Println(room)
//	}
//
//	type Thing struct {
//		name   string
//		weight int
//	}
//
//	type Room struct {
//		name  string
//		thing []Thing
//	}
//
//	func main() {
//		thing := Thing{
//			name:   "духовка",
//			weight: 25,
//		}
//		room := Room{
//			name: "кухня",
//		}
//		room.thing = append(room.thing, thing)
//		fmt.Println(room)
//	}
//TODO:реализацция структур через слайс[]
//type Thing struct {
//	name   string
//	weight int
//}
//type Room struct {
//	name   string
//	things []Thing
//}
//
//func main() {
//
//	//	room := Room{
//	//		name: "баня",
//	//		things: []Thing{
//	//			{name: "печь", weight: 70},
//	//			{name: "стол", weight: 5},
//	//			{name: "ванная", weight: 45},
//	//		},
//	//	}
//	//	fmt.Println(room)
//	//}
//
//TODO:Реализация структур через слайс[]

//		room := Room{
//			name: "кухня",
//			things: []Thing{
//				{name: "стол", weight: 10},
//				{name: "духовка", weight: 45},
//			},
//		}
//		fmt.Println(room)
//	}
//
//	type People struct {
//		name string
//		age  int
//		work string
//	}
//
//	type Person struct {
//		people []People
//	}
//
//	func main() {
//		person := Person{
//			people: []People{
//				{name: "Stas", age: 21, work: "backend"},
//				{name: "Danya", age: 23, work: "frontend"},
//				{name: "Egor", age: 24, work: "fullstack"},
//			},
//		}
//		fmt.Println(person)
//	}
//
// TODO:квиз
//
//	type Thing struct {
//		name   string
//		weight int
//	}
//
//	type Room struct {
//		name  string
//		thing []Thing
//	}
//
//	type House struct {
//		name  string
//		rooms [][]Room
//	}
//
//	func main() {
//		house := House{
//			name: "Нахапино",
//			rooms: [][]Room{
//				{
//					{name: "Кухня", thing: []Thing{
//						{name: "стол", weight: 5},
//						{name: "духовка", weight: 45},
//					}},
//					{name: "Баня", thing: []Thing{
//						{name: "Печь", weight: 75},
//						{name: "Статуя", weight: 190},
//					}},
//					{name: "Спальня", thing: []Thing{
//						{name: "Кровать", weight: 50},
//						{name: "Комод", weight: 25},
//					}},
//				},
//			},
//		}
//		fmt.Println(house)
//	}
//type Thing struct {
//	name   string
//	weight int
//}

//	type Room struct {
//		name  string
//		thing []Thing
//	}
//
//	type House struct {
//		name string
//		room [][]Room
//	}
//
//	func main() {
//		house := House{
//			name: "Нахапино",
//			room: [][]Room{
//				{
//					{name: "Гостинная", thing: []Thing{
//						{"Диван", 120},
//						{"Телевизор", 35},
//					}},
//					{name: "Спальня", thing: []Thing{
//						{name: "Кровать", weight: 52},
//						{name: "Стол", weight: 30},
//					}},
//					{name: "Кухня", thing: []Thing{
//						{name: "Холодильник", weight: 45},
//						{name: "Духовка", weight: 57},
//					}},
//				},
//			},
//		}
//		fmt.Println(house)
//	}
//TODO:вложенные структуры можно инициализировать и без указания их параметров:"name","age" и т.д.
//type Item struct {
//	name   string
//	weight int
//}
//type Invent struct {
//	name string
//	item []Item
//}
//
//func main() {
//	invent := Invent{
//		name: "рюкзак",
//		item: []Item{
//			{"Нож", 300},
//			{"Зажигалка", 150},
//			{"Костюм", 2000},
//		},
//	}
//	fmt.Println(invent)
//}

// TODO:Операция сравнения(==)
//
//	type Thing struct {
//		name   string
//		weight int
//	}
//
//	func main() {
//		axe := Thing{
//			name:   "Топор",
//			weight: 2500,
//		}
//		knife := Thing{
//			name:   "Нож",
//			weight: 1500,
//		}
//		fmt.Println(axe == knife)
//		fmt.Println(knife == Thing{name: "Нож", weight: 1500})
//		fmt.Println(axe == Thing{name: "Топор", weight: 2500})
//	}
//
//	type Item struct {
//		name   string
//		weight int
//	}
//
//	type Room struct {
//		name string
//		item []Item
//	}
//
//	type House struct {
//		name string
//		room [][]Room
//	}
//
//	func main() {
//		house := House{
//			name: "Nahapino",
//			room: [][]Room{
//				{
//					{"Кухня", []Item{
//						{"Духовка", 52},
//						{"Холодильник", 64},
//					}},
//					{"Гостинная", []Item{
//						{"Телевизор", 25},
//						{"Диван", 47},
//					}},
//					{"Спальня", []Item{
//						{"Кровать", 24},
//						{"Комод", 22},
//					}},
//				},
//			},
//		}
//		fmt.Println(house)
//	}
//type Item struct {
//	name   string
//	weight int
//}
//type Room struct {
//	name string
//	item []Item
//}
//type House struct {
//	name string
//	room [][]Room
//}
//
//func main() {
//	house := House{
//		name: "Нахаппппано",
//		room: [][]Room{
//			{
//				{name: "Гостинная", item: []Item{
//					{"Диван", 35},
//					{"Телевизор", 24},
//				}},
//				{name: "Кухня", item: []Item{
//					{"Холодильник", 52},
//					{"Стол", 21},
//				}},
//				{name: "Спальня", item: []Item{
//					{name: "Кровать", weight: 46},
//					{name: "Шкаф", weight: 74},
//				}},
//			},
//		},
//	}
//	fmt.Println(house)
//}

//TODO:Методы,повторение
//Отличие метода от функции в том
//что метод используется с конкретным типом данным,вызываются только экземпляром данного типа данных.

//type Hero struct {
//	name   string
//	health int
//}
//
//// 'h'-приемник у которого тип - структура "Hero"
//// метод "introduce()" использует поля структуры с помощью приемника
//func (h Hero) introduce() string {
//	return "My name is " + h.name
//}
//
//func main() {
//	hero := Hero{
//		name:   "Pudge",
//		health: 7465,
//	}
//	fmt.Println(hero.introduce())
//}

// методы с параметрами(записываются в скобках после имени метода)
//type Hero struct {
//	name   string
//	health int
//}
//
//func (h Hero) say() string {
//	return "Меня зовут " + h.name
//}
//
//// '*' перед структурой - дает возможность запоминать все изменения значений полей в передаваемой структуре)
//func (h *Hero) upHealth(points int) {
//	fmt.Println("Количество здоровья героя", h.name, "было", h.health)
//	h.health += points
//	fmt.Println("Количество здоровья героя", h.name, "стало", h.health)
//}
//func main() {
//	hero := Hero{
//		name:   "Пудж",
//		health: 7645,
//	}
//	fmt.Println(hero.say())
//	hero.upHealth(55)
//}

//TODO:Квиз

//type Hero struct {
//	name   string
//	health int
//}
//
//func (h Hero) say() string {
//	return "Меня зовут " + h.name
//}
//func (h Hero) upHealth(points int) {
//	fmt.Println("У героя", h.name, "было", h.health, "здоровья")
//	h.health += points
//	fmt.Println("У героя", h.name, "стало", h.health, "здоровья")
//}
//func (h Hero) moveTo(place string) {
//	fmt.Println("Герой", h.name, "переместился в", place)
//}
//func main() {
//	hero := Hero{
//		name:   "Пудж",
//		health: 7450,
//	}
//	fmt.Println(hero.say())
//	hero.upHealth(50)
//	hero.moveTo("мид")
//}
//

// TODO:Разница между методами и функциями
//type Hero struct {
//	name    string
//	baseDam int
//	baseDef int
//	minDam  int
//	minDef  int
//	maxDam  int
//	maxDef  int
//}
//
//func randint(min, max int) int {
//	rand.Seed(time.Now().UnixNano())
//	return rand.Intn(max-min+1) + min
//}
//
//func (h Hero) attack() string {
//	return fmt.Sprintf("%s нанес урон, равный %d", h.name, h.baseDam+randint(h.minDam, h.maxDam))
//}
//func (h Hero) defence() string {
//	return fmt.Sprintf("%s заблокировал урон, равный %d", h.name, h.baseDef+randint(h.minDef, h.maxDef))
//}
//func main() {
//	pudge := Hero{
//		name:    "пудж",
//		baseDam: 12,
//		baseDef: 10,
//		minDam:  7,
//		minDef:  5,
//		maxDam:  25,
//		maxDef:  23,
//	}
//	pudgeAttack := pudge.attack()
//	pudgeDefence := pudge.defence()
//	fmt.Println(pudgeAttack)
//	fmt.Println(pudgeDefence)
//}

//TODO:Квиз

//type Person struct {
//	name string
//	age  int
//}
//
//func (p Person) sayName() string {
//	return "Меня зовут " + p.name
//}
//func (p Person) sayAge() int {
//	return p.age
//}
//func main() {
//	stas := Person{
//		name: "Стас",
//		age:  21,
//	}
//	fmt.Println(stas.sayName())
//	fmt.Println("Мне исполнился", stas.sayAge(), "год")
//}

//TODO:Эмбединг в Go
//встраивание типов в структуру

//type Character struct{
//	name string
//	className string
//}
//type Mage struct{
//	Character   //использовали использовали предыдущую структуру
//	def int
//	attack int
//}
//

//type User struct {
//	name  string
//	email string
//}
//
//func (u User) info() string {
//	return fmt.Sprintf("имя: %s\nпочта: %s", u.name, u.email)
//}
//
//func (u User) notify() {
//	fmt.Println("Уведомить пользователя")
//}
//
//type Admin struct {
//	User
//	level string
//}
//
//func main() {
//	user := User{
//		name:  "Стас",
//		email: "stas.kuchma.86@mail.ru",
//	}
//	admin := Admin{
//		User:  user,
//		level: "group",
//	}
//	admin.name = "Ваня"
//	admin.email = "dogmeboss@mail.ru"
//	fmt.Printf("имя админа:%s\nпочта админа:%s\nуровень админа:%s\n", admin.name, admin.email, admin.level)
//	fmt.Println(admin.info())
//	admin.notify()
//	fmt.Printf("имя пользователя:%s\nпочта пользователя:%s\n", user.name, user.email)
//	fmt.Println(user.info())
//	user.notify()
//}

// TODO:Множественное встраивание
//type User struct {
//	name string
//}
//type Manager struct {
//	title string
//}
//type Person struct {
//	User         //встроили первую структуру
//	Manager      //встроили вторую структуру
//}
//
//func main() {
//	user := User{
//		name: "Stas",
//	}
//	manager := Manager{
//		title: "Boss",
//	}
//	person := Person{
//		User:    user,
//		Manager: manager,
//	}
//	fmt.Println(person.name, person.title)
//}

//TODO:Цепочка встраиваний

//	type Item struct {
//		name   string
//		weight int //создал первую структуру
//	}
//
//	type Room struct {
//		name string
//		item []Item //встроил первую структуру во вторую
//	}
//
//	type House struct {
//		name string
//		room [][]Room //встроил вторую структуру (в которой встроена первая структура) в третью
//	}
//
//	func main() {
//		house := House{
//			name: "Поместье Кузьминых",
//			room: [][]Room{
//				{
//					{name: "Кабинет", item: []Item{
//						{name: "Комьютер", weight: 30},
//						{name: "Бутылка Виски", weight: 05},
//					}},
//					{name: "Спальня", item: []Item{
//						{name: "Кровать", weight: 50},
//						{name: "Шкаф", weight: 32},
//					}},
//					{name: "Кухня", item: []Item{
//						{name: "Холодильник", weight: 52},
//						{name: "Стол", weight: 21},
//					}},
//				},
//			},
//		}
//		fmt.Println(house)
//	}
//type Hero struct {
//	health  int
//	damage  int
//	defence int
//}
//type Mage struct {
//	name string
//	hero Hero
//}
//type Warrior struct {
//	name string
//	hero Hero
//}
//
//func (h Hero) attack() string {
//	return fmt.Sprintf("Герой нанес урон равный: %d\n", h.damage)
//}
//func (h Hero) def() string {
//	return fmt.Sprintf("Герой заблокировал %d урона\n", h.defence)
//}
//func (m Mage) attack() string {
//	return fmt.Sprintf("Герой %s нанес урон равный: %d\n", m.name, m.hero.damage)
//}
//func (m Mage) def() string {
//	return fmt.Sprintf("Герой %s заблокировал %d урона\n", m.name, m.hero.defence)
//}
//func (w Warrior) attack() string {
//	return fmt.Sprintf("Герой %s нанес урон равный: %d\n", w.name, w.hero.damage)
//}
//func (w Warrior) def() string {
//	return fmt.Sprintf("Герой %s заблокировал %d урона\n", w.name, w.hero.defence)
//}
//func main() {
//	hero := Hero{
//		health:  100,
//		damage:  25,
//		defence: 20,
//	}
//	fmt.Println(hero.attack())
//	fmt.Println(hero.def())
//
//	mage := Mage{
//		name: "Lloyd",
//		hero: Hero{100, 20, 25},
//	}
//	fmt.Println(mage.attack())
//	fmt.Println(mage.def())
//
//	warrior := Warrior{
//		name: "Axe",
//		hero: Hero{120, 30, 23},
//	}
//	fmt.Println(warrior.attack())
//	fmt.Println(warrior.def())
//}
