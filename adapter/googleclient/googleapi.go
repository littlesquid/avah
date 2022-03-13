package adapter

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GoogleApiLogin() *http.Client {
	fmt.Println("connecting to google api")

	b, err := ioutil.ReadFile("./credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.JWTConfigFromJSON(b, viper.GetString("google.sheet.oauth"))
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := config.Client(oauth2.NoContext)

	return client

}
