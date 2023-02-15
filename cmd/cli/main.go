package main

import (
	"templier.com/pkg/logger"
	"templier.com/pkg/utils"
)

func main() {
	app := NewApplication()
	err := app.RunTemplier()
	if err != nil {
		logger.FatalError(utils.ConstantsError["FatalError"])
	}
}
