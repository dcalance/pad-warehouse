package cassandra

import (
	"github.com/gocql/gocql"
)

var Session *gocql.Session

func Init() {
	var err error

	cluster := gocql.NewCluster("cassandra")
	cluster.Keyspace = "cycling"
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
}
