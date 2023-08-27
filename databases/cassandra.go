package database

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gocql/gocql"
)

type CassandraDB struct {
	Instance *gocql.Session
}

func (c *CassandraDB) Init() {
	DB_CASSANDRA_USERNAME := os.Getenv("DB_CASSANDRA_USERNAME")
	DB_CASSANDRA_PAWSSWORD := os.Getenv("DB_CASSANDRA_PAWSSWORD")
	DB_CASSANDRA_CLUSTERIP := os.Getenv("DB_CASSANDRA_CLUSTERIP")

	// connect to the cluster
	cluster := gocql.NewCluster(DB_CASSANDRA_CLUSTERIP)
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	cluster.ConnectTimeout = time.Second * 10
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: DB_CASSANDRA_USERNAME, Password: DB_CASSANDRA_PAWSSWORD, AllowedAuthenticators: []string{"com.instaclustr.cassandra.auth.InstaclustrPasswordAuthenticator"}}
	session, err := cluster.CreateSession()

	if err != nil {
		log.Println(err)
		return
	}
	c.Instance = session
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

	q := c.Instance.Query("SELECT database_version FROM configurations").Iter()
	q.Scan(&databaseVersion)

	if databaseVersion != 0 && databaseVersion == GetDatabaseVersion() {
		log.Println("Database Version: ", strconv.Itoa(databaseVersion), " No need to update")
		return
	}

	log.Println("Database Version: ", strconv.Itoa(databaseVersion), " Updating it...")

}

func GetDatabaseVersion() int {
	files, _ := os.ReadDir("./databases/sqls/cassandra")
	return len(files)
}
