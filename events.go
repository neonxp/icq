package icq

type CommonEvent struct {
	Type   string `json:"type"`
	SeqNum int    `json:"seqNum"`
}

type ServiceEvent struct {
	CommonEvent
	Data interface{} `json:"eventData"`
}

type BuddyListEvent struct {
	CommonEvent
	Data struct {
		Groups []struct {
			Name    string  `json:"name"`
			ID      int     `json:"id"`
			Buddies []Buddy `json:"buddies"`
		} `json:"groups"`
	} `json:"eventData"`
}

type MyInfoEvent struct {
	CommonEvent
	Data Buddy `json:"eventData"`
}

type TypingStatus string

const (
	StartTyping TypingStatus = "typing"
	StopTyping               = "none"
)

type TypingEvent struct {
	CommonEvent
	Data struct {
		AimID        string       `json:"aimId"`
		TypingStatus TypingStatus `json:"typingStatus"`
	} `json:"eventData"`
}

type IMEvent struct {
	CommonEvent
	Data struct {
		Autoresponse int    `json:"autoresponse"`
		Timestamp    int    `json:"timestamp"`
		Notification string `json:"notification"`
		MsgID        string `json:"msgId"`
		IMF          string `json:"imf"`
		Message      string `json:"message"`
		RawMessage   struct {
			IPCountry     string `json:"ipCountry"`
			ClientCountry string `json:"clientCountry"`
			Base64Msg     string `json:"base64Msg"`
		} `json:"rawMsg"`
		Source Buddy `json:"source"`
	} `json:"eventData"`
}
