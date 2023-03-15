package handler

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/kexirong/repeater/middleware"
	"github.com/kexirong/repeater/models"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	var user models.User
	if err := models.DB.Where("username = ?", username).First(&user).Error; err != nil {

		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	if password != user.Password {
		return echo.ErrUnauthorized
	}

	claims := &middleware.JWTCustomClaims{
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Mobile:   user.Mobile,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(middleware.SigningKey))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"access_token": t,
		"token_type":   "Bearer",
	})
}

type chPasswordModel struct {
	Password        string `json:"password" validate:"required"`
	NewPassword     string `json:"new_password"  validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

func ChangeUserPassword(c echo.Context) error {
	username := c.Param("username")

	var model chPasswordModel
	if err := c.Bind(&model); err != nil {
		return err
	}

	var user models.User
	if err := models.DB.First(&user, "username = ?", username).Error; err != nil {
		return err
	}
	if model.Password != user.Password {
		return echo.NewHTTPError(http.StatusBadRequest, "current password is incorrect")
	}
	if model.NewPassword == "" || model.NewPassword != model.ConfirmPassword {
		return echo.NewHTTPError(http.StatusBadRequest, "two new passwords do not match")
	}
	if err := models.DB.Model(&user).Update("password", model.NewPassword).Error; err != nil {
		return nil
	}
	return c.JSON(http.StatusAccepted, echo.Map{"id": user.ID})
}
