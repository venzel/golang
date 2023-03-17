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

func ToJson[T Animal](t *T) *bytes.Buffer {
	jsonByte, err := json.Marshal(*t)

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

	polvo := Animal{"Povo", 4}
	polvoToJson := ToJson(&polvo)
	fmt.Println(polvoToJson)

	tatuArgs := `{"name":"Tatu","size":10}`
	var tatu Animal
	ToObject(&tatuArgs, &tatu)
	fmt.Println(tatu)

	griloArgs := `{"name":"Grilo","size":20,"color":"red"}`
	var grilo Insect
	ToObject(&griloArgs, &grilo)
	fmt.Println(grilo)
}
