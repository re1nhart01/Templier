package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
	"templier.com/pkg/args"
)

func main() {
	test := map[string]any{}
	file, err := os.ReadFile("test/test_cfg/templier.yaml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(file))
	yaml.Unmarshal(file, &test)
	fmt.Println(test["zxc2254"])
	args.GetArguments()
}
