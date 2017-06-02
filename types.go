package main

// Action is user initiated behavior passed by Slack
type Action struct {
	Actions []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"actions"`
	CallbackID string `json:"callback_id"`
	Team       struct {
		ID     string `json:"id"`
		Domain string `json:"domain"`
	} `json:"team"`
	Channel struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"channel"`
	User struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"user"`
	ActionTs  string `json:"action_ts"`
	MessageTs string `json:"message_ts"`
	// AttachmentID string `json:"attachment_id"`
	Token       string `json:"token"`
	ResponseURL string `json:"response_url"`
}
