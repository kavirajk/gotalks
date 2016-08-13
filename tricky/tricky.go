package tricky

import (
	"fmt"
	"os"
)

type Notifier interface {
	Notify(message string)
}

type SlackNotifier struct {
}

func (s *SlackNotifier) Notify(message string) {

}

var notifier = SlackNotifier{}

func mustEnv(key string) string {
	var value string
	if value = os.Getenv(key); value == "" {
		notifier.Notify(fmt.Sprintf("%s env variable not set.", key))
		os.Exit(1)
	}
	return value
}
