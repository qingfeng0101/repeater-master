package middleware

import (
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash"
	"net/http"
	"time"

	"github.com/kexirong/repeater/models"
	"github.com/labstack/echo/v4"
)

var ErrInvalidAPIKey = echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired api key")

type Cipher string

const Csha256 Cipher = "sha256"
const Csha1 Cipher = "sha1"

func getBs(n int) []byte {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		//在rand.read 失败函数可用
		nano := time.Now().UnixNano()
		for i := 0; i < n; i++ {
			b[i] = byte(nano >> uint(i&63))
		}
	}
	return b
}

func GenAPIKey(cipher Cipher) string {
	var h hash.Hash
	switch cipher {
	case Csha256:
		h = sha256.New()
	case Csha1:
		fallthrough
	default:
		h = sha1.New()
	}
	h.Write(getBs(16))

	return fmt.Sprintf("%x", h.Sum(nil))

}

func APIKey() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			key := c.Request().Header.Get("API-Key")
			if key == "" {
				return ErrInvalidAPIKey
			}
			var apiKey models.APIKey
			if err := models.DB.First(&apiKey, "key = ?", key).Error; err != nil {
				return &echo.HTTPError{
					Code:     ErrInvalidAPIKey.Code,
					Message:  ErrInvalidAPIKey.Message,
					Internal: err,
				}
			}
			if apiKey.UpdatedAt+apiKey.ExpireDelta < uint(time.Now().Unix()) {
				return ErrInvalidAPIKey
			}
			return next(c)
		}
	}
}
