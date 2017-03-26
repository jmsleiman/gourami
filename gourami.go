package main

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
)

// since JWT was the target use-case, and JWT are three segments split by `.`
// we can just run the function on each segment and print out the results.
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		split := strings.SplitN(scanner.Text(), ".", 3)

		if len(split) != 3 {
			continue
		}

		formatAndPrint(split[0])
		formatAndPrint(split[1])
	}
}

// formatAndPrint will "pretty-print" the contents of a base-64 json segment.
// it takes in base64-encoded data, with or without padding, and will parse
// the JSON contents to then print them.
func formatAndPrint(input string) {
	context := log.WithFields(log.Fields{
		"funcName":   "formatAndPrint",
		"sourceFile": "main.go",
		"input":      input,
	})

	if l := len(input) % 4; l > 0 {
		input += strings.Repeat("=", 4-l)
	}

	de64, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		context.WithFields(log.Fields{
			"len": len(input),
			"err": err,
		}).Warn("input was not base64")
		return
	}

	var dump map[string]interface{}

	if err := json.Unmarshal(de64, &dump); err != nil {
		context.WithField("err", err).Warn("input was not json")
		return
	}

	if formatted, err := json.MarshalIndent(dump, "", "\t"); err != nil {
		context.WithField("err", err).Warn("input could not be formatted")
	} else {
		fmt.Println(string(formatted))
	}

}
