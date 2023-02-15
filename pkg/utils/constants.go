package utils

var Constants = map[string]any{
	"DefaultFileName": "Templier.yaml",
}

var ConstantsError = map[string]string{
	"FatalError":                "Unknown fatal error.",
	"ArgumentsLengthError":      "Not enought arguments. --help to look out",
	"InvalidFirstArgumentError": "First argument should to be a component name from yaml file. Or component is invalid. Allow only string(in any case) or(and) numbers",
	"InvalidArrayArgError":      "Invalid value. Arguments -l or -n should to be as array without spaces. Example: -f [first,second,third]",
	"InvalidNameArgError":       "Invalid flag for file name. Example: -f TemplierFileConfig.yaml",
	"InvalidArgument":           "Invalid Argument. --help to look out",
	"InvalidFS":                 "Invalid filesystem operation.Might be caught if file not exists. Current path:",
	"InvalidYmlParse":           "Invalid yaml file. Check is valid",
}

var Regexp = map[string]string{
	"StringRegexp":      "^[a-zA-Z0-9]+$",
	"FileNameRegexp":    "^([a-zA-Z0-9].+)\\.(yaml|yml)$",
	"StringArrayRegexp": "^\\[(.*?),(.*?)\\]$",
}

const Helper = `
	Welcome to Templier.
	The purpose of this application to generate template files.
	 For example you have angular components.
	 And you dont want to always create files,
	 you can just define config yaml file and call this
	 application with component name.
     Which are a default routine in developers live.
	 Usage:
	 Templier.yaml file::
	 zxc2254:
	files:
	- name: "@aboba@.component.tsx"
		content:
		labels:
			componentName: Zalupe
			componentPropType: zalupa
		data: |
			import React from 'react';
			
			type {{componentPropType}}componentType = {};
			
			const {{componentName}} = () => {
				return (
				<div>
					zxczxc daun;
				</div>
				)
			}
	- name: "@aboba@.component.tsx"
		content:
		labels: 
			- componentPropType
			- componentName
		data: |
			import React from 'react';
			
			type {{componentPropType}}componentType = {};
			
			const {{componentName}} = () => {
				return (
				<div>
					zxczxc daun;
				</div>
				)
			}     
			
	<------------------------------------------->
	Flags:
	-l - flag to set labels which define in {{ }} brackets
	-n - flag to set file names for files inside @ @ brackets
	--log - flag to set logger on, which will log every step
	--help - current flag
	-f - filename of file with components	
`

var ReservedArguments = map[string]string{
	"Labels":   "-l",
	"Names":    "-n",
	"Logger":   "--log",
	"FileName": "-f",
	"Help":     "--help",
}

var SystemInfo = map[string]string{
	"Version": "0.0.1",
	"Author":  "re1nhart",
	"Name":    "Templier",
}
