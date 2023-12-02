package main

import (
	"fmt"
	"os"
	"strings"
)

var standardPaths = map[string]any{
	"/usr/local/bin": nil,
	"/usr/bin":       nil,
	"/usr/sbin":      nil,
	"/bin":           nil,
	"/sbin":          nil,
}

func main() {
	if err := mainE(); err != nil {
		panic(err)
	}
}

func mainE() error {
	pathArg := os.Args[1]
	orderedPathArgs := os.Args[2:]

	if pathArg == "" {
		return fmt.Errorf("first argument must be the PATH string")
	}
	newParts := []string{}
	parts := strings.Split(pathArg, ":")

	for _, part := range parts {
		if _, ok := standardPaths[part]; ok {
			continue
		}
		newParts = append(newParts, part)
	}

	return nil
}
