package db

import (
	"api/config"
	"api/model"
	"errors"
	"github.com/kamagasaki/go-utils/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertPinjaman(requestData model.Pinjamann) error {
	db := mongo.MongoConnect(DBATS)
	insertedID := mongo.InsertOneDoc(db, config.PinjamanColl, requestData)
	if insertedID == nil {
		return errors.New("couldn't insert data")
	}
	return nil
}

func GetAllPinjaman(filter bson.M) ([]model.Pinjamann, error) {
	db := mongo.MongoConnect(DBATS)
	var datas []model.Pinjamann
	err := mongo.GetAllDocByFilter[model.Pinjamann](db, config.PinjamanColl, filter, &datas)
	if err != nil {
		return nil, err
	}
	return datas, nil
}

func GetOnePinjaman(filter bson.M) (model.Pinjamann, error) {
	db := mongo.MongoConnect(DBATS)
	var data model.Pinjamann
	err := mongo.GetOneDocByFilter[model.Pinjamann](db, config.PinjamanColl, filter, &data)
	if err != nil {
		return model.Pinjamann{}, err
	}
	return data, nil
}
