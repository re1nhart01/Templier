package logger

import (
	"fmt"
	"syscall"

	"templier.com/pkg/utils"
)

const labelMessage = "@TEMPLIER Error: "

func FatalError(errMsg string, args ...string) {
	fmt.Println(labelMessage+errMsg, args)
	syscall.Exit(0)
}

func LogHelper() {
	fmt.Print(utils.Helper)
	syscall.Exit(0)
}
