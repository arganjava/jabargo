package models

import (
	"gopkg.in/mgo.v2/bson"
	"log"
	"gopkg.in/mgo.v2"
	"time"
)

const (
	// CollectionArticle holds the name of the articles collection
	CollectionChannels = "channel"
)

func channelCollection() *mgo.Collection {

	return GetSession().C(CollectionChannels)
}

type Channel struct {
	Id        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string        `json:"name" form:"name" binding:"required" bson:"title"`
	Type      string        `json:"type" form:"type" binding:"required" bson:"type"`
	Users     []UserChannel `json:"users" form:"users" bson:"users"`
	Chat      []Message     `json:"chat" form:"chats" bson:"chat"`
	CreatedAt time.Time     `json:"createdAt" bson:"createdAt"`
	UpdateAt  time.Time     `json:"updateAt" bson:"updateAt"`
}

type UserChannel struct {
	Id        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name      string        `json:"name" form:"name" binding:"required" bson:"name"`
	Email     string        `json:"email" form:"email"  bson:"email"`
	Company   string        `json:"company" form:"company" bson:"company"`
	Privilege string        `json:"privilege" form:"privilege" bson:"privilege"`
	Phone     string        `json:"phone" form:"phone" bson:"phone"`
}

type Message struct {
	Id           bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Sender       string        `json:"sender" form:"sender" binding:"required" bson:"sender"`
	Acknowledges []*string     `json:acknowledges" form:"acknowledges" bson:"acknowledges"`
	CreatedAt    time.Time     `json:"createdAt" bson:"createdAt"`
	UpdateAt     time.Time     `json:"updateAt" bson:"updateAt"`
}

func (channel *Channel) Post() (error, *Channel) {
	result := Channel{}
	channel.Id = bson.NewObjectId()
	err := channelCollection().Insert(&channel)
	if err != nil {
		log.Fatal(err)
	}
	err = channelCollection().Find(bson.M{"_id": channel.Id}).One(&result)
	return err, &result;
}

func (this Channel) Get(channel *Channel) (error, *[]Channel) {
	result := []Channel{}
	err := channelCollection().Find(nil).All(&result)
	if err != nil {
		log.Fatal(err)
	}
	return err, &result;
}

func (channel Channel) AddUser(idUser *bson.ObjectId) (error, *Channel) {
	user := User{}
	channelId := bson.M{"_id": channel.Id}
	err := channelCollection().Find(bson.M{"_id": channel.Id}).One(&channel)
	err = userCollection().Find(bson.M{"_id": idUser}).One(&user)
	pushToArray := bson.M{"$push": bson.M{"users": &user}}
	err = channelCollection().Update(channelId, pushToArray)
	pushToArray = bson.M{"$push": bson.M{"channels": &channel}}
	err = userCollection().Update(bson.M{"_id": user.Id}, pushToArray)

	updateChannel := Channel{}
	err = channelCollection().Find(bson.M{"_id": channel.Id}).One(&updateChannel)
	if err != nil {
		log.Fatal(err)
	}
	return err, &updateChannel;
}

func (channel Channel) RemoveUser(idUser *bson.ObjectId) (error, *Channel) {
	user := User{}
	channelId := bson.M{"_id": channel.Id}
	err := channelCollection().Find(bson.M{"_id": channel.Id}).One(&channel)
	err = userCollection().Find(bson.M{"_id": idUser}).One(&user)
	pushToArray := bson.M{"$pull": bson.M{"users": &user}}
	err = channelCollection().Update(channelId, pushToArray)
	pushToArray = bson.M{"$pull": bson.M{"channels": &channel}}
	err = userCollection().Update(bson.M{"_id": user.Id}, pushToArray)

	updateChannel := Channel{}
	err = channelCollection().Find(bson.M{"_id": channel.Id}).One(&updateChannel)
	if err != nil {
		log.Fatal(err)
	}
	return err, &updateChannel;
}
