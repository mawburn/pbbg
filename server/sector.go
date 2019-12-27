package main

type SectorPlayer struct {
	Id string `json:"id"`
}

type SectorObject struct {
	Id       string  `json:"id"`
	Type     string  `json:"type"`
	Size     string  `json:"size"`
	Quantity *uint32 `json:"quantity"`
}

type Sector struct {
	Id        string          `json:"id"`
	Celestial *Celestial      `json:"celestial"`
	Objects   []*SectorObject `json:"objects"`
	Players   []*SectorPlayer `json:`
}

func getSector(id string) { 
}
