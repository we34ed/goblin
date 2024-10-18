package main

import "fmt"

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
//	type warrior struct {
//		name     string
//		username string
//	}
//
//	type mage struct {
//		warrior
//		damage int
//	}
//
//	func (w warrior) say() {
//		fmt.Println(w.name, w.username)
//	}
//
//	func (m mage) say() {
//		fmt.Printf("%s %s\n%d\n", m.name, m.username, m.damage)
//	}
//
//	func (m mage) talk(s string) {
//		fmt.Println(s)
//	}
//
//	func main() {
//		stas := warrior{
//			name:     "Stas",
//			username: "Kuchma",
//		}
//		stas.say()
//		vanya := mage{
//			warrior: warrior{name: "Slava", username: "Rudan"},
//			damage:  120,
//		}
//		vanya.say()
//		vanya.talk("i kill yoy")
//	}
type person struct {
	name string
	age  int
}
type botan struct {
	person
	work string
}

func (p person) talk() {
	fmt.Printf("my name is %s\n my age is %d\n", p.name, p.age)
}
func (b botan) say() {
	fmt.Printf("my name is %s\n my age is %d\n im working in %s\n", b.name, b.age, b.work)
}
func main() {
	firstPerson := person{
		name: "Stas",
		age:  21,
	}
	secondPerson := botan{
		person: person{name: "Slavik", age: 21},
		work:   "doter",
	}
	firstPerson.talk()
	secondPerson.say()
}
