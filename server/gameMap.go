package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type MapObject struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Max uint32 `json:"max"`
}

type MapCelestial struct {
	Id   string `json:"id"`
	Name string `"json:"name"`
	Type string `json:"type"`
}

type MapSector struct {
	Id        string        `json:"id"`
	Celestial *MapCelestial `json:"celestial"`
	Objects   []*MapObject  `json:"objects"`
}

type MapSystem struct {
	Id      string        `json:"id"`
	Sectors [][]MapSector `json:"sectors"`
}

type GameMap struct {
	Systems []MapSystem `json:"systems"`
}

var pGameMap GameMap

func getGameMap(w http.ResponseWriter, r *http.Request) {
	jsonMap, _ := json.Marshal(getGameMapStruct())

	w.Write(jsonMap)
}

func getGameMapStruct() GameMap {
	// if we have it persisted, no reason to check disk
	if len(pGameMap.Systems) > 0 {
		return pGameMap
	}

	mapFile, err := os.Open("./static/map.json")

	if err != nil {
		fmt.Println(err)
	}

	defer mapFile.Close()

	byteValue, _ := ioutil.ReadAll(mapFile)

	var gameMap GameMap

	umErr := json.Unmarshal(byteValue, &gameMap)

	if umErr != nil {
		fmt.Println(umErr)
	}

	pGameMap = gameMap

	return gameMap
}
