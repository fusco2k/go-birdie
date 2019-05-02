package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/fusco2k/go-birdie/auth"
	"github.com/fusco2k/go-birdie/models"
	"github.com/mrjones/oauth"
)

func main() {

	//------handling tag------
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

	//------handling tag end------

	//------handling keys------

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
		err := os.MkdirAll(homedir+"/.go-birdie", 0777)
		if err != nil {
			log.Println(err)
		}

		file, err := os.OpenFile(homedir+"/.go-birdie/cfg.key", os.O_CREATE|os.O_RDWR, 0777)
		if err != nil {
			log.Println(err)
		}
		file.WriteString(`{"authenticate":"false"}`)
		file.Sync()
		file.Close()
	}
	//opens the cfg.txt file, later implementations will use to retain the auth tokens
	//if does no exist, create the file using the rw permission
	file, err := os.OpenFile(homedir+"/.go-birdie/cfg.key", os.O_RDWR, 0666)
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
	fmt.Println(newSet)
	//------handling keys end------
	if newSet.Authenticated == false {
		file.Close()
		os.Remove(homedir + "/.go-birdie/cfg.key")
		file, err := os.OpenFile(homedir+"/.go-birdie/cfg.key", os.O_CREATE|os.O_RDWR, 0777)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer file.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
		//creates a new scanner for reading user input from CLI
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("enter the twitter api key")
		scanner.Scan()
		tk := scanner.Text()
		fmt.Println("enter the twitter api key secret")
		scanner.Scan()
		ts := scanner.Text()
		//authenticate the user
		at, as := auth.Authenticate(tk, ts)
		//generate the key set
		newSet.APIKey = tk
		newSet.APISecretKey = ts
		newSet.Authenticated = true
		newSet.AccessToken = at
		newSet.AccessTokenSecret = as
		//write to the cfg.key
		file.WriteString(`{
			"authenticated":true,
			"api_key":"` + newSet.APIKey + `",
			"api_secret_key":"` + newSet.APISecretKey + `",
			"access_token":"` + newSet.AccessToken + `",
			"access_token_secret":"` + newSet.AccessTokenSecret + `"}`)
		file.Sync()
	}
	fmt.Println(newSet)
	//------handling auth------

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

	//------handling auth end------

	//------handling tweet------

	//send post request to twitter using keys
	res, err := client.PostForm("https://api.twitter.com/1.1/statuses/update.json", url.Values{"status": []string{*tweetTag}})
	//prints status code
	fmt.Println(res.StatusCode)

	//------handling tweet end------
}
