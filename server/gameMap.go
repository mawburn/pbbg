package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Object struct {
	Id       string  `json:"id"`
	Type     string  `json:"type"`
	Size     string  `json:"size"`
	Quantity *uint32 `json:"quantity"`
}

type Celestial struct {
	Id   string `json:"id"`
	Name string `"json:"name"`
	Type string `json:"type"`
}

type Sector struct {
	Id        string     `json:"id"`
	Celestial *Celestial `json:"celestial"`
	Objects   []*Object  `json:"objects"`
}

type System struct {
	Id      string     `json:"id"`
	Sectors [][]Sector `json:"sectors"`
}

type GameMap struct {
	Systems []System `json:"systems"`
}

func gameMap() GameMap {
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

	return gameMap
}
