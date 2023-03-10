package fs

import (
	"fmt"
	"os"
	"strings"

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

func WriteFileSync(path, content string, slices ...string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	fullPath := fmt.Sprintf("%s%s%s", wd, OS_SLASH, path)
	err = os.MkdirAll(fullPath, 0777)
	if err != nil {
		return "", err
	}
	for i := 0; i < len(slices); i++ {
		fullPath += fmt.Sprintf("%s%s", OS_SLASH, slices[i])
	}
	err = os.WriteFile(fullPath, []byte(content), 0777)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(utils.ConstantsText["LOGWritingFile"], strings.Join(slices, ","), content), nil
}
