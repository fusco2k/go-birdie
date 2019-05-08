### Go-Birdie

It's a simple CLI twitter app with the main function of posting tweets from a simple command line.

There is also a way to authenticate from the app, you gonna need the consumers keys that I will not provide, but you can generate your own creating an app from the developer console on https://developer.twitter.com/.

## Instructions

Simple executing the app will thrown you into the authentication process and the app will aks if you want to generate a new cfg.key file where the app store the key values.

You then will have to re-run the app using now `go-birdie -t "...tweet goes here"`.
The app will not accept more than 1 argument, an empty tweet or flag without argument.

If you try to tweet 2 times the same thing, Twitter will not accept and will return an error.

### Always remember, do not share your customers keys, if you do so, have in mind that only do with trusted people and that you will have the risk of it leaking on the web. I take no responsibility on misusage. 