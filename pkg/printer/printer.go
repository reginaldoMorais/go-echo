package printer

import (
	"fmt"
	"go-echo/pkg/color"
	"time"
)

func Infof(format string, msg string) {
	printf(format, "INFO ", color.Info(msg))
}

func Info(msg string) {
	print("INFO ", color.Info(msg))
}

func Warnf(format string, msg string) {
	printf(format, "WARN ", color.Warn(msg))
}

func Warn(msg string) {
	print("WARN ", color.Warn(msg))
}

func Errorf(format string, msg string) {
	printf(format, "ERROR", color.Fatal(msg))
}

func Error(msg string) {
	print("ERROR", color.Fatal(msg))
}

func printf(format string, level string, msg string) {
	currentTime := time.Now()
	fmt.Printf(
		color.Yellow("%d-%d-%d %d:%d:%d:%d - %s -  "+color.Info(format)+"\n"),
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

func print(level string, msg string) {
	currentTime := time.Now()
	fmt.Printf(
		color.Yellow("%d-%d-%d %d:%d:%d:%d - %s -  %s\n"),
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
