package routes

import (
	"github.com/ahmed-deftoner/url-shortener/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func ResolveURL(c *fiber.Ctx) error {
	url := c.Params("url")
	r := database.CreateClient(0)
	defer r.Close()
	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "could not find short"})
	} else if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "could not connect to db"})
	}

	rdr := database.CreateClient(1)
	defer rdr.Close()

	_ = rdr.Incr(database.Ctx, "counter")
	return c.Redirect(value, 301)
}
