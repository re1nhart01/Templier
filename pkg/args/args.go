package args

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"

	"templier.com/pkg/logger"
	"templier.com/pkg/utils"
)

// first argument must be component from yml file
// --help - helper
// --log - logger for all operations
// -f - filename
// -l - labels
// -n - for components name
// -p - pathname where we need to create file
// templier compName -l [aboba, aababa] -n [zxc, zxc zxc]

type Arguments struct {
	Component    string
	System       string
	WithLabel    []string
	WithNames    []string
	WithFileName string
	WithPathName string
	WithHelp     bool
	WithLogger   bool
}

func GetArguments() Arguments {
	osArgs := os.Args
	if len(osArgs) < 2 {
		logger.FatalError(utils.ConstantsError["ArgumentsLengthError"])
	}
	usableArguments := osArgs[1:]
	argsv := Arguments{
		System: osArgs[0],
	}
	argsv.getArgumentsDictionary(usableArguments)
	return argsv
}

func (argsv Arguments) PrintArgs() {
	val := reflect.ValueOf(argsv)

	values := make([]interface{}, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		values[i] = val.Field(i).Interface()
	}

	fmt.Println(values)
}

func (argsv *Arguments) getArgumentsDictionary(args []string) {

	usableArgumentsLength := len(args)
	if args[0] != utils.ReservedArguments["Help"] {
		re, _ := regexp.Compile(utils.Regexp["StringRegexp"])

		if isValid := re.Match([]byte(args[0])); !isValid {
			logger.FatalError(utils.ConstantsError["InvalidFirstArgumentError"])
		}
		argsv.Component = args[0]
	} else {
		argsv.WithLogger = true
	}

	for i := 1; i < usableArgumentsLength; i++ {
		switch args[i] {
		case utils.ReservedArguments["Help"]: // --help
			argsv.WithHelp = true
			break

		case utils.ReservedArguments["Logger"]: //--log
			argsv.WithLogger = true
			break

		case utils.ReservedArguments["FileName"]: // -f
			if usableArgumentsLength <= i+1 {
				logger.FatalError(utils.ConstantsError["InvalidNameArgError"])
			}
			argsv.WithFileName = checkIsNextArgValid(args[i+1], "FileNameRegexp", "InvalidNameArgError")
			i += 1
			break

		case utils.ReservedArguments["Labels"]: // -l
			if usableArgumentsLength <= i+1 {
				logger.FatalError(utils.ConstantsError["InvalidArrayArgError"])
			}
			list := checkIsNextArgValid(args[i+1], "StringArrayRegexp", "InvalidArrayArgError")
			argsv.WithLabel = strings.Split(list[1:len(list)-1], ",")
			i += 1
			break

		case utils.ReservedArguments["Names"]: // -n
			if usableArgumentsLength <= i+1 {
				logger.FatalError(utils.ConstantsError["InvalidArrayArgError"])
			}
			list := checkIsNextArgValid(args[i+1], "StringArrayRegexp", "InvalidArrayArgError")
			argsv.WithNames = strings.Split(list[1:len(list)-1], ",")
			i += 1
			break
		case utils.ReservedArguments["Path"]:
			if usableArgumentsLength <= i+1 {
				logger.FatalError(utils.ConstantsError["InvalidArrayArgError"])
			}
			path := checkIsNextArgValid(args[i+1], "PathRegexp", "asdsasa")
			argsv.WithPathName = path
			i += 1
			break

		default: // other which not defined
			// logger.FatalError(utils.ConstantsError["InvalidArgument"])
		}
	}
}

func checkIsNextArgValid(next, regexpKey, errorkey string) string {
	re, err := regexp.Compile(utils.Regexp[regexpKey])
	if err != nil {
		panic(err)
	}
	if isValid := re.Match([]byte(next)); !isValid {
		logger.FatalError(utils.ConstantsError[errorkey])
	}
	return next
}
