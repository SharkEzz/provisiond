package logging

import (
	"fmt"
	"time"
)

// TODO: refactor this

func generateString(data string) string {
	dt := time.Now()

	return fmt.Sprintf("[%s] %s", dt.Format("2006-01-02 15:04:05"), data)
}

func Log(text string) string {
	return generateString(text)
}

func LogOut(text string) {
	fmt.Println(Log(text))
}
