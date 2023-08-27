package main

func main() {
	var db Database
	tmp := new(CassandraDB)
	db = tmp

	db.Init()
	println("Conetado no banco...")
	db.UpdateDatabase()

	defer db.Close()

}
