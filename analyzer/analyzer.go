/*
Package analyzer reads users messages and try to fine answer in answers.json File
if message contains some keyword form file
*/
package analyzer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Message struct {
	Text   string `json:"text"`
	Answer string `json:"answer"`
}

var msgs []Message

// Init read json file with messages and answers

func init() {
	raw, err := ioutil.ReadFile("./analyzer/answers.json")
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	err = json.Unmarshal(raw, &msgs)
	if err != nil {
		fmt.Printf("Json error: %v\n", err)
		os.Exit(1)
	}
}

// GetAnswer check messages in file and make answers if have found it
func GetAnswer(text string) string {
	for _, msg := range msgs {
		if strings.Contains(text, msg.Text) {
			return msg.Answer
		}
	}
	return ""
}
