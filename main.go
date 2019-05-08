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
	//gets home env key
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Panicln(err)
	}
	//------handling tag------
	var cfg string
	//create a flag to interact with the user and get the tweet message
	tweetTag := flag.String("t", "", `tweets the message under ""`)
	//parse the flags to use the flags content
	flag.Parse()
	//aks if wish to authenticate
	if flag.NArg() == 0 {
		fmt.Println("do you wish to start a new authentication process? (y/n)")
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("enter the twitter api key")
		scanner.Scan()
		authDecision := scanner.Text()
		switch authDecision {
		case "y":
			//check if app folder exist
			checkDirectory(homedir)
			//create a new cfg.key
			createCfg(homedir)
			//prints the result
			fmt.Println(`now you are authenticated, next time use the -t followed by the "...tweet text...", exiting`)
			os.Exit(0)
		case "n":
			fmt.Println(`ok, next time use the -t followed by the "...tweet text...", exiting`)
			os.Exit(0)
		default:
			fmt.Println("sorry, thats not an option, exiting now.")
			os.Exit(0)
		}
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
	//check if app folder exist
	if !checkDirectory(homedir) {
		log.Fatalln("you are not authenticated, start the program using no flags")
	}
	//opens the cfg.txt file
	//if does no exist, create the file using the rw permission
	file, err := os.OpenFile(homedir+"/.go-birdie/cfg.key", os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("could not open the cfg.key: %v", err)
	}
	defer file.Close()
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

	//------handling auth ------

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

func checkDirectory(homedir string) bool {
	if _, err := os.Stat(homedir + "/.go-birdie"); os.IsNotExist(err) {
		fmt.Println()
		//creates folders to store data
		err := os.MkdirAll(homedir+"/.go-birdie", 0777)
		if err != nil {
			log.Println(err)
		}
		return false
	}
	return true
}

func createCfg(homedir string) {
	//model for receive cfg keys
	keySet := models.Key{}
	//remove the older cfg.key, if exists, and create a new one
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

func dummyKey() (ck, cs, at, as string){
ck = "testeck"
cs = "testecs"
at = "testeat"
as = "testeas"

return ck, cs, at, as
}
