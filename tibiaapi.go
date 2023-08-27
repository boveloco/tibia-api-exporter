package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var TIBIA_API_URL = "api.tibiadata.com"
var API_GET_WORLDS = "v3/worlds"
var API_GET_STATISTICS = "v3/killstatistics"

type CreatureStatistic struct {
	Name  string `json:"race"`
	Count int    `json:"last_day_killed"`
	World string
	Date  uint16
}

type World struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Pvp      string `json:"pvp_type"`
	Battleye string `json:"battleye_date"`
}

type WorldRequest struct {
	Players_Online int     `json:"players_online"`
	Worlds         []World `json:"regular_worlds"`
}

func (c *CreatureStatistic) getStatistics() {

}

func getWorlds() []World {
	var uri = fmt.Sprintf("https://%s/%s", TIBIA_API_URL, API_GET_WORLDS)

	res, err := http.Get(uri)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		return nil
	}
	println(res)

	var r *WorldRequest

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	err = json.Unmarshal(body, r)

	if err != nil {
		println(err)
		return nil
	}

	return r.Worlds
}
