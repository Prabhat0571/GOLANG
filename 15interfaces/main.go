package main

import "fmt"

type dog struct{}

func (d dog) say() string {
	return "woof"

}

type cat struct{}

func (c cat) say() string {
	return "meow"
}

func main() {
	c := cat{}
	fmt.Println("Cat says: ", c.say())
}
