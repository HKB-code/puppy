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

func main() {
	fmt.Println(dog.WhenGrownUp(Bark()))
	fmt.Println(dog.WhenGrownUp(Barks()))
}
