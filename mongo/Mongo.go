package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/healthy-service/model"
	"log"
)

var session *mgo.Session

var DBName = "healthy"

func GetSession() (*mgo.Session, error) {
	if session == nil {
		session, err := mgo.Dial("localhost:27017")
		if err != nil {
			return nil, err
		}
		return session, err
	}

	newSession := session.Copy()

	return newSession, nil
}

func GetAllFoods() ([]model.Food, error) {
	session, _ := GetSession()
	defer session.Close()
	db := session.DB(DBName)

	foods := make([]model.Food, 0)
	err := db.C("foods").Find(bson.M{}).All(&foods)
	logError(err)

	return foods, err
}

func GetAllRecipes() ([]model.Recipe, error) {
	session, _ := GetSession()
	defer session.Close()
	db := session.DB(DBName)

	recipes := make([]model.Recipe, 0)
	err := db.C("recipes").Find(bson.M{}).All(&recipes)
	logError(err)

	return recipes, err
}

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateFood(food model.Food) (model.Food, error) {
	session, _ := GetSession()
	defer session.Close()
	db := session.DB(DBName)

	err := db.C("foods").UpdateId(food.ID, &food)

	logError(err)

	return food, err
}

func AddFood(food model.Food) (model.Food, error) {
	session, _ := GetSession()
	defer session.Close()
	db := session.DB(DBName)

	err := db.C("foods").Insert(&food)
	logError(err)

	return food, err
}
