package main

// Player struct
type Player struct {
	Name   string
	Level  int
	Health int
}

type Football struct {
	Person Player
}

func main() {
	p1 := Player{
		Name:   "P1",
		Level:  1,
		Health: 10,
	}

	println(p1.Name)

	// anonyms struct
	sample := struct {
		name string
	}{name: "Hello"}

	sample.name = "World"
	println(sample.name)

	p2 := Football{
		Person: p1,
	}

	println(p2.Person.Level)
}
