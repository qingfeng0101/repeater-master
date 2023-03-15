package handler

import (
	"net/http"
	"time"

	"github.com/kexirong/repeater/middleware"
	"github.com/kexirong/repeater/models"
	"github.com/labstack/echo/v4"
)

func ReadAPIKeys(c echo.Context) error {
	id := c.Param("id")
	if id != "" {
		var apiKey models.APIKey
		if err := models.DB.First(&apiKey, id).Error; err != nil {
			return err
		}
		return c.JSON(http.StatusOK, apiKey)
	}

	var apiKeys []models.APIKey
	if err := models.DB.Find(&apiKeys).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, apiKeys)
}

func AddAPIKey(c echo.Context) error {
	var apiKey models.APIKey
	err := c.Bind(&apiKey)
	if err != nil {
		return err
	}
	apiKey.Key = middleware.GenAPIKey(middleware.Csha256)

	if err := models.DB.Create(&apiKey).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, echo.Map{"id": apiKey.ID})
}

func DeleteAPIKey(c echo.Context) error {
	param := c.Param("id")

	var apiKey models.APIKey
	result := models.DB.First(&apiKey, param)
	if result.RowsAffected == 0 {
		return result.Error
	}

	if err := models.DB.Delete(&apiKey).Error; err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)

}

func RefreshAPIKey(c echo.Context) error {
	param := c.Param("id")

	var apiKey models.APIKey
	result := models.DB.First(&apiKey, param)
	if result.RowsAffected == 0 {
		return result.Error
	}

	if err := models.DB.Model(&apiKey).Update("updated_at", time.Now().Unix()).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, echo.Map{"id": apiKey.ID})
}
