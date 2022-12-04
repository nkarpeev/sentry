package telegramSentryClient

type Config struct {
	ApiToken string `toml:"api_key"`
	ChatId   string `toml:"chat_id"`
	LogLevel string `toml:"log_level"`
}

func NewConfig() *Config {
	return &Config{
		ApiToken: "", //todo
		ChatId:   "0",
		LogLevel: "debug",
	}
}
