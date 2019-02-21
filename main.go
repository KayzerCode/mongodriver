package main

import (
	"fmt"
	"os"

	"github.com/globalsign/mgo"
	log "github.com/sirupsen/logrus"
)

const (
	url = "localhost"
)

func main() {

	dbName := ""
	// Get dbName from Args
	if 1 == len(os.Args) {
		log.Warnf("No dbName specified, listing all dbNames")
		listAllDbs()
	} else {
		dbName = os.Args[1]
		dbInfo(dbName)
	}

	// Look up for dbName
	// db := session.DB(dbName)
	// if db == nil {
	// 	// not found
	// 	log.Errorf("db `%v` not found, listing all dbNames", dbName)
	// } else {
	// 	// listing collections in db
	// 	cols, err := db.CollectionNames()
	// 	if err != nil {
	// 		log.Warnf("No collections in db `%v`", dbName)
	// 	}
	// 	fmt.Printf("Collectiond in db `%v`:\n", dbName)
	// 	for i, v := range cols {
	// 		fmt.Printf("[%3v] - %v\n", i+1, v)
	// 	}

	// }
}

func listAllDbs() {

	// Connect to Mongodb server
	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	log.Infof("Successfully connected to mongodb server at %v", url)

	// listing all dbs
	dbNames, err := session.DatabaseNames()
	if err != nil {
		log.Warn(err)
	}
	for i, v := range dbNames {
		fmt.Printf("[%3v] - %v\n", i+1, v)
	}
	return
}

func dbInfo(dbName string) {
	// Connect to Mongodb server
	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	log.Infof("Successfully connected to mongodb server at %v", url)

	fmt.Printf("Searching for '%v' on server..\n", dbName)

	dbNames, err := session.DatabaseNames()
	if err != nil {
		log.Warn(err)
	} else {
		for i, v := range dbNames {
			if v == dbName {
				fmt.Printf("Use %v\n", v)
				return
			}
			i++
		}
		fmt.Printf("Db '%v' is not found \n", dbName)
	}
}
