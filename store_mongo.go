package telegram

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type mongoStore struct {
	mongo *mgo.Database
}


func NewMongoStore(mongo *mgo.Database) Store {
	return &mongoStore{mongo}
}

func (s mongoStore) SetState(user int, state string, field []string) error {
	ss := State{Userid: user, State: state, Field: field}
	c := s.mongo.C("userstate")
	c.RemoveAll(bson.M{"userid": user})
	return c.Insert(&ss)
}

func (s mongoStore) getState(user int) State {
	c := s.mongo.C("userstate")
	var state State
	c.Find(bson.M{"userid": user}).One(&state)
	return state
}
