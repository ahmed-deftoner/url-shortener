package routes

import (
	"time"

	"github.com/ahmed-deftoner/url-shortener/helpers"
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type respomse struct {
	URL           string        `json:"url"`
	CustomShort   string        `json:"short"`
	Expiry        time.Duration `json:"expiry"`
	RateRemaining int           `json:"rate_limit"`
	RateLimitRest time.Duration `json:"rate_limit_remaining"`
}

func ShortenURL(c *fiber.Ctx) error {
	body := new(request)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse json"})
	}

	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "wrong url"})
	}

	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "fuck u!"})
	}
	body.URL = helpers.EnforceHTTP(body.URL)
}
