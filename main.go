package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fusco2k/go-birdie/models"
)

func main() {
	var cfg string
	//create a flag to interact with the user and get the tweet message
	tweetTag := flag.String("t", "", `tweets the message under ""`)
	//parse the flags to use the flags content
	flag.Parse()
	//crash the execution if using more than 1 argument
	if flag.NArg() > 0 {
		log.Fatalln(`the flag "t" only accepts 1 argument, check usage and use "" if necessary`)
	}
	//prints the flag content to the console for debugging purpouse
	fmt.Println("tweet: ", *tweetTag)

	//model for receive cfg keys
	newSet := models.Key{}
	//opens the cfg.txt file, later implementations will use to retain the auth tokens
	//if does no exist, create the file using the rw permission
	file, err := os.OpenFile("cfg.key", os.O_CREATE, 0600)
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
	//decodes the key from cfg file
	err = json.Unmarshal([]byte(cfg), &newSet)
	if err != nil {
		fmt.Println(err.Error())
	}
}
