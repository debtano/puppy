package puppy

import (
	"fmt"

	"github.com/debtano/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof Woof Woof"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func from10() {
	fmt.Println("i am from version 1.0")
}
