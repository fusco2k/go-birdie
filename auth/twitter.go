package auth

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/mrjones/oauth"
)

var at, as string

var c *oauth.Consumer

//Authenticate a new user
func Authenticate(key, secret string) (at, as string) {

	consumerKey := "yDsQeHa6P2tfOwA1lHiWGIgCZ"

	consumerSecret := "yIo6AsfmD82OyJpvMwVD7MftqPxVxmgS3eZHPqVTHWl9EEu4hP"

	c = oauth.NewConsumer(
		consumerKey,
		consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		},
	)

	token, requestURL, err := c.GetRequestTokenAndUrl("oob")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(requestURL)
	scan := bufio.NewScanner(os.Stdin)
	fmt.Println("enter the twitter api key")
	scan.Scan()
	verificationCode := scan.Text()

	tk, err := c.AuthorizeToken(token, verificationCode)
	if err != nil {
		log.Fatal(err)
	}

	at = tk.Token
	fmt.Println(at)
	as = tk.Secret
	fmt.Println(as)

	return at, as
}
