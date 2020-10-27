package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Menu(header string, options []string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("\n\n")
	fmt.Println(header + ":")
	fmt.Println("")

	for index, option := range options {
		fmt.Printf("%d: %s\n", index+1, option)
	}

	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	// Remove the new line
	text = strings.Replace(text, "\n", "", -1)

	fmt.Println(text)

}

func main() {
	fmt.Println(" --------------------------")
	fmt.Println("|                          |")
	fmt.Println("|     Build Log Editor     |")
	fmt.Println("|     Weinberg Software    |")
	fmt.Println("|                          |")
	fmt.Println(" --------------------------")

	Menu("Select Project", []string{"Zenith Cruzer", "Vans RV-9"})

	// Loop until exit

}
