package logging

import (
	"fmt"
	"time"
)

func generateString(data string) string {
	dt := time.Now()

	return fmt.Sprintf("[%s] %s", dt.Format("2006-01-02 15:04:05"), data)
}

func Log(text string) string {
	return generateString(text)
	// switch level {
	// case "info":
	// 	return fmt.Sprintf("\x1B[34m%s\x1b[0m", generateString(text))
	// case "warn":
	// 	return fmt.Sprintf("\x1b[33m%s\x1b[0m", generateString(text))
	// case "error":
	// 	return fmt.Sprintf("\x1b[31m%s\x1b[0m", generateString(text))
	// case "debug":
	// 	return fmt.Sprintf("\x1b[34m%s\x1b[0m", generateString(text))
	// case "success":
	// 	return fmt.Sprintf("\x1b[32m%s\x1b[0m", generateString(text))
	// default:
	// 	return text
	// }
}

func LogOut(text string) {
	fmt.Println(Log(text))
}
