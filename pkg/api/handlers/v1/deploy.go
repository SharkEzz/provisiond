package v1

import (
	"net/http"
	"time"

	"github.com/SharkEzz/provisiond/pkg/executor"
	"github.com/SharkEzz/provisiond/pkg/loader"
	"github.com/gofiber/fiber/v2"
)

func HandlePostDeploy(c *fiber.Ctx) error {
	config := string(c.Body())
	if config == "" {
		c.Status(http.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": "request body cannot be empty",
		})
	}
	ch := c.Locals("channel").(chan string)

	deployment, err := loader.GetLoader(string(config)).Load()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(err)
	}

	go func() {
		time.Sleep(time.Second)
		executor.NewExecutor(deployment, ch).ExecuteJobs()
	}()

	return c.JSON(fiber.Map{
		"started": true,
	})
}
