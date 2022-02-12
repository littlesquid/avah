package messaging

import (
	handler "avah/common/http"
	"avah/constant"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func Oauth() {

	oauthUrl := viper.GetString("line.client.oauth.url")
	accessToken := viper.GetString("line.access.token")

	url := fmt.Sprintf("%s?%s=%s",
		oauthUrl,
		constant.ACCESS_TOKEN,
		accessToken)

	lineClient := handler.HttpClient{
		Method: http.MethodGet,
		Url:    url,
	}

	log.Fatalf("url: %s", lineClient.Url)
}
