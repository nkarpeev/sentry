package sentryClient

type SentryClient interface {
	Send(payload string) error
}
