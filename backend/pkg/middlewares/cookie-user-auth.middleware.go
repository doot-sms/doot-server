package middlewares

import (
	"github.com/doot-sms/doot-server/pkg/token"
	"github.com/doot-sms/doot-server/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func NewCookieUserAuthMiddleware(tokenMaker token.TokenMaker, config utils.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {

		accessToken := c.Cookies(config.AccessTokenCookieName)

		payload, err := tokenMaker.VerifyToken(accessToken)

		if err != nil {
			c.ClearCookie(config.AccessTokenCookieName)
			c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "invalid token",
			})
		}

		// accessible to next handlers
		c.Locals("user", payload)

		return c.Next()
	}
}
