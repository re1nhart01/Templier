package utils

var Constants = map[string]any{}

var ConstantsError = map[string]string{
	"ArgumentsLengthError":      "Not enought arguments. --help to look out",
	"InvalidFirstArgumentError": "First argument should to be a component name from yaml file. Or component is invalid. Allow only string(in any case) or(and) numbers",
	"InvalidArrayArgError":      "Invalid value. Arguments -l or -n should to be as array without spaces. Example: -f [first,second,third]",
	"InvalidNameArgError":       "Invalid flag for file name. Example: -f TemplierFileConfig.yaml",
	"InvalidArgument":           "Invalid Argument. --help to look out",
}

var Regexp = map[string]string{
	"StringRegexp":      "^[a-zA-Z0-9]+$",
	"FileNameRegexp":    "^([a-zA-Z0-9].+)\\.(yaml|yml)$",
	"StringArrayRegexp": "^\\[(.*?),(.*?)\\]$",
}

var ReservedArguments = map[string]string{
	"Labels":   "-l",
	"Names":    "-n",
	"Logger":   "--log",
	"FileName": "-f",
	"Help":     "--help",
}
