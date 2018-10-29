package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"gopkg.in/mgo.v2"
	"errors"
)

const (
	CollectionToken = "token"
)

func TokenCollection() *mgo.Collection {

	return GetSession().C(CollectionToken)
}

type Token struct {
	Id          bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	AccessToken string        `json:"accessToken,omitempty" bson:"accessToken"`
	UserId      bson.ObjectId `json:"userId,omitempty" bson:"userId"`
	User        UserChannel   `json:"user,omitempty" bson:"user"`
	CreatedAt   time.Time     `json:"createdAt" bson:"createdAt"`
	UpdateAt    time.Time     `json:"updateAt" bson:"updateAt"`
}

func (token Token) Create() (error, *Token) {
	token.Id = bson.NewObjectId()
	var tokenFind *Token
	err := TokenCollection().Find(bson.M{"userId": token.UserId}).One(&tokenFind);
	if err != nil && tokenFind == nil {
		err = TokenCollection().Insert(&token)
		return err, &token;
	}
	token.UpdateAt = time.Now()
	token.CreatedAt = tokenFind.CreatedAt
	err = TokenCollection().RemoveId(tokenFind.Id)
	if err != nil {
		return err, nil;
	}
	err = TokenCollection().Insert(&token)
	if err != nil {
		return err, nil;
	}
	return nil, &token;
}

func (token Token) FindToken(accessToken string) (error, *Token) {
	token.Id = bson.NewObjectId()
	var tokenFind *Token
	err := TokenCollection().Find(bson.M{"accessToken": accessToken}).One(&tokenFind);
	if err == nil && tokenFind != nil {
		return err, tokenFind;
	}
	err = errors.New("token not found")
	return err, nil;
}
