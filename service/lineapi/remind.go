package messaging

import (
	handler "avah/common/http"
	"avah/constant"
	"avah/service/lineapi/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func Remind() {
	requestHeader := map[string]string{
		constant.CONTENT_TYPE: "application/x-www-form-urlencoded",
	}

	message := model.Message{
		Type: constant.MESSAGE_TYPE_TEXT,
		Text: "Hi! from Avah",
	}

	requestBody := model.PushMessage{
		To: "",
		Messages: []model.Message{
			message,
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatal(err)
	}

	lineClient := handler.HttpClient{
		Method:        http.MethodPost,
		Url:           viper.GetString("line.client.oauth.url"),
		RequestHeader: requestHeader,
		RequestBody:   jsonData,
	}

	log.Fatalf("url: %s", lineClient.Url)
}
