package fs

import (
	"fmt"
	"os"

	"templier.com/pkg/logger"
	"templier.com/pkg/utils"
)

func ReadFile(fileName string) []byte {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	currFile := fileName
	if currFile == "" {
		currFile = utils.Constants["DefaultFileName"].(string)
	}
	workingFile := fmt.Sprintf("%s%s%s", wd, OS_SLASH, currFile)
	content, err := os.ReadFile(workingFile)
	if err != nil {
		logger.FatalError(utils.ConstantsError["InvalidFS"], workingFile)
	}
	return content
}
