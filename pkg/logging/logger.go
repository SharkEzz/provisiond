package logging

import (
	"fmt"
	"time"
)

const (
	INFO int = iota
	ERROR
	WARN
	SUCCESS

	green  string = "\033[32m"
	red    string = "\033[0;31m"
	yellow string = "\033[0;33m"
	blue   string = "\033[0;34m"
	reset  string = "\033[0;39m"
)

func generateString(data string) string {
	dt := time.Now()

	return fmt.Sprintf("[%s] %s", dt.Format("2006-01-02 15:04:05"), data)
}

func getColor(logType int) string {
	switch logType {
	case INFO:
		return blue
	case ERROR:
		return red
	case WARN:
		return yellow
	case SUCCESS:
		return green
	}

	return reset
}

func Log(text string) string {
	return generateString(text)
}

func LogOut(text string, logType int) {
	fmt.Print(getColor(logType))
	fmt.Print(Log(text))
	fmt.Print(reset + "\n")
}
