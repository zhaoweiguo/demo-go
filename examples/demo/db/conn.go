package db

import (
	"gopkg.in/mgo.v2"
	"log"
)

var collection *mgo.Collection

func Conn(host, db, table string) {
	session, err := mgo.Dial(host)
	if err != nil {
		log.Println("connect mongoDB fail!")
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	collection = session.DB(db).C(table)
}
