package main

import (
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	var db Database
	var wg sync.WaitGroup
	currentTime := time.Now().UTC()

	tmp := new(CassandraDB)
	db = tmp

	db.Init()
	defer db.Close()

	err := db.UpdateDatabase()
	if err != nil {
		panic("Error updating database")
	}
	valid, err := db.ValidateExecution()
	if err != nil {
		panic("Error while validating execution")
	}
	if !valid {
		today := currentTime.Format("2006-01-02")
		log.Printf("Already executed for: %s", today)
		os.Exit(0)
	}

	worlds := getWorlds()

	for _, world := range worlds {
		log.Printf("Processing World: %s", world.Name)
		wg.Add(1)
		s := getStatistics(world.Name)
		go db.WriteStatistics(s, world.Name, &wg)
	}
	wg.Wait()

	db.SetLastExecution()

}
