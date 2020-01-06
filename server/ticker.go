package main

import (
	//	"fmt"
	"math/rand"
	"time"
)

func runTicker() {
	ticker := time.NewTicker(1 * time.Second)
	for _ = range ticker.C {
		runActions()
	}
}

func runActions() {
	for _, sys := range systemActions {
		go func() {
			for _, action := range actions {
				// shuffle player actions
				rand.Seed(time.Now().UnixNano())
				rand.Shuffle(len(sys[action]), func(i, j int) { sys[action][i], sys[action][j] = sys[action][j], sys[action][i] })
			}
		}()
	}
}
