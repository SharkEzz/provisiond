package api

import (
	"fmt"
	"net/http"

	handlers_v1 "github.com/SharkEzz/provisiond/pkg/api/handlers/v1"
	"github.com/SharkEzz/provisiond/pkg/logging"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
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

	ch := make(chan string)

	v1.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	v1.Use(func(c *fiber.Ctx) error {
		fmt.Println(logging.Log(fmt.Sprintf("Handling %s request from %s on path '%s'", c.Method(), c.IP(), c.Path())))

		c.Locals("channel", ch)

		if string(c.Request().URI().LastPathSegment()) == "healthcheck" ||
			string(c.Request().URI().LastPathSegment()) == "ws" {
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

	v1.Get("/healthcheck", handlers_v1.HandleGetHealthcheck)
	v1.Post("/deploy", handlers_v1.HandlePostDeploy)
	v1.Get("/ws", websocket.New(handlers_v1.HandleGetWebsocket))

	return &API{
		host:     host,
		port:     port,
		password: password,
		fiber:    app,
	}
}

func (a *API) Start() {
	fmt.Println(logging.Log(fmt.Sprintf("Starting API on %s:%d", a.host, a.port)))

	a.fiber.Listen(fmt.Sprintf("%s:%d", a.host, a.port))
}
