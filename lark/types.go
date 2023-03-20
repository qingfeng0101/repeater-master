package lark

type Markdown struct {
	Tag string `json:"tag"`
	Content string `json:"content"`
}

type Element struct {
	Elements []Markdown `json:"elements"`
}
type SendMsg struct {
	Card Element `json:"card"`
	MsgType string `json:"msg_type"`
	UserIds []string `json:"user_ids"`
	Content Texts `json:"content"`
}
type Texts struct {
	Text string `json:"text"`
}