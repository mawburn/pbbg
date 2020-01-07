package main

import (
	"encoding/json"
	"fmt"
)

const LG_MAX int = 10000000
const MD_MAX int = 1000000
const SM_MAX int = 100000

// SectorObject - objects represented as the player is looking at them
type SectorObject struct {
	MapObject
	Quantity uint32 `json:"quantity"`
}

// Sector - sectors as the player is looking at them
type Sector struct {
	SystemID  string          `json:"systemId"`
	Celestial *MapCelestial   `json:"celestial"`
	Objects   []*SectorObject `json:"objects"`
	Players   []string        `json:"players"`
}

func getSector(ID string) (Sector, error) {
	sectorResult, err := dbConns.Redis.Get("sector:"+ID).Bytes()

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

func movePlayer(userID string, fromSector string, toSector string) {
	user := []string{userID}

	updatePlayers(fromSector, []string{}, user)
	updatePlayers(toSector, user, []string{})
}

/**
* built for when queing comes into effect
**/
func updatePlayers(sectorID string, add []string, remove []string) {
	sectorByte, err := dbConns.Redis.Get("sector:"+sectorID).Bytes()

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

	redisErr := dbConns.Redis.Set("sector:"+sectorID, j, 0).Err()

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

	for secID, sec := range m.Sectors {
		s := Sector{
			SystemID:  sec.SystemID,
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
					ID:   obj.ID,
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

		err = dbConns.Redis.Set("sector:" +secID, j, 0).Err()

		if err != nil {
			panic(err)
		}
	}
}
