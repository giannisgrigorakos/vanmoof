package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"vanmoof/internal/mexicanWave"
)

func main() {
	fmt.Print("Enter the word that you want to Mexican Wave and press Enter⏎: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred while reading input. Please try again", err)
		return
	}

	input = strings.TrimSuffix(input, "\n") // remove the pressed enter from the end
	result := mexicanWave.Wave(input)
	fmt.Println(result)
}
