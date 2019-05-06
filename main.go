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
	//show instructions if no arg is used
	if flag.NArg() == 0 {
		log.Fatalln(`use -a for start a new authentication process or -t to tweet`)
	}
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
	keySet := models.Key{}
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
	err = json.Unmarshal([]byte(cfg), &keySet)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(keySet)
	//------handling keys end------
	if keySet.Authenticated == false {
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
		//authenticate the user
		keySet.APIKey, keySet.APISecretKey, keySet.AccessToken, keySet.AccessTokenSecret = auth.Authenticate()
		keySet.Authenticated = true
		//write to the cfg.key
		file.WriteString(
			`{"authenticated":true,
			"api_key":"` + keySet.APIKey + `",
			"api_secret_key":"` + keySet.APISecretKey + `",
			"access_token":"` + keySet.AccessToken + `",
			"access_token_secret":"` + keySet.AccessTokenSecret + `"}`)
		file.Sync()
	}
	fmt.Println(keySet)
	//------handling auth------

	//config new costumer usign the api keys
	costumer := oauth.NewConsumer(keySet.APIKey, keySet.APISecretKey, oauth.ServiceProvider{
		RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
		AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
		AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
	})
	//config new user using token keys
	user := oauth.AccessToken{
		Token:  keySet.AccessToken,
		Secret: keySet.AccessTokenSecret,
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
