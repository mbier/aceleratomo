package mongo

import (
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

const (
	MongoDBHosts = "localhost:27017"
	AuthDatabase = "acelerato"
	AuthUserName = ""
	AuthPassword = ""
	TestDatabase = "acelerato-test"
)

func GetMongoSession() *mgo.Session {

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{MongoDBHosts},
		Timeout:  60 * time.Second,
		Database: AuthDatabase,
		Username: AuthUserName,
		Password: AuthPassword,
	}

	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}

	mongoSession.SetMode(mgo.Monotonic, true)

	return mongoSession;
}