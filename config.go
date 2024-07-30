package minifier

import (
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Next func(c *fiber.Ctx) bool

	MinifyHTML, MinifyCSS, MinifyJS, MinifyXML, MinifyJSON bool
}

var ConfigDefault = Config{Next: nil, MinifyHTML: true}

func configDefault(config ...Config) Config {
	if len(config) < 1 {
		return ConfigDefault
	}

	return config[0]
}
