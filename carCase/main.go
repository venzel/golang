package main

import (
	"fmt"
	model "model/auto"
)

func main() {
    auto := model.Auto{
        Name: "Auto x",
        Year: 2001,
    }

    car := model.Car{
        Auto: auto,
        Color: "Red",
    }

    car.SetColor("Yellow")

    fmt.Println(car)
    fmt.Println(car.GetColor())
}