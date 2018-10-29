package models

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"log"
	"time"
	"golang.org/x/crypto/bcrypt"

	"errors"
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
		log.Panicln(&err)
		return err, nil;
	}
	pushToArray := bson.M{"$push": bson.M{"channels": bson.M{"_id": channel.Id, "name": channel.Name}}}
	user.Channels = append(user.Channels, ChannelUser{channel.Id, channel.Name})
	hash := Hash{}
	hashResult, errHash := hash.Generate(user.Password)
	if errHash != nil {
		return errHash, nil;
	}
	user.Password = hashResult
	err = userCollection().Insert(&user)
	if err != nil {
		return err, nil;
	}

	err = userCollection().Find(bson.M{"_id": user.Id}).One(&result)
	if err != nil {
		return err, nil;
	}
	pushToArray = bson.M{"$push": bson.M{"users": bson.M{"_id": user.Id, "name": user.Name}}}
	err = channelCollection().Update(bson.M{"_id": channel.Id}, pushToArray)
	if err != nil {
		return err, nil;
	}
	return nil, &result;
}

func (user User) Update() (error, *User) {
	err := userCollection().Update(user.Id, &user)
	if err != nil {
		return err, nil;
	}
	err = userCollection().Find(bson.M{"_id": user.Id}).One(&user)
	return err, &user;
}

func (user User) Delete() *User {
	result := User{}
	err := userCollection().Remove(user.Id)
	if err != nil {
		log.Fatal(err)
	}
	err = userCollection().Find(bson.M{"name": "Ale"}).One(&result)
	return &result;
}

func (user User) VerifyUser(email string) (error, *User) {
	err := userCollection().Find(bson.M{"email": email}).One(&user)
	if err != nil {
		err = errors.New("not found user or password")
		return err, nil;
	}
	return nil, &user;
}

type Hash struct{}

func (c *Hash) Generate(input string) (string, error) {
	saltedBytes := []byte(input)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash, nil
}

func (c *Hash) Compare(hash string, s string) error {
	incoming := []byte(s)
	existing := []byte(hash)
	return bcrypt.CompareHashAndPassword(existing, incoming)
}
