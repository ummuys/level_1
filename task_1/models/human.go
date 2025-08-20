package models

import "fmt"

type human struct {
	Name string
	Age  int
}

func (h human) Greetings() {
	fmt.Printf("Hello! My name is %s!\nI am %d years old :)", h.Name, h.Age)
}
