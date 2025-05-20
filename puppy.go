package puppy

import (
	"fmt"

	"github.com/ClaudioRoncaglio/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func From11() {
	fmt.Println("Sono alla versione 1.1.0")
}

func From12() {
	fmt.Println("Sono alla versione 1.2.0")
}
