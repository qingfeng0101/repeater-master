package handler

import (
	"encoding/json"
	"net/http"

	"github.com/kexirong/repeater/models"
	"github.com/kexirong/repeater/sender"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
)

func ReadNotifyRules(c echo.Context) error {
	id := c.Param("id")
	if id != "" {
		var nr models.NotifyRule
		if err := models.DB.Preload(clause.Associations).First(&nr, id).Error; err != nil {
			return err
		}
		return c.JSON(http.StatusOK, nr)
	}

	var nrs []models.NotifyRule
	if err := models.DB.Unscoped().Preload(clause.Associations).Find(&nrs).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, nrs)
}

func ReadNotifyRulesCache(c echo.Context) error {
	cache := sender.NotifyRuleRepos().Cache()

	return c.JSON(http.StatusOK, cache)
}

func ReadNotifyRulesIndexCacheTree(c echo.Context) error {
	rs := sender.NotifyRuleRepos()

	return c.JSON(http.StatusOK, rs.IndexCache.ToMap())
}

type ruleForm struct {
	models.NotifyRule
	ContactIDs  []int `json:"contact_ids" validate:"required"`
	SenderSetID int   `json:"sender_set_id" validate:"required"`
}

func AddNotifyRule(c echo.Context) error {
	var rf ruleForm

	if err := c.Bind(&rf); err != nil {
		return err
	}
	nr := rf.NotifyRule
	if nr.Name == "" {
		nr.Name = "default"
		nr.LabelSetList = []byte("[]")
	}

	var cs []models.Contact
	if err := models.DB.Find(&cs, rf.ContactIDs).Error; err != nil {
		return err
	}
	var ss models.SenderSet
	if err := models.DB.First(&ss, rf.SenderSetID).Error; err != nil {
		return err
	}
	nr.Contacts = cs
	nr.SenderSet = &ss

	if err := sender.NotifyRuleRepos().Regist(&nr); err != nil {
		return err
	}

	if err := models.DB.Create(&nr).Error; err != nil {
		sender.NotifyRuleRepos().Delete(sender.RuleName(nr.Name))
		return err
	}

	return c.JSON(http.StatusCreated, echo.Map{"id": nr.ID})

}

func UpdateNotifyRule(c echo.Context) error {
	param := c.Param("id")

	var nr models.NotifyRule
	if err := models.DB.First(&nr, param).Error; err != nil {
		return err
	}

	var rf ruleForm
	if err := c.Bind(&rf); err != nil {
		return err
	}

	var cs []models.Contact
	if err := models.DB.Find(&cs, rf.ContactIDs).Error; err != nil {
		return err
	}
	var ss models.SenderSet
	if err := models.DB.First(&ss, rf.SenderSetID).Error; err != nil {
		return err
	}

	nr.SenderFilter = rf.SenderFilter
	// nr.LabelSetList = rf.LabelSetList
	// nr.Contacts = cs

	nr.SenderSet = &ss
	drule := sender.Rule{}

	if err := json.Unmarshal(nr.SenderSet.Sets, &drule.SenderSet); err != nil {
		return err
	}

	if err := json.Unmarshal(rf.LabelSetList, &drule.LabelSetList); err != nil {
		return err
	}

	if len(drule.LabelSetList) > 0 {
		nr.LabelSetList = rf.LabelSetList
	}

	if err := json.Unmarshal(nr.SenderFilter, &drule.SenderFilter); err != nil {
		return err
	}
	for _, contact := range cs {
		drule.Contacts = append(drule.Contacts, sender.ContactUsername(contact.Username))
	}

	if err := sender.NotifyRuleRepos().Update(sender.RuleName(nr.Name), &drule); err != nil {
		return err
	}

	// if err := models.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&nr).Error; err != nil {

	if err := models.DB.Save(&nr).Error; err != nil {
		return err
	}

	if err := models.DB.Model(&nr).Association("Contacts").Replace(cs); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{"id": nr.ID})

}

func DeleteNotifyRule(c echo.Context) error {
	param := c.Param("id")

	var nr models.NotifyRule
	result := models.DB.First(&nr, param)
	if result.RowsAffected == 0 {
		return result.Error
	}

	if err := models.DB.Delete(&nr).Error; err != nil {
		return err
	}
	sender.NotifyRuleRepos().Delete(sender.RuleName(nr.Name))
	return c.NoContent(http.StatusNoContent)

}

func UnscopedDeleteNotifyRule(c echo.Context) error {
	param := c.Param("id")

	var nr models.NotifyRule
	result := models.DB.Unscoped().First(&nr, param)
	if result.RowsAffected == 0 {
		return result.Error
	}

	if err := models.DB.Unscoped().Select("Contacts").Delete(&nr).Error; err != nil {
		return err
	}
	sender.NotifyRuleRepos().Delete(sender.RuleName(nr.Name))
	return c.NoContent(http.StatusNoContent)

}
