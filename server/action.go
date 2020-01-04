package main

import (
	"encoding/json"
//	"fmt"
	"io/ioutil"
	"os"
)

type Action struct {
	Class  string `json:"class"`
	Weight uint8  `json:"weight"`
}

type PlayerAction struct {
	UserId     string
	ActionType string
	Command    string
	// we'll probably need more stuff here or in command
}

var actions map[string]*Action

// System ID mapped to a map that contains an array of player actions
// This is done because we want to process actions in hard precedence based on type
var systemActions map[string]map[string][]PlayerAction

func getAction(id string) *Action {
	return actions[id]
}

func addAction(sysId string, pAction PlayerAction) {
	if _, ok := systemActions[sysId]; !ok {
		systemActions[sysId][pAction.ActionType] =  []PlayerAction{pAction}
		return
	}

	if _, ok := systemActions[sysId][pAction.ActionType]; !ok {
		systemActions[sysId][pAction.ActionType] =  []PlayerAction{pAction}
		return
	}


	var userLastAction int

	userLastAction = -1

	for i, v := range systemActions[sysId][pAction.ActionType] {
		if v.UserId == pAction.UserId {
			userLastAction = i
		}
	}

	if userLastAction != -1 {
		systemActions[sysId][pAction.ActionType][userLastAction] = systemActions[sysId][pAction.ActionType][len(systemActions[sysId][pAction.ActionType])-1]
		systemActions[sysId][pAction.ActionType] = systemActions[sysId][pAction.ActionType][:len(systemActions[sysId][pAction.ActionType])-1]
	}

	systemActions[sysId][pAction.ActionType] = append(systemActions[sysId][pAction.ActionType], pAction)
}

func initActions() {
	actionFile, err := os.Open("./static/actions.json")

	if err != nil {
		panic(err)
	}

	defer actionFile.Close()

	byteValue, err := ioutil.ReadAll(actionFile)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(byteValue, &actions)

	if err != nil {
		panic(err)
	}
}
