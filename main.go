package main

func main() {
	var db Database
	tmp := new(CassandraDB)
	db = tmp

	db.Init()
	db.UpdateDatabase()

	worls := getWorlds()
	println(worls)
	defer db.Close()

}
