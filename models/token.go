package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Token struct {
	Id          bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	AccessToken string        `json:"accessToken,omitempty" bson:"accessToken"`
	User        UserChannel   `json:"user,omitempty" bson:"user"`
	CreatedAt   time.Time     `json:"createdAt" bson:"createdAt"`
	UpdateAt    time.Time     `json:"updateAt" bson:"updateAt"`
}
