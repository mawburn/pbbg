package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Action struct {
	Class  string `json:"class"`
	Weight uint8  `json:"weight"`
}

var actions map[string]*Action

func getAction(id string) *Action {
	return actions[id]
}

func initActions() {
	actionFile, err := os.Open("./static/actions.json")

	if err != nil {
		fmt.Println(err)
	}

	defer actionFile.Close()

	byteValue, _ := ioutil.ReadAll(actionFile)

	umErr := json.Unmarshal(byteValue, &actions)

	if umErr != nil {
		fmt.Println(umErr)
	}
}
