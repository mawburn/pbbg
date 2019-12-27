package main

import (
  "fmt"
  "math/rand"
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
	Objects   []SectorObject `json:"objects"`
	Players   []SectorPlayer `json:"players"`
}

func getSector(id string) {
}

func generateSectors(c *redis.Client) {
	m := getGameMapStruct()

	for _, sys := range m.Systems {
		for _, row := range sys.Sectors {
			for _, col := range row {
				var s Sector

				s.Id = col.Id
				s.SystemId = sys.Id
				s.Celestial = col.Celestial
        s.Players = []SectorPlayer{}
        
				for _, obj := range col.Objects {
          if obj == nil {
            s.Objects = append(s.Objects, nil)
            continue
          }
          
					var so SectorObject

					so.Id = obj.Id
					so.Type = obj.Type
					so.Size = obj.Size

					var q int

					if obj.Size == "lg" {
						q = rand.Intn(LG_MAX-MD_MAX) + MD_MAX + 1
					} else if obj.Size == "md" {
						q = rand.Intn(MD_MAX-SM_MAX) + SM_MAX + 1
					} else {
						q = rand.Intn(SM_MAX-10000) + 10000
					}
          
          so.Max = uint32(q)
          so.Quantity = uint32(q)
          
          s.Objects = append(s.Objects, so)
				}
        
         j, _ := json.Marshal(s)
         fmt.Println(string(j))
			}
		}
	}
}
