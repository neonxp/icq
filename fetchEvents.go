package icq

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func (a *API) FetchEvents(ctx context.Context, ch chan interface{}) error {
	fetchResp := &struct {
		Response struct {
			Data struct {
				FetchBase string            `json:"fetchBaseURL"`
				PollTime  int               `json:"pollTime"`
				Events    []json.RawMessage `json:"events"`
			} `json:"data"`
		} `json:"response"`
	}{}
	for {
		b := []byte{}
		u := a.fetchBase
		if u == "" {
			v := url.Values{}
			v.Set("aimsid", a.token)
			v.Set("first", "1")
			u = a.baseUrl + "/fetchEvents?" + v.Encode()
		}
		req, err := http.Get(u)
		if err != nil {
			return err
		}
		b, err = ioutil.ReadAll(req.Body)
		req.Body.Close()
		if err := json.Unmarshal(b, fetchResp); err != nil {
			return err
		}
		a.fetchBase = fetchResp.Response.Data.FetchBase
		for _, e := range fetchResp.Response.Data.Events {
			ce := &CommonEvent{}
			if err := json.Unmarshal(e, ce); err != nil {
				return err
			}
			switch ce.Type {
			case "service":
				ev := &ServiceEvent{}
				if err := json.Unmarshal(e, ev); err != nil {
					return err
				}
				ch <- ev
			case "buddylist":
				ev := &BuddyListEvent{}
				if err := json.Unmarshal(e, ev); err != nil {
					return err
				}
				ch <- ev
			case "myInfo":
				ev := &MyInfoEvent{}
				if err := json.Unmarshal(e, ev); err != nil {
					return err
				}
				ch <- ev
			case "typing":
				ev := &TypingEvent{}
				if err := json.Unmarshal(e, ev); err != nil {
					return err
				}
				ch <- ev
			case "im":
				ev := &IMEvent{}
				if err := json.Unmarshal(e, ev); err != nil {
					return err
				}
				ch <- ev
			default:
				ch <- ce
			}
		}
		select {
		case <-time.After(time.Duration(fetchResp.Response.Data.PollTime)):
		case <-ctx.Done():
			return nil
		}
	}
	return nil
}
