package checkup

import (
	"encoding/json"
	"fmt"

	"checkup/notifier/discord"
	"checkup/notifier/mail"
	"checkup/notifier/mailgun"
	"checkup/notifier/pushover"
	"checkup/notifier/slack"
)

func notifierDecode(typeName string, config json.RawMessage) (Notifier, error) {
	switch typeName {
	case mail.Type:
		return mail.New(config)
	case slack.Type:
		return slack.New(config)
	case mailgun.Type:
		return mailgun.New(config)
	case pushover.Type:
		return pushover.New(config)
	case discord.Type:
		return discord.New(config)
	default:
		return nil, fmt.Errorf(errUnknownNotifierType, typeName)
	}
}
