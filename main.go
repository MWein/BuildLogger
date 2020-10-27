package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

const sc = "\u001B7"
const rc = "\u001B8"

func Menu(header string, options []string) int {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(rc + sc)
		fmt.Printf("\n\n")
		fmt.Println(header + ":")
		fmt.Println("")
	
		for index, option := range options {
			fmt.Printf("%d: %s\n", index+1, option)
		}
	
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// Remove the new line crap
		text = strings.Replace(text, "\r", "", -1)
		text = strings.Replace(text, "\n", "", -1)
	
		input, err := strconv.Atoi(text)
	
		if err != nil || input > len(options) || input <= 0 {
			fmt.Println("Invalid Input")
			continue
		}

		return input - 1
	}
}

func main() {
	fmt.Println(" --------------------------")
	fmt.Println("|                          |")
	fmt.Println("|     Build Log Editor     |")
	fmt.Println("|     Weinberg Software    |")
	fmt.Println("|                          |")
	fmt.Println(" --------------------------")

	hello := Menu("Select Project", []string{"Zenith Cruzer", "Vans RV-9"})

	fmt.Println(hello)
	// Loop until exit

}
