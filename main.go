package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/mrjones/oauth"

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
	//validates the flag argument
	if *tweetTag == "" {
		log.Fatalln(`tweet empty, exiting`)
	} else if len(*tweetTag) > 280 {
		log.Fatalln(`sorry, more than 280 characteres`)
	}
	//prints the flag content to the console for debugging purpouse
	fmt.Println("tweet: ", *tweetTag)

	//model for receive cfg keys
	newSet := models.Key{}
	//gets home env key
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Panicln(err)
	}
	//check if app folder exist
	if _, err := os.Stat(homedir + "/.go-birdie"); os.IsNotExist(err) {
		fmt.Println("dir does not exist, creating it...")
		//creates folders to store data
		os.MkdirAll(homedir+"/.go-birdie", 0777)
	}
	//opens the cfg.txt file, later implementations will use to retain the auth tokens
	//if does no exist, create the file using the rw permission
	file, err := os.OpenFile(homedir+"/.go-birdie/cfg.key", os.O_CREATE, 0600)
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
	//config new costumer usign the api keys
	costumer := oauth.NewConsumer(newSet.APIKey, newSet.APISecretKey, oauth.ServiceProvider{
		RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
		AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
		AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
	})
	//config new user using token keys
	user := oauth.AccessToken{
		Token:  newSet.AccessToken,
		Secret: newSet.AccessTokenSecret,
	}
	//generate a client using costumer and user keys
	client, err := costumer.MakeHttpClient(&user)
	if err != nil {
		log.Println(err)
	}
	//send post request to twitter using keys
	res, err := client.PostForm("https://api.twitter.com/1.1/statuses/update.json", url.Values{"status": []string{*tweetTag}})
	//prints status code
	fmt.Println(res.StatusCode)
}
