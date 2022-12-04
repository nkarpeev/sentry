package telegramSentryClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Client struct {
	config *Config
	logger *log.Logger
}

func New(config *Config) *Client {
	return &Client{
		config: config,
		logger: log.New(),
	}
}

func (c *Client) Send(msg string) error {
	var err error
	var response *http.Response
	var url = fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", c.config.ApiToken)

	body, _ := json.Marshal(map[string]string{
		"chat_id": c.config.ChatId,
		"text":    msg,
	})

	response, err = http.Post(
		url,
		"application/json",
		bytes.NewBuffer(body),
	)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	log.Infof("Message '%s' was sent", msg)
	log.Infof("Response JSON: %s", string(body))

	return nil
}
