package logger

import "fmt"

const labelMessage = "@TEMPLIER Error: "

func LogError(errMsg string) {
	fmt.Println(labelMessage + errMsg)
}
