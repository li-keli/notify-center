package db

import (
	"github.com/globalsign/mgo"
	"github.com/sirupsen/logrus"
)

var globalS *mgo.Session

func NewMongo() {
	var url = "172.16.2.161:21000"

	dialInfo := &mgo.DialInfo{
		Addrs:  []string{url},
		Source: "notification_db",
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		logrus.Fatal("mongodb connection error: ", err, url)
	}
	session.SetMode(mgo.Monotonic, true)
	globalS = session
}

func connect(db, collection string) (*mgo.Session, *mgo.Collection) {
	s := globalS.Copy()
	c := s.DB(db).C(collection)
	return s, c
}

func Insert(db, collection string, docs ...interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Insert(docs...)
}

func FindOne(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).One(result)
}

func FindAll(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).All(result)
}

func Update(db, collection string, query, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Update(query, update)
}

func Remove(db, collection string, query interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Remove(query)
}
