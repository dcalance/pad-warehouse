package cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

func Init() {
	var err error

	cluster := gocql.NewCluster("127.0.0.1:9043")
	cluster.Keyspace = "cycling"
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("cassandra init done")
}
