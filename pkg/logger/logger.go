package logger

import (
	"fmt"
	"strings"
	"syscall"

	"templier.com/pkg/utils"
)

const labelMessage = "@TEMPLIER Error: "

func FatalError(errMsg string, args ...string) {
	fmt.Println(labelMessage+errMsg, "|", strings.Join(args, " and "))
	syscall.Exit(0)
}

func LogHelper() {
	fmt.Print(utils.Helper)
	syscall.Exit(0)
}

func LogSteps(content ...string) {
	for i := 0; i < len(content); i++ {
		fmt.Println(fmt.Sprintf("Step %d: %s", i, content[i]))
	}
}
