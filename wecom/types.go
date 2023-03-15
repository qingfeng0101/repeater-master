package wecom

type BaseMsg struct {
	ToUser  string `json:"touser,omitempty"`
	ToParty string `json:"toparty,omitempty"`
	MsgType string `json:"msgtype"`
	AgentID uint   `json:"agentid"`
	// TextCard textCard `json:"textcard"`
}

type TextCard struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type Text struct {
	Content string `json:"content"`
}

type Markdown Text

type Result struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
