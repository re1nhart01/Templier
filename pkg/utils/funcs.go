package utils

import (
	"fmt"
	"strings"
)

func CheckIsAllKeysExists(mappable map[string]any, keys ...string) bool {
	isExists := true
	for _, value := range keys {
		if mappable[value] == nil {
			isExists = false
		}
	}
	return isExists
}

func AssertEq(cond bool, msg ...string) {
	if !cond {
		panic(fmt.Sprintf("AssertionError: Condition is not equal. Trace: " + strings.Join(msg, " and \n")))
	}
}
