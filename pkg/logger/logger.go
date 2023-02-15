package logger

import (
	"fmt"
	"syscall"
)

const labelMessage = "@TEMPLIER Error: "

func FatalError(errMsg string) {
	fmt.Println(labelMessage + errMsg)
	syscall.Exit(0)
}
