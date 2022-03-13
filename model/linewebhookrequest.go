package model

type LineWebHookRequest struct {
	Destination string `json:"destination"`
	Events      []struct {
		Type    string `json:"type"`
		Message struct {
			Type string `json:"type"`
			ID   string `json:"id"`
			Text string `json:"text"`
		} `json:"message"`
		Timestamp  int64  `json:"timestamp"`
		Source     Source `json:"source"`
		ReplyToken string `json:"replyToken"`
		Mode       string `json:"mode"`
	} `json:"events"`
}

type Source struct {
	Type   string `json:"type"`
	UserID string `json:"userId"`
}
