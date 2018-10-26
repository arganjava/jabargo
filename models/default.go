package models

import (
		"os"
)

var GO_APP_ENV = os.Getenv("GO_APP_ENV")

var (
	DB       = config(os.Getenv("DB_PRODUCTION"), os.Getenv("DB_STAGING"), "jabargo")
	USERNAME = config(os.Getenv("DB_PRODUCTION_USERNAME"), os.Getenv("DB_PRODUCTION_USERNAME"), "")
	PASSWORD = config(os.Getenv("DB_PRODUCTION_PASSWORD"), os.Getenv("DB_STAGING_PASSWORD"), "")
	ADDRESS  = config(os.Getenv("DB_PRODUCTION_ADDRESS"), os.Getenv("DB_STAGING_PASSWORD"), "127.0.0.1:27017")
)

const (
	DefaultChannel = "5bcdd71c9afd13895a508af8"
)

func config(production string, staging string, local string) string {
	var val string
	func(production string, staging string) {
		switch GO_APP_ENV {
		case "production":
			val = production
		case "staging":
			val = staging
		default:
			val = local
		}
	}(production, staging)
	return val
}


/*type OperatinCRUD interface {
	Create(collection string, class interface{}) interface{}
	Update(class interface{}) interface{}
	Delete(class interface{}) interface{}
	GetAll(class interface{}) []interface{}
	GetOne(class... interface{}) interface{}
}

func Create(collection string, class *struct{}) *struct{} {

	return class;
}

func Update(collection string, class *struct{}) *struct{} {

	return class;
}

func Delete(collection string, class *struct{}) *struct{} {

	return class;
}

func ReadOne(collection string, class *struct{}) *struct{} {

	return class;
}

func ReadAll(collection string, class *[]interface{}) *[]struct{} {
	err := GetSession().C(collection).Find(nil).All(&class)
	if err != nil {
		log.Fatal(err)
	}
	return class;
}
*/