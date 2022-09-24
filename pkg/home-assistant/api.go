package haas

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type HomeAssistantAPI struct {
	Url   string
	Token string
}

type ResponsePost []string

func NewHomeAssistantAPI(url, token string) *HomeAssistantAPI {
	return &HomeAssistantAPI{url, token}
}

func (api *HomeAssistantAPI) Post(path string, data interface{}) (bool, error) {
	url := fmt.Sprintf("%s/%s", api.Url, path)
	json_data, err := json.Marshal(data)

	if err != nil {
		return false, err
	}

	req, _ := http.NewRequest("POST", url, bytes.NewReader(json_data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", api.Token))

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
		return false, err
	}

	if res.StatusCode >= 300 {
		log.Fatalf("status code: %d", res.StatusCode)
		return false, fmt.Errorf("status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)

	var response *ResponsePost
	json.Unmarshal(body, &response)

	defer res.Body.Close()

	if err != nil {
		log.Fatal(err)
		return false, err
	}

	return true, nil
}
