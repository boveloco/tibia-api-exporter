package main

import "log"

func main() {
	var db Database
	tmp := new(CassandraDB)
	db = tmp

	db.Init()
	db.UpdateDatabase()

	worlds := getWorlds()

	for _, world := range worlds {
		log.Printf("Processing World: %s", world.Name)
		s := getStatistics(world.Name)
		db.WriteStatistics(s, world.Name)
	}

	defer db.Close()

}
