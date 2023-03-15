package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JWTCustomClaims struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	jwt.StandardClaims
}

const SigningKey = "repeaterjwtsecret!"

func skip(c echo.Context) bool {
	return c.Request().Method == "OPTIONS" || c.Path() == "/api/user/login" || c.Path() == "/api/prometheus/alert/receiver"
}

func JWT() echo.MiddlewareFunc {

	config := middleware.JWTConfig{
		Skipper:    skip,
		Claims:     &JWTCustomClaims{},
		SigningKey: []byte(SigningKey),
	}
	return middleware.JWTWithConfig(config)
}
