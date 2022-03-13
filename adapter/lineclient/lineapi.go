package adapter

import (
	handler "avah/common/http"
	"avah/constant"
	"avah/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

func Oauth() {
	oauthUrl := viper.GetString("line.client.oauth.url.verify")
	accessToken := viper.GetString("line.access.token")

	url := fmt.Sprintf("%s?%s=%s",
		oauthUrl,
		constant.ACCESS_TOKEN,
		accessToken)

	lineClient := handler.HttpClient{
		Method: http.MethodGet,
		Url:    url,
	}

	fmt.Println("url: %s", lineClient.Url)

	lineClient.Request()

	fmt.Println(lineClient.ResponseBody)
}

func Remind(textMessage string, to string) {
	accessToken := viper.GetString("line.access.token")
	requestHeader := map[string]string{
		constant.CONTENT_TYPE:  "application/json",
		constant.AUTHORIZATION: fmt.Sprintf("Bearer %s", accessToken),
	}

	message := model.Message{
		Type: constant.MESSAGE_TYPE_TEXT,
		Text: textMessage,
	}

	requestBody := &model.PushMessage{
		To: to,
		Messages: []model.Message{
			message,
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Errorf("convert request body failed: ", err)
	}

	lineClient := handler.HttpClient{
		Method:        http.MethodPost,
		Url:           viper.GetString("line.client.message.url.push"),
		RequestHeader: requestHeader,
		RequestBody:   jsonData,
	}

	lineClient.Request()
}
