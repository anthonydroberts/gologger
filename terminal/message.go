package terminal

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func Msg(msgType string, message string) {
	timeString := time.Now().Format("15:04:05")

	switch msgType {
	case "success":
		col := color.New(color.FgCyan)
		col.Printf("\n%s Gologger > %s\n", timeString, message)
	case "fail":
		col := color.New(color.FgRed)
		col.Printf("\n%s Gologger > %s\n", timeString, message)
	case "print":
		col := color.New(color.FgHiWhite, color.Bold)
		col.Printf("\n%s Gologger > %s\n", timeString, message)
	default:
		fmt.Printf("\n%s Gologger > %s\n", timeString, message)
	}
}
