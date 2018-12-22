package icq

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// SendMessage with `message` text to `to` participant
func (a *API) SendMessage(message Message) (*MessageResponse, error) {
	parse, _ := json.Marshal(message.Parse)
	v := url.Values{}
	v.Set("aimsid", a.token)
	v.Set("r", strconv.FormatInt(time.Now().Unix(), 10))
	v.Set("t", message.To)
	v.Set("message", message.Text)
	v.Set("mentions", strings.Join(message.Mentions, ","))
	if len(message.Parse) > 0 {
		v.Set("parse", string(parse))
	}
	b, err := a.send("/im/sendIM", v)
	if err != nil {
		return nil, err
	}
	r := &Response{}
	if err := json.Unmarshal(b, r); err != nil {
		return nil, err
	}
	if r.Response.StatusCode != 200 {
		return nil, fmt.Errorf("failed to send message: %s", r.Response.StatusText)
	}
	return r.Response.Data, nil
}
