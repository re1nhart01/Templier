package main

import (
	"gopkg.in/yaml.v3"
	"templier.com/pkg/args"
	"templier.com/pkg/fs"
	"templier.com/pkg/logger"
	"templier.com/pkg/utils"
)

type Templier struct {
	Author  string
	Version string
	Name    string
}

func NewApplication() Templier {
	return Templier{
		Author:  utils.SystemInfo["Author"],
		Name:    utils.SystemInfo["Name"],
		Version: utils.SystemInfo["Version"],
	}
}

func (app *Templier) RunTemplier() (err error) {
	args := args.GetArguments()
	if args.Component == "" && args.WithLogger {
		logger.LogHelper()
		return
	}
	workingFile := fs.ReadFile(args.WithFileName)
	if len(workingFile) <= 0 {
		logger.FatalError(utils.ConstantsError["FatalError"])
	}
	mappableYaml := make(map[string]any)
	err = yaml.Unmarshal(workingFile, &mappableYaml)
	if err != nil {
		logger.FatalError(utils.ConstantsError["InvalidYmlParse"])
	}

	return nil
}
