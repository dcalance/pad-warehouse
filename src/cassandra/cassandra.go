package cassandra

import (
	"github.com/gocql/gocql"
)

var Session *gocql.Session

func Init(clusterIp string) {
	var err error

	cluster := gocql.NewCluster(clusterIp)
	cluster.Keyspace = "cycling"
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
}
