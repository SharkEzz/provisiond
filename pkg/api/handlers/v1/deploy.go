package v1

import (
	"net/http"

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

	deployment, err := loader.GetLoader(string(config)).Load()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(err)
	}

	err = executor.NewExecutor(deployment).ExecuteJobs()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(err)
	}

	return c.JSON(fiber.Map{
		"done": true,
	})
}
