package logs

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/text"
)

func LogError(err string) {
	panic(text.FgRed.Sprintf("[ERROR] - %s", err))
}
func LogSuccess(who string, data string) {
	fmt.Println(text.BgGreen.Sprintf("[%s]", who) + " " + data)
}
func LogWarning(who string, data string) {
	fmt.Println(text.BgYellow.Sprintf("[%s]", who) + " " + data)
}

func CheckError(err error) bool {
	if err != nil {
		LogError(err.Error())
		return true
	}
	return false
}
