package middlewares

import (
	"github.com/doot-sms/doot-server/pkg/token"
	"github.com/doot-sms/doot-server/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type CookieAuthMiddleware struct {
	tokenMaker token.TokenMaker
	config     utils.Config
}

func (cam *CookieAuthMiddleware) NewCookieAuthMiddleware(tokenMaker token.TokenMaker) fiber.Handler {
	return func(c *fiber.Ctx) error {

		accessToken := c.Cookies(cam.config.AccessTokenCookieName)

		payload, err := tokenMaker.VerifyToken(accessToken)

		if err != nil {
			c.ClearCookie(cam.config.AccessTokenCookieName)
			c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "invalid token",
			})
		}

		// accessible to next handlers
		c.Locals("user", payload)

		return c.Next()
	}
}
