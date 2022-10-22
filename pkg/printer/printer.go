package printer

import (
	"fmt"
	"go-echo/pkg/color"
	"time"
)

func Infof(format string, msg any) {
	printf(color.Info(format), "INFO ", msg)
}

func Info(msg any) {
	print("INFO ", color.Info(msg))
}

func Warnf(format string, msg any) {
	printf(color.Warn(format), "WARN ", msg)
}

func Warn(msg any) {
	print("WARN ", color.Warn(msg))
}

func Errorf(format string, msg ...any) {
	printf(color.Fatal(format), "ERROR", msg)
}

func Error(msg any) {
	print("ERROR", color.Fatal(msg))
}

func printf(format string, level string, msg any) {
	currentTime := time.Now()
	date := color.Yellow("%d-%d-%d %d:%d:%d:%d")
	fullFormat := date + " - %s -    " + format + "\n"

	fmt.Printf(
		fullFormat,
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		currentTime.Hour(),
		currentTime.Minute(),
		currentTime.Second(),
		currentTime.Nanosecond()/100000,
		color.Purple(level),
		msg)
}

func print(level string, msg any) {
	currentTime := time.Now()
	date := color.Yellow("%d-%d-%d %d:%d:%d:%d")
	fullFormat := date + " - %s -    %s\n"

	fmt.Printf(
		fullFormat,
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		currentTime.Hour(),
		currentTime.Minute(),
		currentTime.Second(),
		currentTime.Nanosecond()/100000,
		color.Purple(level),
		msg)
}
