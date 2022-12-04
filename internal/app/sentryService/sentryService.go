package sentryService

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/nkarpeev/telegram-logger/internal/app/sentryClient"
	"github.com/nkarpeev/telegram-logger/internal/app/sentryClient/telegramSentryClient"
)

var (
	configPath string
	sender     sentryClient.SentryClient
)

type Payload struct {
	Msg string
}

func init() {
	flag.StringVar(&configPath, "tg-config-path", "configs/telegramClient.toml", "path to config file")
	flag.Parse()

	config := telegramSentryClient.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	sender = telegramSentryClient.New(config)
}

func Write(payload Payload) error {
	err := sender.Send(payload.Msg)

	if err != nil {
		return err
	}

	return nil
}

// func Write(msg string) error { //todo add interface

// 	time.Sleep(time.Duration(5) * time.Second)

// 	f, err := os.Create("telegram-logger.txt")

// 	if err != nil {
// 		panic(err)
// 	}

// 	defer f.Close()

// 	_, err = f.WriteString(msg)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return err
// }
