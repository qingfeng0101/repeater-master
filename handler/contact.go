package handler

import (
	"errors"
	"net/http"

	"github.com/kexirong/repeater/models"
	"github.com/kexirong/repeater/sender"
	"github.com/kexirong/repeater/wecom"
	"github.com/labstack/echo/v4"
)

func ReadContacts(c echo.Context) error {
	id := c.Param("id")
	if id != "" {
		var contact models.Contact
		if err := models.DB.First(&contact, id).Error; err != nil {
			return err
		}
		return c.JSON(http.StatusOK, contact)
	}

	var contacts []models.Contact
	if err := models.DB.Unscoped().Find(&contacts).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, contacts)
}

func AddContact(c echo.Context) error {
	var contact models.Contact
	err := c.Bind(&contact)
	if err != nil {
		return err
	}

	if err := models.DB.Create(&contact).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, echo.Map{"id": contact.ID})
}

func DeleteContact(c echo.Context) error {
	param := c.Param("id")
	var contact models.Contact
	result := models.DB.Unscoped().First(&contact, param)
	if result.RowsAffected == 0 {
		return result.Error
	}

	if err := models.DB.Unscoped().Delete(&contact).Error; err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)

}

func UpdateContact(c echo.Context) error {
	param := c.Param("id")

	var contact models.Contact
	result := models.DB.First(&contact, param)
	if result.RowsAffected == 0 {
		return result.Error
	}

	err := c.Bind(&contact)
	if err != nil {
		return err
	}

	if err := models.DB.Save(&contact).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, echo.Map{"id": contact.ID})
}

func PostSyncContact(c echo.Context) error {

	s := sender.SenderRepos().Get("wecom-default")
	wc, ok := (s).(*wecom.WeCom)
    if !ok {
		return errors.New("wecom-default is nil, not *wecom.WeCom")
	}

		departments, err := wc.GetDepartmentList()
		if err != nil {
			return err
		}

		var users []wecom.User
		for _, department := range departments {
			if department.ParentID != 1 {
				continue
			}

			us, err := wc.GetUserList(department.ID, true)
			if err != nil {
				return err
			}
			users = append(users, us...)
		}

		for _, u := range users {
			contact := models.Contact{
				Username: u.UserID,
				WeCom:    u.UserID,
				Name:     u.Name,
				Mobile:   u.Mobile,
				Email:    u.Email,
				DingTalk: u.UserID,
			}
			models.DB.Create(&contact)
		}
		return c.NoContent(http.StatusAccepted)






}
