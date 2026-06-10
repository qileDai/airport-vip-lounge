package middleware

import (
	"airport-vip-lounge/src/repository"
	"time"

	"github.com/gofiber/fiber/v2"
)

func AuditMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()

		log := &repository.AuditLog{
			Action:    c.Method() + " " + c.Path(),
			UserID:    c.Get("X-User-ID", "anonymous"),
			Details:   string(c.Body()),
			IPAddress: c.IP(),
			UserAgent: c.Get("User-Agent"),
		}
		repository.NewAuditLogRepository().Create(log)

		return err
	}
}

func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		duration := time.Since(start)

		println(
			"[", time.Now().Format("2006-01-02 15:04:05"), "]",
			c.Method(),
			c.Path(),
			c.Response().StatusCode(),
			duration.String(),
		)
		return err
	}
}
