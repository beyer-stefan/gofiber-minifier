package minifier

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
	"github.com/tdewolff/minify/v2/json"
	"github.com/tdewolff/minify/v2/xml"
	"regexp"
)

func New(config ...Config) fiber.Handler {
	var (
		cfg Config = configDefault(config...)
	)

	return func(c *fiber.Ctx) error {
		var (
			err      error
			origBody []byte
			m        *minify.M
			mimetype string
		)

		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		if err = c.Next(); err != nil {
			return err
		}

		if c.Response().StatusCode() != fiber.StatusOK {
			return nil
		}

		// make a copy of the original body
		origBody = c.Response().Body()
		// reset the body in the response to allow us to start over and
		// write a minified version of the original body
		c.Response().ResetBody()

		m = minify.New()

		if cfg.MinifyHTML {
			m.Add("text/html", &html.Minifier{
				// avoid breaking things, e.g. Shoelace.style web components or LinkedIn sharing
				KeepEndTags:      true,
				KeepDocumentTags: true,
			})
		}

		if cfg.MinifyCSS {
			m.Add("text/css", &css.Minifier{})
		}

		if cfg.MinifyJS {
			m.AddRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), &js.Minifier{})
		}

		if cfg.MinifyXML {
			m.AddRegexp(regexp.MustCompile("xml$"), &xml.Minifier{})
		}

		if cfg.MinifyJSON {
			m.AddRegexp(regexp.MustCompile("json$"), &json.Minifier{})
		}

		mimetype = string(c.Response().Header.Peek("Content-Type")[:])
		if err = m.Minify(mimetype, c.Response().BodyWriter(), bytes.NewReader(origBody)); err != nil {
			log.Warnf("%s '%s'", err.Error(), mimetype)
			// Minifying does not work (aka: returned an error),
			// so we fail in a gentle way by writing the original (un-minified) body
			c.Response().BodyWriter().Write(origBody)
		}

		return nil
	}
}
