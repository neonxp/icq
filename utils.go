package icq

import (
	"strings"
)

// Command is sugar on top of IMEvent that represented standard ICQ bot commands
type Command struct {
	From      string
	Command   string
	Arguments []string
}

// ParseCommand from IMEvent
// Command must starts from '.' or '/'. Arguments separated by space (' ')
func ParseCommand(event *IMEvent) (*Command, bool) {
	message := event.Data.Message
	parts := strings.Split(message, " ")
	if len(parts) == 0 {
		return nil, false
	}
	if parts[0][0] != '.' && parts[0][0] != '/' {
		return nil, false
	}

	return &Command{
		From:      event.Data.Source.AimID,
		Command:   string(parts[0][1:]),
		Arguments: parts[1:],
	}, true
}
