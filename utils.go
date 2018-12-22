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
	cmd := string(parts[0][1:])
	return &Command{
		From:      event.Data.Source.AimID,
		Command:   strings.ToLower(cmd),
		Arguments: parts[1:],
	}, true
}
