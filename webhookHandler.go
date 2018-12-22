package icq

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetWebhookHandler returns http.HandleFunc that parses webhooks
// Warning! Not fully functional at ICQ now!
func (a *API) GetWebhookHandler(cu chan<- Update, e chan<- error) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if r.Method != http.MethodPost {
			e <- fmt.Errorf("incorrect method: %s", r.Method)
			return
		}
		wr := &WebhookRequest{}
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			e <- err
			return
		}
		if err := json.Unmarshal(b, wr); err != nil {
			e <- err
			return
		}
		for _, u := range wr.Updates {
			cu <- u
		}
	}
}
