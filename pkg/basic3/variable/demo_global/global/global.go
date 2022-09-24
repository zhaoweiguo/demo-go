package global

type Person struct {
	Name string
	Age  int
}

var person Person = Person{"zwg", 123}

func Get() Person {
	return person
}

func Do1() {
	person = Person{
		Name: "gordon",
		Age:  100,
	}
}

func Do2() {
	person.Age = 99
}

func Do3() {
	person = Person{
		Name: "xinxi",
		Age:  90,
	}
}
