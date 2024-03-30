package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func foo() error {
	var err *os.PathError = nil
	return err
}

const (
	a = iota + 1
	_
	b
	c
)

func change(a *int) {
	t := *a * 2
	a = &t
}

type model struct {
	Name string `json:"test"`
	Age  int    `json:"age"`
}

func main() {
	data := `{"name":"Alex","age":30}`
	var model model
	err := json.Unmarshal([]byte(data), &model)
	if err != nil {
		panic(err)
	}
	fmt.Println(model)
}
