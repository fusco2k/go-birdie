package models

//Key basic struct to unmarshall data
type Key struct {
	APIKey       string `json:"APIKey,omitempty"`
	APISecretKey string `json:"APISecretKey,omitempty"`
	AccessToken string `json:"AccessToken,omitempty"`
	AccessTokenSecret string `json:"AccessTokenSecret,omitempty"`
}