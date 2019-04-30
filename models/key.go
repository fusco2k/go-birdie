package models

//Key basic struct to unmarshall data
type Key struct {
	APIKey            string `json:"api_key,omitempty"`
	APISecretKey      string `json:"api_secret_key,omitempty"`
	AccessToken       string `json:"access_token,omitempty"`
	AccessTokenSecret string `json:"access_token_secret,omitempty"`
}
