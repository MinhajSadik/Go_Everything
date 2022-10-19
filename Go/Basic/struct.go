package Basic

import (
	"fmt"
)

type Programmers struct {
	id    int
	name  string
	pType string
	field string
}

func displayInfo(programmer Programmers) {
	fmt.Printf("id:> %v, name:> %v, pType:> %v, field:> %v\n", programmer.id, programmer.name, programmer.pType, programmer.field)
}

func (p *Programmers) changeName(name string) {
	p.name = name
}

func Struct() {
	fmt.Println("Struct Learning...")

	minhaj := Programmers{1, "minhaj", "backend", "cse"}

	displayInfo(minhaj)

	minhaj.changeName("rahman")

	fmt.Println(minhaj.name)

}
