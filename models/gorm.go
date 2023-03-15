package models

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) {
	// db.Migrator().DropTable(&SenderSet{}, &NotifyRule{})
	err := db.AutoMigrate(
		&WecomCredential{},
		&DingCredential{},
		&LarkCredential{},
		&StmpCredential{},
		&SenderSet{},
		&Contact{},
		&User{},
		&NotifyRule{},
		&PrometheusDataSource{},
		&APIKey{},
		&PrometheusFakeDataSource{},
	)
	if err != nil {
		panic(err)
	}
}
