package main

import (
	"encoding/json"

	"github.com/go-redis/redis/v7"
)

var LG_MAX int = 10000000
var MD_MAX int = 1000000
var SM_MAX int = 100000

type SectorPlayer struct {
	Id string `json:"id"`
}

type SectorObject struct {
	MapObject
	Max      uint32 `json:"max"`
	Quantity uint32 `json:"quantity"`
}

type Sector struct {
	Id        string          `json:"id"`
	SystemId  string          `json:"-"`
	Celestial *MapCelestial   `json:"celestial"`
	Objects   []*SectorObject `json:"objects"`
	Players   []SectorPlayer  `json:"players"`
}

func getSector(id string) {

}

func updatePlayers(c *redis.Client, sectorId string, add []SectorPlayer, remove []SectorPlayer) {
	sectorByte, err := c.Get(sectorId).Bytes()

	if err != nil {
		panic(err)
	}

	var s Sector

	umErr := json.Unmarshal(sectorByte, &s)

	if umErr != nil {
		panic(umErr)
	}

	var newPlayers []SectorPlayer

	for _, sp := range s.Players {
		if !containsPlayer(sp, remove) {
			newPlayers = append(newPlayers, SectorPlayer{Id: sp.Id})
		}
	}

	for _, a := range add {
		newPlayers = append(newPlayers, a)
	}

	j, _ := json.Marshal(s)

	redisErr := c.Set(s.Id, j, 0).Err()

	if redisErr != nil {
		panic(redisErr)
	}
}

func containsPlayer(p SectorPlayer, list []SectorPlayer) bool {
	for _, li := range list {
		if li.Id == p.Id {
			return true
		}
	}

	return false
}

func generateSectors(c *redis.Client) {
	m := getGameMapStruct()

	for _, sys := range m.Systems {
		for _, row := range sys.Sectors {
			for _, col := range row {
				s := Sector{
					Id:        col.Id,
					SystemId:  sys.Id,
					Celestial: col.Celestial,
					Players:   []SectorPlayer{},
				}

				for _, obj := range col.Objects {
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

				j, _ := json.Marshal(s)

				err := c.Set(s.Id, j, 0).Err()

				if err != nil {
					panic(err)
				}
			}
		}
	}
}
