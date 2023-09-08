package main

import (
	"log"
	"sync"
)

func main() {
	var db Database
	var wg sync.WaitGroup

	tmp := new(CassandraDB)
	db = tmp

	db.Init()
	err := db.UpdateDatabase()

	if err != nil {
		panic("Error updating database")
	}
	worlds := getWorlds()

	// channel responses
	resChan := make(chan bool, len(worlds))

	for _, world := range worlds {
		log.Printf("Processing World: %s", world.Name)
		s := getStatistics(world.Name)
		go db.WriteStatistics(s, world.Name, resChan, &wg)
		wg.Add(1)
	}
	wg.Wait()

	defer db.Close()

}
