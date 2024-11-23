package minifier

import (
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Next func(c *fiber.Ctx) bool

	SuppressWarnings, MinifyHTML, MinifyCSS, MinifyJS, MinifyXML, MinifyJSON bool
}

var ConfigDefault = Config{Next: nil}

func configDefault(config ...Config) Config {
	if len(config) < 1 {
		return ConfigDefault
	}

	return config[0]
}
