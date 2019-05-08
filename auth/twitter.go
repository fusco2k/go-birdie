package auth

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/mrjones/oauth"
)

//Authenticate a new user
func Authenticate() (ck, cs, at, as string) {
	//creates a new scanner for reading user input from CLI
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("enter the twitter api key:")
	scanner.Scan()
	consumerKey := scanner.Text()
	fmt.Println("enter the twitter api key secret:")
	scanner.Scan()
	consumerSecret := scanner.Text()
	//creates a new consumer to start the auth process
	c := oauth.NewConsumer(
		consumerKey,
		consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		},
	)
	//generate the get request using oob for command line
	//receive the token and the url to continue the process
	token, requestURL, err := c.GetRequestTokenAndUrl("oob")
	if err != nil {
		log.Fatal(err)
	}
	//shows the url to user auth with twitter
	fmt.Println(requestURL)
	//asks the pin confirmation
	scan := bufio.NewScanner(os.Stdin)
	fmt.Println("enter the twitter validation code:")
	scan.Scan()
	verificationCode := scan.Text()
	//process the auth token
	tk, err := c.AuthorizeToken(token, verificationCode)
	if err != nil {
		log.Fatal(err)
	}
	//return the codes to generate the client and the conf file
	return consumerKey, consumerSecret, tk.Token, tk.Secret
}
