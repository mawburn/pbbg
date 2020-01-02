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

type PlayerSector struct {
	SectorId string   `json:"sectorId"`
	Players  []string `json:"players"`
}

func playerSector(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId").(string)

	sector, err := getCurrentSector(userId)

	if err != nil {
		Err500(w, []string{err.Error()})
		return
	}

	jSector, err := json.Marshal(sector)

	if err != nil {
		Err500(w, []string{"Error marshalling sector"})
		return
	}

	w.Write(jSector)
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

func getCurrentSector(userId string) (PlayerSector, error) {
	playerVal, err := dbConns.Redis.Get("player-" + userId).Result()

	if err != nil {
		return PlayerSector{}, fmt.Errorf("Get Current Sector - Error retrieving player")
	}

	var player Player

	err = json.Unmarshal([]byte(playerVal), &player)

	if err != nil {
		return PlayerSector{}, fmt.Errorf("Get Current Sector - Error unmarshalling player")
	}

	sectorVal, err := dbConns.Redis.Get(player.CurSectorId).Result()

	if err != nil {
		return PlayerSector{}, fmt.Errorf("Get Current Sector - Error retrieving sector")
	}

	var sector Sector

	err = json.Unmarshal([]byte(sectorVal), &sector)

	if err != nil {
		return PlayerSector{}, fmt.Errorf("Get Current Sector - Error unmarshalling sector")
	}

	return PlayerSector{
		SectorId: player.CurSectorId,
		Players:  sector.Players,
	}, nil
}
