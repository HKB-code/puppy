package puppy

import (
	"fmt"

	"github.com/HKB-code/dog"
)

func Bark() string {
	return "woof"
}

func Barks() string {
	return "woofs woof woo"
}

func BigBark(){
fmt.Println(dog.WhenGrownUp(Bark()))
}
func BigBarks() {
	
	fmt.Println(dog.WhenGrownUp(Barks()))
}
