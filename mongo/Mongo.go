package mongo

import (
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

var DBName = "healthy"

func GetSession() (*mgo.Session, error) {
	if session == nil {
		session, err := mgo.Dial("localhost:27017")
		if err != nil {
			return nil,err
		}
		return session,err
	}

	newSession := session.Copy()

	return newSession,nil
}
