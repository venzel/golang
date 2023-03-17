package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Animal struct {
	Name string `json:"name"`
	Size int    `json:"size"`
}

type Insect struct {
	Name  string `json:"name"`
	Size  int    `json:"size"`
	Color string `json:"color"`
}

func ToJson(a *Animal) *bytes.Buffer {
	jsonByte, err := json.Marshal(*a)

	if err != nil {
		log.Fatal()
	}

	return bytes.NewBuffer(jsonByte)
}

func ToObject[T Animal | Insect](args *string, t *T) {
	if err := json.Unmarshal([]byte(*args), &t); err != nil {
		log.Fatal()
	}
}

func main() {
	peixe := Animal{"Peixe", 3}
	peixeToJson := ToJson(&peixe)
	fmt.Println(peixeToJson)

	tatuArgs := `{"name":"Cago","size":10}`
	var cago Animal
	ToObject(&tatuArgs, &cago)
	fmt.Println(cago.Name)

	griloArgs := `{"name":"Grilo","size":20,"color":"red"}`
	var grilo Insect
	ToObject(&griloArgs, &grilo)
	fmt.Println(grilo.Color)
}
