package models

import (
	"gopkg.in/mgo.v2"
	"time"
)

func GetSession() *mgo.Database {
	info := &mgo.DialInfo{
		Addrs:    []string{ADDRESS}, //[]string{"ds113692.mlab.com:13692"},
		Timeout:  60 * time.Second,
		Database: DB,       // "orsungo-accessmanagement",
		Username: USERNAME, //"arganherokudb1",
		Password: PASSWORD, // "$erver25082017",
	}
	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}

	db := session.DB(DB)
	return db;
}
