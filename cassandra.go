package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gocql/gocql"
)

var CASSANDRA_SQL_PATH = "./sqls/cassandra"
var CASSANDRA_KEYSPACE = os.Getenv("DB_CASSANDRA_KEYSTORE")

type CassandraDB struct {
	Instance *gocql.Session
}

func (c *CassandraDB) Init() {
	// DB_CASSANDRA_USERNAME := os.Getenv("DB_CASSANDRA_USERNAME")
	// DB_CASSANDRA_PAWSSWORD := os.Getenv("DB_CASSANDRA_PAWSSWORD")
	DB_CASSANDRA_CLUSTERIP := os.Getenv("DB_CASSANDRA_CLUSTERIP")
	log.Println("Connecting Database...")

	// connect to the cluster
	cluster := gocql.NewCluster(DB_CASSANDRA_CLUSTERIP)
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	cluster.ConnectTimeout = time.Second * 10
	session, err := cluster.CreateSession()

	if err != nil {
		log.Println(err)
		return
	}
	c.Instance = session
	log.Println("Database Connected..")
}

func (c *CassandraDB) Write(data *interface{}) bool {
	err := c.Instance.Query("INSERT INTO sleep_centre.sleep_study (name, study_date, sleep_time_hours) VALUES ('James', '2018-01-07', 8.2);").Exec()
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (c *CassandraDB) Close() {
	c.Instance.Close()
}

func (c *CassandraDB) UpdateDatabase() {

	// Return average sleep time for James
	var databaseVersion int = 0

	q := c.Instance.Query("SELECT database_version FROM " + CASSANDRA_KEYSPACE + ".configurations").Iter()
	q.Scan(&databaseVersion)

	if databaseVersion != 0 && databaseVersion == GetDatabaseVersion() {
		log.Println("Database Version: ", strconv.Itoa(databaseVersion), " No need to update")
		return
	}

	log.Println("Database Version: ", strconv.Itoa(databaseVersion), " Updating it...")

	files, _ := os.ReadDir(CASSANDRA_SQL_PATH)

	for i := databaseVersion; i < len(files); i++ {
		println("Applying update: ", files[i].Name())
		queries := getFileQueries(files[i].Name())
		for _, query := range queries {
			err := c.Instance.Query(query).Exec()
			if err != nil {
				log.Fatal("Err while applying update: ", files[i].Name(), err)
			}
		}
	}
}

func getFileQueries(file string) (queries []string) {
	f, err := os.Open(CASSANDRA_SQL_PATH + "/" + file)

	// Check for the error that occurred during the opening of the file
	if err != nil {
		fmt.Println(err)
	}

	// read the file line by line using a scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		s := scanner.Text()
		s = strings.Replace(s, "<<KEYSPACE>>", CASSANDRA_KEYSPACE, -1)

		queries = append(queries, s)
	}
	// check for the error that occurred during the scanning

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	// Close the file
	defer f.Close()

	return []string(queries)
}

func GetDatabaseVersion() int {
	files, _ := os.ReadDir(CASSANDRA_SQL_PATH)
	return len(files)
}
