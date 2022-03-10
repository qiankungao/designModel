package main

import "fmt"

type API interface {
	Say(name string) string
}

func NewApi(name string) API {
	switch name {
	case "chinese":
		return &Chinese{}
	case "English":
		return &English{}
	}
	return nil
}

type Chinese struct {
}

func (c *Chinese) Say(name string) string {
	return name + "你好"
}

type English struct {
}

func (e *English) Say(name string) string {
	return "hello " + name
}
func main() {
	fmt.Println(NewApi("chinese").Say("乾坤"))
	fmt.Println(NewApi("English").Say("乾坤"))
}
