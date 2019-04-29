package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//creates a new scanner for reading user input from CLI
	scanner := bufio.NewScanner(os.Stdin)

	//loops scanning each new line and prints to the console
	for scanner.Scan() {
		_, err := fmt.Println(scanner.Text())
		if err != nil {
			fmt.Fprintln(os.Stderr, "err: ", err)
		}
	}
}
