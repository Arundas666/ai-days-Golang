package openAI

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
)
const (
	ApiEndpoint = "https://api.openai.com/v1/chat/completions"
)
func GetAnswerFromOpenAI(question string, apiKey string) (string, error) {
	client := resty.New()
	response, err := client.R().
		SetAuthToken(apiKey).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"model":    "gpt-3.5-turbo",
			"messages": []map[string]interface{}{{"role": "user", "content": question}},
		}).
		Post(ApiEndpoint)
	if err != nil {
		return "", err
	}

	var data map[string]interface{}
	err = json.Unmarshal(response.Body(), &data)
	if err != nil {
		return "", err
	}

	content := data["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	return content, nil

}
