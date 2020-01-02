package main

import (
	"encoding/json"
	"fmt"
)

var LG_MAX int = 10000000
var MD_MAX int = 1000000
var SM_MAX int = 100000

type SectorObject struct {
	MapObject
	Quantity uint32 `json:"quantity"`
}

type Sector struct {
	SystemId  string          `json:"systemId"`
	Celestial *MapCelestial   `json:"celestial"`
	Objects   []*SectorObject `json:"objects"`
	Players   []string        `json:"players"`
}

func getSector(id string) (Sector, error) {
	sectorResult, err := dbConns.Redis.Get(id).Bytes()

	if err != nil {
		return Sector{}, fmt.Errorf("Unable to get sector")
	}

	var sector Sector

	err = json.Unmarshal(sectorResult, &sector)

	if err != nil {
		return Sector{}, fmt.Errorf("Unable to get sector")
	}

	return sector, nil
}

func movePlayer(userId string, fromSector string, toSector string) {
	user := []string{userId}

	updatePlayers(fromSector, []string{}, user)
	updatePlayers(toSector, user, []string{})
}

/**
* built for when queing comes into effect
**/
func updatePlayers(sectorId string, add []string, remove []string) {
	sectorByte, err := dbConns.Redis.Get(sectorId).Bytes()

	if err != nil {
		panic(err)
	}

	var s Sector

	umErr := json.Unmarshal(sectorByte, &s)

	if umErr != nil {
		panic(umErr)
	}

	var newPlayers []string

	for _, sp := range s.Players {
		if !containsPlayer(sp, remove) {
			newPlayers = append(newPlayers, sp)
		}
	}

	for _, a := range add {
		newPlayers = append(newPlayers, a)
	}

	if newPlayers != nil {
		s.Players = newPlayers
	} else {
		s.Players = []string{}
	}

	j, _ := json.Marshal(s)

	redisErr := dbConns.Redis.Set(sectorId, j, 0).Err()

	if redisErr != nil {
		panic(redisErr)
	}
}

func containsPlayer(p string, list []string) bool {
	for _, li := range list {
		if li == p {
			return true
		}
	}

	return false
}

func generateSectors() {
	m := getGalaxyMapStruct()

	for secId, sec := range m.Sectors {
		s := Sector{
			SystemId:  sec.SystemId,
			Celestial: sec.Celestial,
			Players:   []string{},
		}

		for _, obj := range sec.Objects {
			if obj == nil {
				s.Objects = append(s.Objects, nil)
				continue
			}

			so := &SectorObject{
				MapObject: MapObject{
					Id:   obj.Id,
					Type: obj.Type,
					Max:  obj.Max,
				},
				Quantity: obj.Max,
			}

			s.Objects = append(s.Objects, so)
		}

		j, err := json.Marshal(s)

		if err != nil {
			panic(err)
		}

		err = dbConns.Redis.Set(secId, j, 0).Err()

		if err != nil {
			panic(err)
		}
	}
}
