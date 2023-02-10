package utils

var Constants = map[string]any{}

var ConstantsError = map[string]string{
	"ArgumentsLengthError": "Not enought arguments. --help to look out",
}

var Regexp = map[string]string{
	"StringRegexp": "`([a-zA-Z0-9]+)`gm",
}
