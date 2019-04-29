package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func main() {
	var cfg string
	//opens the cfg.txt file, later implementations will use to retain the auth tokens
	//if does no exist, create the file using the rw permission
	file, err := os.OpenFile("cfg.txt", os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	//creates a new scanner to read the cfg file
	s := bufio.NewScanner(file)
	//loops scanning each new line and concatenates to the cfg string
	for s.Scan() {
		str := s.Text()
		cfg += str + " "
	}
	//prints the file output
	fmt.Println(cfg)

	//creates a new scanner to read user input from CLI
	scanner := bufio.NewScanner(os.Stdin)
	//loops scanning each new line and prints to the console
	for scanner.Scan() {
		_, err := fmt.Println(scanner.Text())
		if err != nil {
			fmt.Fprintln(os.Stderr, "err: ", err)
		}
	}

	//http.ListenAndServe(":8080", nil)
}
