package main

import "fmt"

func main() {
	dna := "aaccccccddekkkkaaakk"

	count := 0
	sum := 0
	letter := string(dna[0])
	letterWinner := string(dna[0])
	letterIndex := 0

	for i, current := range dna {
		if string(current) == letter {
			sum++
		} else {
			letter = string(current)
			sum = 1
		}

		if sum > count {
			count = sum
			letterWinner = string(current)
			letterIndex = i - count + 1
		}
	}

	fmt.Println(count)
	fmt.Println(letterWinner)
	fmt.Println(letterIndex)

	slice := map[string]int{}
	sequences := []string{}
	aux := string(dna[0])
	sequence := ""
	size := len(dna)
	count = 0
	letterWinner = string(dna[0])
	
	firstLetter := func(letter string) string {
		return string(letter[0])
	}
	
	for i, v := range dna {
		if string(v) == aux {
			sequence += string(v)
		}
		
		if string(v) != aux || size == i+1 {
			sequences = append(sequences, sequence)
			aux = string(v)		
			sequence = string(v)
		}
	}

	for _, v := range sequences {
		if len(v) > slice[firstLetter(v)] {
			slice[firstLetter(v)] = len(v)
		}
	}

	for key, value := range slice {
		if value > count {
			count = value
			letterWinner = key
		}
	}

	fmt.Println("String:",dna)
	fmt.Println("List:", sequences)
	fmt.Println("Map:", slice)

	fmt.Println("Letter winner:", letterWinner)
	fmt.Println("Amount:", count)
}