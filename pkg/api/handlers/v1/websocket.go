package v1

import (
	"github.com/gofiber/websocket/v2"
)

func HandleGetWebsocket(c *websocket.Conn) {
	ch := c.Locals("channel").(chan string)

	hasError := make(chan bool)
	running := true

	go func() {
		for {
			if _, _, err := c.ReadMessage(); err != nil {
				hasError <- true
			}
		}
	}()

	for running {
		select {
		case msg := <-ch:
			c.WriteMessage(1, []byte(msg))
		case <-hasError:
			running = false
		}
	}
}
