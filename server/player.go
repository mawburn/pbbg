package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-redis/redis/v7"
)

type PlayerMove struct {
	Direction string `json:"direction"`
}

type Player struct {
	CurSystemId string `json:"system"`
	CurSectorId string `json:"sectorId"`
}

func playerMove(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var m PlayerMove
	err := decoder.Decode(&m)

	if err != nil {
		Err500(w, []string{"Unable to parse request"})
		return
	}

	userId := r.Context().Value("userId").(string)

	playerVal, err := dbConns.Redis.Get("player-" + userId).Result()

	if err != nil {
		Err500(w, []string{"Unable to retrieve player"})
		return
	}

	var playerInfo Player

	err = json.Unmarshal([]byte(playerVal), &playerInfo)

	if err != nil {
		Err500(w, []string{"Unable to retrieve player"})
		return
	}

	switch m.Direction {
	case "up":
	case "down":
	case "left":
	case "right":
		fmt.Println(playerInfo.CurSectorId)
		break
	default:
		Err500(w, []string{"Invalid Direction"})
		return
	}

	rerr := dbConns.Redis.Set("key", m.Direction, 0).Err()
	if rerr != nil {
		panic(rerr)
	}

	w.Write(json.RawMessage(`{"precomputed": true}`))
}

func getToken(r *http.Request, c *redis.Client) string {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	if false {
		return "X"
	}

	b := make([]byte, 16)
	rand.Read(b)
	fmt.Sprintf("%x", b)

	return string(b)
}
