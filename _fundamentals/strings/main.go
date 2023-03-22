package main

import (
	"fmt"
	"strings"
)

func main() {
	names := make([]string, 0)

	names = append(names, "tiago")
	names = append(names, "marcos")
	names = append(names, "cintia")

	for _, name := range names {
		println(name)
	}

	all := strings.Join(names, ", ") // tiago, marcos, cintia

	fmt.Println(all)

	exists := strings.Contains(all, "Tiago")

	fmt.Println(exists) // False

	index := strings.Index(all, "marcos")

	fmt.Println(index) // 7

	word := "Taigo vai brincar de bola"

	parts := strings.Split(word, " ")

	fmt.Println(parts) // [Taigo vai brincar de bola]

	for _, part := range parts {
		fmt.Println(part)
	}

	fmt.Println(len(word)) // 25

	namex := "tiago"
	namey := "toago"

	fmt.Println(strings.ToUpper(namex)) // TIAGO
	fmt.Println(strings.ToUpper(namey)) // TIAGO

	final := strings.Split(namex, "")

	fmt.Println(final) // [t i a g o]

	fmt.Println(strings.Replace(namey, "o", "x", 1)) // txago
	fmt.Println(strings.ReplaceAll(namey, "o", "x")) // txagx

	nameLoop := "marcos"

	for _, letter := range nameLoop {
		fmt.Println(string(letter))
	}

	namec := "marcos" + "boigia"

	fmt.Println(namec)

	namesx := ""
	namesx += namex
	namesx += namey

	fmt.Println(namesx)

	fnl := fmt.Sprintf("%s%s", namex, namey)

	fmt.Println(fnl)

	vni := 10.2

	fnx := fmt.Sprintf("%f", vni)

	fmt.Println(fnx)
}
