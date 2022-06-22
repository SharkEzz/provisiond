package api

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type API struct {
	host     string
	port     uint16
	password string
	fiber    *fiber.App
}

func NewAPI(host string, port uint16, password string) *API {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		AppName:               "provisiond",
	})

	v1 := app.Group("/v1")

	v1.Use(func(c *fiber.Ctx) error {
		logrus.Infof("Handling %s request from %s on path '%s'", c.Method(), c.IP(), c.Path())

		if string(c.Request().URI().LastPathSegment()) == "healthcheck" {
			return c.Next()
		}

		requestPassword := c.Get("password", "")

		if requestPassword != password {
			c.Status(http.StatusUnauthorized)
			return c.JSON(fiber.Map{
				"error": "unauthorized",
			})
		}

		return c.Next()
	})

	v1.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
		})
	})
	v1.Post("/deploy", func(c *fiber.Ctx) error {
		// TODO
		return c.JSON(fiber.Map{
			"wip": true,
		})
	})

	return &API{
		host:     host,
		port:     port,
		password: password,
		fiber:    app,
	}
}

func (a *API) Start() {
	logrus.Infof("Starting API server on '%s:%d'", a.host, a.port)

	a.fiber.Listen(fmt.Sprintf("%s:%d", a.host, a.port))
}
