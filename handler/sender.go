package handler

import (
	"encoding/json"
	"net/http"

	"github.com/kexirong/repeater/models"
	"github.com/kexirong/repeater/sender"
	"github.com/labstack/echo/v4"
)

func ReadSenders(c echo.Context) error {
	list := []map[string]string{}
	for k, v := range sender.SenderRepos().List() {
		list = append(list, map[string]string{"index": k, "type": v.Type(), "hash": v.Hash()})
	}
	return c.JSON(http.StatusOK, list)
}

func ReadSenderSets(c echo.Context) error {

	id := c.Param("id")
	if id != "" {
		var set models.SenderSet
		if err := models.DB.First(&set, id).Error; err != nil {
			return err
		}
		return c.JSON(http.StatusOK, set)
	}

	var sets []models.SenderSet
	if err := models.DB.Find(&sets).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, sets)

}

func AddSenderSet(c echo.Context) error {

	var set models.SenderSet

	if err := c.Bind(&set); err != nil {
		return err
	}
	if set.Name == "" {
		set.Name = "default"
	}

	if set.Name != "default" {
		var count int64
		models.DB.Model(&models.SenderSet{}).Where("name = ?", "default").Count(&count)
		//if count == 0 {
		//	return echo.NewHTTPError(http.StatusBadRequest, "please create a default entry first")
		//}
	}

	if err := models.DB.Create(&set).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{"id": set.ID})
}
func DeleteSenderSet(c echo.Context) error {
	param := c.Param("id")
	var set models.SenderSet
	result := models.DB.Unscoped().First(&set, param)
	if result.RowsAffected == 0 {
		return result.Error
	}

	if err := models.DB.Unscoped().Delete(&set).Error; err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
func UpdateSenderSet(c echo.Context) error {
	param := c.Param("id")

	var set models.SenderSet
	result := models.DB.First(&set, param)
	if result.RowsAffected == 0 {
		return result.Error
	}

	err := c.Bind(&set)
	if err != nil {
		return err
	}

	if err := models.DB.Save(&set).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, echo.Map{"id": set.ID})
}

func AddSend(c echo.Context) error {
	name := c.Param("name")
	var payload Payload
	if err := c.Bind(&payload); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if name == "" {
		name = "default"
	}
	var ss models.SenderSet
	result := models.DB.Where("name = ?", name).First(&ss)
	if result.RowsAffected == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, name+" not exist")
	}
	// if result.Error != nil {
	// 	return result.Error
	// }

	var sets []string
	if err := json.Unmarshal(ss.Sets, &sets); err != nil {
		c.Logger().Error(err)
		return err
	}
	var contacts []models.Contact
	result = models.DB.Where("username IN ?", payload.To).Find(&contacts)
	if result.RowsAffected == 0 || result.Error != nil {
		c.Logger().Error(result.Error)
		return result.Error
	}
	for _, set := range sets {
		s := sender.SenderRepos().Get(set)
		err := s.Send(contacts, payload.Subject, payload.ContentType, payload.Content)
		if err != nil {
			c.Logger().Error(err)
		}
	}
	return c.NoContent(http.StatusAccepted)
}
