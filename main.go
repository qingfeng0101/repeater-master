package main

import (
	"fmt"
	"github.com/kexirong/repeater/config"
	"github.com/kexirong/repeater/dingtalk"
	"github.com/kexirong/repeater/email"
	"github.com/kexirong/repeater/lark"
	"net/http"

	//"github.com/kexirong/repeater/email"
	"github.com/kexirong/repeater/logging"
	"github.com/kexirong/repeater/models"
	"github.com/kexirong/repeater/router"
	"github.com/kexirong/repeater/sender"
	"github.com/kexirong/repeater/wecom"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func main() {
	// db, err := gorm.Open(mysql.Open("uommp:Pmmou180!@tcp(192.168.2.180:3306)/ommp"), &gorm.Config{})
	conf := config.Get()
	if conf.DataPath == ""|| conf.DataPath == " " {
		conf.DataPath = "data.db"
	}
	db, err := gorm.Open(sqlite.Open(conf.DataPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	models.AutoMigrate(db)
	models.DB = db

	registAllRender(models.DB)
	registAllNotifyRule(models.DB)
	createDefaultAccount(models.DB)

	e := echo.New()
	// e.Use(em.LoggerWithConfig(em.LoggerConfig{Output: logging.LoggerWriter()}))
	e.Logger.SetLevel(log.INFO)
	e.Logger.SetOutput(logging.LoggerWriter())
	e.HTTPErrorHandler = customHTTPErrorHandler

	e.Use(middleware.CORS())

	router.AddOld(e)
	router.AddApi(e)


	e.Start(conf.Listen)

}

func customHTTPErrorHandler(err error, c echo.Context) {
	c.Logger().Error(err)
	if c.Response().Committed {
		return
	}

	he, ok := err.(*echo.HTTPError)
	if !ok {
		he = &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	if c.Request().Method == http.MethodHead {
		err = c.NoContent(he.Code)
	} else {
		err = c.JSON(he.Code, echo.Map{"reason": he.Message})
	}
	if err != nil {
		c.Logger().Error(err)
	}
}

func registAllRender(db *gorm.DB) {
	var scs []models.StmpCredential
	if err := db.Find(&scs).Error; err != nil {
		panic(err)
	}

	for _, sc := range scs {
		email.StmpCredentialRegist(sc)
	}

	var wcs []models.WecomCredential
	if err := db.Find(&wcs).Error; err != nil {
		panic(err)
	}

	for _, wc := range wcs {
		wecom.WecomCredentialRegist(wc)
	}
	var dcs []models.DingCredential
	if err := db.Find(&dcs).Error; err != nil {
		panic(err)
	}

	for _, dc := range dcs {
		dingtalk.WecomCredentialRegist(dc)
	}
	var lcs []models.LarkCredential
	if err := db.Find(&lcs).Error; err != nil {
		panic(err)
	}

	for _, lc := range lcs {
		lark.LarkCredentialRegist(lc)
	}
}

func registAllNotifyRule(db *gorm.DB) {
	var nrs []models.NotifyRule
	if err := db.Preload(clause.Associations).Find(&nrs).Error; err != nil {
		panic(err)
	}

	for _, nr := range nrs {
		if err := sender.NotifyRuleRepos().Regist(&nr); err != nil {
			fmt.Println(err)
		}

	}

}

func createDefaultAccount(db *gorm.DB) {
	user := models.User{
		Name:     "Admin",
		Username: "admin",
		Password: "admin123",
	}

	if err := db.FirstOrCreate(&user, "username = ?", "admin").Error; err != nil {
		panic(err)
	}
}
