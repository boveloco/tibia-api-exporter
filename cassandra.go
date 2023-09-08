package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gocql/gocql"
)

var CASSANDRA_SQL_PATH = "./sqls/cassandra"
var CASSANDRA_KEYSPACE = os.Getenv("DB_CASSANDRA_KEYSTORE")

type CassandraDB struct {
	Instance *gocql.Session
}

func (c *CassandraDB) Init() {
	DB_CASSANDRA_USERNAME := os.Getenv("DB_CASSANDRA_USERNAME")
	DB_CASSANDRA_PAWSSWORD := os.Getenv("DB_CASSANDRA_PAWSSWORD")
	DB_CASSANDRA_CLUSTERIP := os.Getenv("DB_CASSANDRA_CLUSTERIP")
	log.Println("Connecting Database...")

	// connect to the cluster
	cluster := gocql.NewCluster(DB_CASSANDRA_CLUSTERIP)
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	cluster.ConnectTimeout = time.Second * 10
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: DB_CASSANDRA_USERNAME, Password: DB_CASSANDRA_PAWSSWORD, AllowedAuthenticators: []string{"org.apache.cassandra.auth.PasswordAuthenticator"}}
	session, err := cluster.CreateSession()

	if err != nil {
		log.Println(err)
		return
	}
	c.Instance = session
	log.Println("Database Connected..")
}

func (c *CassandraDB) WriteStatistics(data []CreatureStatistic, world string, res chan bool, wg *sync.WaitGroup) {
	var errRet bool = false
	log.Printf("Writing statistics for world: %s", world)
	for _, statistic := range data {
		name := strings.Replace(statistic.Name, "'", "-", -1)
		query := fmt.Sprintf("INSERT INTO %s.creature_statistics (name, count, day, world) VALUES ('%s', %d, %d, '%s');", CASSANDRA_KEYSPACE, name, statistic.Count, time.Now().Unix(), world)
		err := c.Instance.Query(query).Exec()

		if err != nil {
			log.Println(err)
			errRet = true
		}
	}
	log.Printf("Statistics Written for world: %s", world)

	res <- errRet
	wg.Done()
}

func (c *CassandraDB) Close() {
	c.Instance.Close()
}

func (c *CassandraDB) UpdateDatabase() error {

	// Return average sleep time for James
	var databaseVersion int = 0

	q := c.Instance.Query("SELECT database_version FROM " + CASSANDRA_KEYSPACE + ".configurations").Iter()
	q.Scan(&databaseVersion)

	if databaseVersion != 0 && databaseVersion == GetDatabaseVersion() {
		log.Println("Database Version: ", strconv.Itoa(databaseVersion), " No need to update")
		return nil
	}

	log.Println("Database Version: ", strconv.Itoa(databaseVersion), " Updating it...")

	files, _ := os.ReadDir(CASSANDRA_SQL_PATH)

	for i := databaseVersion; i < len(files); i++ {
		println("Applying update: ", files[i].Name())
		queries := getFileQueries(files[i].Name())
		for _, query := range queries {
			err := c.Instance.Query(query).Exec()
			if err != nil {
				return errors.New(fmt.Sprintf("Err while applying update: %s. Err: %s", files[i].Name(), err.Error()))
			}
		}
	}
	return nil
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
