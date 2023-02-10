package args

import (
	"os"

	"templier.com/pkg/logger"
	"templier.com/pkg/utils"
)

// --help
// -l
// templier compName -l [aboba, aababa] -n [zxc, zxc zxc]

type possibleArguments interface {
	~string | ~[]string
}

type Arguments struct {
}

func (argsv *Arguments) getArgumentsDictionary() map[string]any {
	args := os.Args
	if len(args) < 2 {
		logger.LogError(utils.ConstantsError["ArgumentsLengthError"])
		return map[string]any{}
	}
	argumentsDict := make(map[string]any)
	for index, arg := range args[1:] {

	}
	return argumentsDict
}

func GetArguments() []string {
	argsv := Arguments{}
	argsv.getArguments()
	return []string{}
}
