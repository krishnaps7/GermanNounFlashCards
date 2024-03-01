package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type GermanNoun struct {
	Noun    string
	Meaning string
	Learned bool
}

func main() {
	// Open the file
	file, err := os.Open("GermanVocab.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Slice to store German nouns with gender
	germanNouns := make([]GermanNoun, 0)
	// i := 0
	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		fields := strings.Fields(line)
		germanNouns = append(germanNouns, GermanNoun{fmt.Sprintf("%v %v", fields[2], fields[3]), fields[1], false})
		// fmt.Printf("%v : %v\n", germanNouns[i].Noun, germanNouns[i].Meaning)
		// i++
		// time.Sleep(time.Second)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	var input string
	var repeat string
	for {
		start := rand.Intn(700)
	comeback:
		fmt.Print("\033[H\033[2J\n") //Clears entire screen
		for j := start; j < start+5; j++ {
			fmt.Print("\r")
			fmt.Printf("%v", germanNouns[j].Noun)
			fmt.Scanf("%s", &input)
			fmt.Printf("%v", germanNouns[j].Meaning)
			time.Sleep(time.Second)
			fmt.Print("\033[2K") //clears line
			fmt.Print("\r")      // bring cursor	 back to beginning of the line
			if input == "y" {
				germanNouns[j].Learned = true
			}
			// goto comeback
		}
		fmt.Print("\nrepeating(enter for yes or anyother for no)?")
		fmt.Scanf("%s", &repeat)
		if repeat == "" {
			goto comeback
		}
	}
}
