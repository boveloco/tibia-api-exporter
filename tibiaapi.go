package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var TIBIA_API_URL = "api.tibiadata.com"
var API_GET_WORLDS = "v3/worlds"
var API_GET_STATISTICS = "v3/killstatistics"

type CreatureStatistic struct {
	Name  string `json:"race"`
	Count int    `json:"last_day_killed"`
}

type World struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Pvp      string `json:"pvp_type"`
	Battleye string `json:"battleye_date"`
	Online   int    `json:"players_online"`
	Status   string `json:"status"`
	Premium  bool   `json:"premium_only"`
}

type Worlds struct {
	Players_Online int     `json:"players_online"`
	Worlds         []World `json:"regular_worlds"`
}

type WorldRequest struct {
	Request Worlds `json:"Worlds"`
}

type CreatureStatistics struct {
	World      string              `json:"world"`
	Statistics []CreatureStatistic `json:"entries"`
}

type CreatureStatisticsRequest struct {
	Request CreatureStatistics `json:"killstatistics"`
}

func getStatistics(world string) []CreatureStatistic {
	var uri = fmt.Sprintf("https://%s/%s/%s", TIBIA_API_URL, API_GET_STATISTICS, world)
	log.Printf("Getting statistics for world: %s", world)

	var s CreatureStatisticsRequest
	req, err := http.Get(uri)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		return nil
	}

	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&s)

	if err != nil {
		return nil
	}
	return s.Request.Statistics
}

func getWorlds() []World {
	var uri = fmt.Sprintf("https://%s/%s", TIBIA_API_URL, API_GET_WORLDS)

	var w WorldRequest
	req, err := http.Get(uri)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		return nil
	}

	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&w)

	if err != nil {
		return nil
	}

	return w.Request.Worlds
}
