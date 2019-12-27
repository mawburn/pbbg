package main

import (
	"encoding/json"
	"net/http"
)

type PlayerMove struct {
	Direction string `json:"direction"`
}

type Player struct {
	Id string
}

func playerMove(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var m PlayerMove
	err := decoder.Decode(&m)

	if err != nil {
		panic(err)
	}

	w.Write(json.RawMessage(`{"precomputed": true}`))
}
