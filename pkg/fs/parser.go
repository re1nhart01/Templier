package fs

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"templier.com/pkg/args"
	"templier.com/pkg/logger"
	"templier.com/pkg/utils"
)

type IArgs interface {
	PrintArgs()
}

var wg sync.WaitGroup

func GenerateFilesFromTemplate(component map[string]any, args args.Arguments) (err error) {
	files := component["files"].([]interface{})
	if files == nil {
		logger.FatalError(utils.ConstantsError["InvalidComponentConf"])
	}
	filesLen := len(files)
	loggerChan := make(chan string, 1)

	parsedCount := parseCount(args.WithCount, filesLen)
	countType := reflect.TypeOf(parsedCount).String()
	if countType == "int" {
		if parsedCount.(int) < 0 || parsedCount.(int) > filesLen {
			logger.FatalError(fmt.Sprintf(utils.ConstantsError["InvalidCountOutOfRange"], 0, filesLen-1))
		}
		selectedFile := files[parsedCount.(int)].(map[string]any)
		generateFileFromTemplate(loggerChan, args, selectedFile, parsedCount.(int), false)
		return
	}
	castedFromTo := parsedCount.([]int)
	from := castedFromTo[0]
	to := castedFromTo[1]

	if from > to || from < 0 || to > filesLen {
		logger.FatalError(fmt.Sprintf(utils.ConstantsError["InvalidCountOutOfRange"], 0, filesLen-1))
	}

	wg.Add(to - from)

	for i := from; i < to; i++ {

		go generateFileFromTemplate(loggerChan, args, files[i].(map[string]any), i, true)

		select {
		case msg := <-loggerChan:
			fmt.Println(msg)
		}
	}

	wg.Wait()

	return nil
}

func generateFileFromTemplate(loggerChan chan string, args args.Arguments, file map[string]any, index int, wgEnabled bool) {
	if wgEnabled {
		defer wg.Done()
	}

	if !utils.CheckIsAllKeysExists(file, "name", "content") {
		logger.FatalError(utils.ConstantsError["InvalidComponentFile"])
	}
	if index >= len(args.WithNames) {
		logger.FatalError(fmt.Sprintf(utils.ConstantsError["InvalidNotEnoughNames"], index, len(args.WithNames)+1))
	}
	lastName := args.WithNames[index]
	fileLabels := args.WithLabel
	firstName := file["name"].(string)
	content := file["content"].(map[string]any)
	if !utils.CheckIsAllKeysExists(content, "data") {
		logger.FatalError(utils.ConstantsError["InvalidContentSection"])
	}

	// find all template in file
	replaceRegexp := regexp.MustCompile(utils.Regexp["TemplateRegexp"])
	fileData := content["data"].(string)

	// check possible labels and if flag array less than possible, throw fatal error
	foundedPossibleLabels := replaceRegexp.FindAllString(fileData, -1)
	if len(foundedPossibleLabels) > len(fileLabels) {
		logger.FatalError(
			utils.ConstantsError["InvalidNotEnoughLabels"],
			fmt.Sprintf(utils.ConstantsError["InvalidNotEnoughLabelsSecond"],
				len(foundedPossibleLabels), len(fileLabels)))
	}

	// loop through labels flag, and replace all content if file
	for i := 0; i < len(fileLabels); i++ {
		found := replaceRegexp.FindString(fileData)
		if found != "" {
			fileData = strings.Replace(fileData, found, fileLabels[i], 1)
		}
	}

	fNameRegexp := regexp.MustCompile(utils.Regexp["FileNameTemplateRegexp"])
	correctFileName := firstName

	isFoundFileNameLabel := fNameRegexp.FindString(firstName)
	if isFoundFileNameLabel != "" {
		correctFileName = strings.Replace(firstName, isFoundFileNameLabel, lastName, 1)
	}

	pathWithoutWd := fmt.Sprintf("%s%s", args.WithPathName, OS_SLASH)
	// make filesystem operations to generate file structure
	_, err := WriteFileSync(pathWithoutWd, fileData, correctFileName)
	if err != nil {
		logger.FatalError(utils.ConstantsError["InvalidFSWriting"])
	}
	//write data logger into log channel and done wait group
	loggerChan <- fmt.Sprintf(utils.ConstantsText["LOGWritingFile"], correctFileName, fileData)
}

func parseCount(countString string, maxCount int) any {
	if countString == "" {
		return []int{0, maxCount}
	}
	c, err := strconv.Atoi(countString)
	if err != nil {

		fromTo := strings.Split(countString, ":")
		if len(fromTo) != 2 {
			logger.FatalError(utils.ConstantsError["InvalidCountError"])
		}

		castedFrom, _ := strconv.Atoi(fromTo[0])
		castedTo, _ := strconv.Atoi(fromTo[1])

		castedFromTo := []int{castedFrom, castedTo}

		return castedFromTo
	}
	return c
}
