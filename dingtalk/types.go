package dingtalk

type SendMess struct {
	WebHook string `json:"webhook"`
	AgentId uint `json:"agent_id"`
	UseridList string `json:"userid_list"`
	DeptIdList string `json:"dept_id_list"`
	AccessToken string `json:"accessToken"`
	Mobiles
	Msg Msg `json:"msg"`
}
type Mobiles struct {
	Mobile string `json:"mobile"`
	SupportExclusiveAccountSearch bool `json:"support_exclusive_account_search"`
}
type Msg struct {
	Msgtype string `json:"msgtype"`
	Markdown MarkDown `json:"markdown"`
	Text Text `json:"text"`
}

type MarkDown struct {
	Title string `json:"title"`
	Text string `json:"text"`
}
type Text struct {
	Content string `json:"content"`
}
