package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

// PlayerAction - holds information about actions players have taken
type PlayerAction struct {
	UserID     string
	ActionType string
	Command    string
	// we'll probably need more stuff here or in command
}

var actions []string

// System ID mapped to a map that contains an array of player actions
// This is done because we want to process actions in hard precedence based on type
var systemActions map[string]map[string][]PlayerAction

var lock = sync.RWMutex{}

func addAction(sysID string, pAction PlayerAction) {
	lock.Lock()
	defer lock.Unlock()

	// If the system doesn't exist, then go ahead an add the player's action
	if _, ok := systemActions[sysID]; !ok {
		systemActions[sysID][pAction.ActionType] = []PlayerAction{pAction}
		return
	}

	// Check if the player already has actions queued & remove them
	var userLastAction int

	userLastAction = -1

	for i, v := range systemActions[sysID][pAction.ActionType] {
		if v.UserID == pAction.UserID {
			userLastAction = i
		}
	}

	if userLastAction != -1 {
		systemActions[sysID][pAction.ActionType][userLastAction] = systemActions[sysID][pAction.ActionType][len(systemActions[sysID][pAction.ActionType])-1]
		systemActions[sysID][pAction.ActionType] = systemActions[sysID][pAction.ActionType][:len(systemActions[sysID][pAction.ActionType])-1]
	}

	// If the action type doesn't have any actions, just add the player's action
	if _, ok := systemActions[sysID][pAction.ActionType]; !ok {
		systemActions[sysID][pAction.ActionType] = []PlayerAction{pAction}
		return
	}

	systemActions[sysID][pAction.ActionType] = append(systemActions[sysID][pAction.ActionType], pAction)
	return
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

	fmt.Println("Loaded Actions", string(byteValue))
}
