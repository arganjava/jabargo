package models

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

const (
	CollectionUser = "user"
)

// Article model
type User struct {
	Id        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name      string        `json:"name" form:"name" binding:"required" bson:"name"`
	Email     string        `json:"email" form:"email"  bson:"email"`
	Company   string        `json:"company" form:"company" bson:"company"`
	Password  string        `json:"password" form:"password" bson:"password"`
	Privilege string        `json:"privilege" form:"privilege" bson:"privilege"`
	Phone     string        `json:"phone" form:"phone" bson:"phone"`
	Channels  []ChannelUser `json:"channels" form:"channels" bson:"channels"`
	CreatedAt time.Time     `json:"createdAt" bson:"createdAt"`
	UpdateAt  time.Time     `json:"updateAt" bson:"updateAt"`
	CreatedBy time.Time     `json:"createdBy" bson:"createdBy"`
	UpdateBy  time.Time     `json:"updateBy" bson:"updateBy"`
}

type ChannelUser struct {
	Id   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string        `json:"name" form:"name" binding:"required" bson:"name"`
}

func userCollection() *mgo.Collection {

	return GetSession().C(CollectionUser)
}

func (user User) GET(mapParam *map[string]interface{}) (error, *[]User) {
	result := []User{}
	pipe := userCollection().Pipe([]bson.M{
		{"$match": bson.M{
			"channels._id": bson.M{"$in": []bson.ObjectId{bson.ObjectIdHex(DefaultChannel)}},
		}}})
	err := pipe.All(&result);
	if err != nil {
		log.Fatal(err)
		return err, nil;
	}
	return err, &result;
}

// once user registered will have default channel called social
// and channel.users appended new user
func (user User) POST() (error, *User) {
	result := User{}
	channel := Channel{}
	err := channelCollection().Find(bson.M{"_id": bson.ObjectIdHex(DefaultChannel)}).One(&channel)
	if err != nil {
		log.Fatal(err)
		return err, nil;
	}
	pushToArray := bson.M{"$push": bson.M{"channels": bson.M{"_id": channel.Id, "name": channel.Name}}}
	user.Id = bson.NewObjectId()
	user.Channels = append(user.Channels, ChannelUser{channel.Id, channel.Name})
	err = userCollection().Insert(&user)
	if err != nil {
		log.Fatal(err)
		return err, nil;
	}
	//err = userCollection().Update(bson.M{"_id": user.Id}, pushToArray)

	err = userCollection().Find(bson.M{"_id": user.Id}).One(&result)
	if err != nil {
		log.Fatal(err)
		return err, nil;
	}
	pushToArray = bson.M{"$push": bson.M{"users": bson.M{"_id": user.Id, "name": user.Name}}}
	err = channelCollection().Update(bson.M{"_id": channel.Id}, pushToArray)
	if err != nil {
		log.Fatal(err)
		return err, nil;
	}
	return nil, &result;
}

func (this User) PUT(user *User) *User {
	result := User{}
	err := userCollection().Insert(&user)
	if err != nil {
		log.Fatal(err)
	}
	err = userCollection().Find(bson.M{"name": "Ale"}).One(&result)
	return &result;
}

func (this User) DELETE(user *User) *User {
	result := User{}
	err := userCollection().Insert(&user)
	if err != nil {
		log.Fatal(err)
	}
	err = userCollection().Find(bson.M{"name": "Ale"}).One(&result)
	return &result;
}
