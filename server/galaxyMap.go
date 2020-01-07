package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// MapObject - objects in space as represtented on the map
type MapObject struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Max  uint32 `json:"max"`
}

// MapCelestial - celestials in space as represtented on the map
type MapCelestial struct {
	ID   string `json:"id"`
	Name string `"json:"name"`
	Type string `json:"type"`
}

// MapSector - sectors as represented on the map
type MapSector struct {
	SystemID  string        `json:"systemId"`
	Xpos      uint8         `json:"x"`
	Ypos      uint8         `json:"y"`
	Celestial *MapCelestial `json:"celestial"`
	Objects   []*MapObject  `json:"objects"`
}

// GalaxyMap - the full game map
type GalaxyMap struct {
	Sectors map[string]MapSector  `json:"sectors"`
	Systems map[string][][]string `json:"systems"`
}

var galaxyMap GalaxyMap

func getGalaxyMap(w http.ResponseWriter, r *http.Request) {
	jsonMap, _ := json.Marshal(getGalaxyMapStruct())

	w.Write(jsonMap)
}

func getGalaxyMapStruct() GalaxyMap {
	// if we have it persisted, no reason to check disk
	if len(galaxyMap.Systems) > 0 {
		return galaxyMap
	}

	mapFile, err := os.Open("./static/map.json")

	if err != nil {
		fmt.Println(err)
	}

	defer mapFile.Close()

	byteValue, _ := ioutil.ReadAll(mapFile)

	var gameMap GalaxyMap

	umErr := json.Unmarshal(byteValue, &gameMap)

	if umErr != nil {
		fmt.Println(umErr)
	}

	galaxyMap = gameMap

	return gameMap
}
