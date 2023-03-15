package models

type Contact struct {
	BaseModel
	Username string `gorm:"column:username; type:varchar(31); not null; uniqueIndex" json:"username"` //用户名
	Name     string `gorm:"column:name; type:varchar(31); not null" json:"name"`                      //姓名
	Email    string `gorm:"column:email; type:varchar(127);" json:"email"`                            //邮箱
	Mobile   string `gorm:"column:mobile; type:varchar(31)" json:"mobile"`                            //手机号
	WeCom    string `gorm:"column:wecom; type:text" json:"wecom"` //企业微信id
	DingTalk string `gorm:"column:dingtalk; type:text" json:"dingtalk"`    //钉钉id
	Lark string `gorm:"column:lark; type:text" json:"lark"`    //钉钉id
}

func (Contact) TableName() string {
	return "contact"
}

type User struct {
	BaseModel
	Name     string `gorm:"column:name; type:varchar(31); not null" json:"name"`                      //姓名
	Username string `gorm:"column:username; type:varchar(31); not null; uniqueIndex" json:"username"` //用户名
	Password string `gorm:"column:password; type:varchar(31)" json:"password"`                        //密码
	Email    string `gorm:"column:email; type:varchar(127);" json:"email"`                            //邮箱
	Mobile   string `gorm:"column:mobile; type:varchar(31)" json:"mobile"`                            //手机号
}

func (User) TableName() string {
	return "user"
}

type APIKey struct {
	BaseModel
	Name        string `gorm:"column:name; type:varchar(31); not null" json:"name"` //名称
	Key         string `gorm:"column:key; type:varchar(127);" json:"key"`
	ExpireDelta uint   `gorm:"column:expire_delta" json:"expire_delta"`
}

func (APIKey) TableName() string {
	return "api_key"
}
