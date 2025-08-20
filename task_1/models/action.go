package models

type Action struct {
	human
}

func NewAction(name string, age int) *Action {
	return &Action{human: human{Name: name, Age: age}}
}
