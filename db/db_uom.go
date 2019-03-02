package db

import (
	"context"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/mongo/options"

	"github.com/mongodb/mongo-go-driver/bson"

	"logika/plm/models"
)

// CreateUOM create new uom
func CreateUOM(uom models.UnitOfMeasure) (string, error) {
	log.Println("Create a uom:", uom.Code)

	res, err := db.Collection("plm_uoms").InsertOne(context.TODO(), uom)
	if err != nil {
		log.Print(err)
		return "", err
	}

	log.Println("Inserted uom code: ", res.InsertedID)
	return fmt.Sprintf("%v", res.InsertedID), nil
	// fmt.Println("Inserted uom: ", uom.Code)
}

// UpdateUOM update uom
func UpdateUOM(code string, uom models.UnitOfMeasure) (int64, error) {
	log.Println("Update uom:", uom.Code)

	filter := bson.M{"code": code}
	data := bson.D{
		{"$set", bson.D{
			{"name", uom.Name},
		}},
		{"$currentDate", bson.D{
			{"lastModified", true},
		}},
	}

	res, err := db.Collection("plm_uoms").UpdateOne(context.TODO(), filter, data)
	if err != nil {
		log.Print(err)
		return 0, err
	}

	log.Println("Updated uom code: ", res.UpsertedID)
	return res.UpsertedCount, nil
}

// DeleteUOM get type
func DeleteUOM(code string) int64 {
	log.Println("Delete ptype: ", code)

	filter := bson.M{"code": code}
	res, err := db.Collection("plm_uoms").DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Print(err)
	}

	log.Println("Delete Count : ", res.DeletedCount)
	return res.DeletedCount
}

// GetUOM get uom
func GetUOM(code string) models.UnitOfMeasure {
	log.Println("Get uom:", code)

	filter := bson.M{"code": code}
	var uom models.UnitOfMeasure
	err := db.Collection("plm_uoms").FindOne(context.TODO(), filter).Decode(&uom)
	if err != nil {
		log.Print(err)
	}

	log.Println("Get uom: ", uom.Code)
	return uom
}

// GetAllUOM get uom
func GetAllUOM() []models.UnitOfMeasure {
	log.Println("Get all uom")

	var plm_uoms []models.UnitOfMeasure
	cur, err := db.Collection("plm_uoms").Find(context.TODO(), bson.D{})
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	} else {
		for cur.Next(context.TODO()) {
			var uom models.UnitOfMeasure
			err := cur.Decode(&uom)
			if err != nil {
				log.Print(err)
			} else {
				plm_uoms = append(plm_uoms, uom)
			}
		}
	}

	log.Println("Return all uom: ", plm_uoms)
	return plm_uoms
}

// GetLimitUOM get uom
func GetLimitUOM(offset int64, limit int64) []models.UnitOfMeasure {
	log.Println("Get limit uom")

	options := options.Find()
	options.SetLimit(limit)
	options.SetSkip(offset)

	var plm_uoms []models.UnitOfMeasure
	cur, err := db.Collection("plm_uoms").Find(context.TODO(), bson.D{}, options)
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	} else {
		for cur.Next(context.TODO()) {
			var uom models.UnitOfMeasure
			err := cur.Decode(&uom)
			if err != nil {
				log.Print(err)
			} else {
				plm_uoms = append(plm_uoms, uom)
			}

		}
	}

	log.Println("Return all uom: ", plm_uoms)
	return plm_uoms
}

// CountUOM cout all uom
func CountUOM() int64 {
	log.Println("Count all uom")

	cnt, err := db.Collection("plm_uoms").Count(context.TODO(), bson.D{})
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	}

	log.Println("Count all uom: ", cnt)
	return cnt
}

// SearchUOM get uom
func SearchUOM(text string, offset int64, limit int64) []models.UnitOfMeasure {
	log.Println("Search uom")

	options := options.Find()
	options.SetLimit(limit)
	options.SetSkip(offset)

	filter := bson.M{"$text": bson.M{"$search": text}}

	var plm_uoms []models.UnitOfMeasure
	cur, err := db.Collection("plm_uoms").Find(context.TODO(), filter, options)
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	} else {
		for cur.Next(context.TODO()) {
			var uom models.UnitOfMeasure
			err := cur.Decode(&uom)
			if err != nil {
				log.Print(err)
			} else {
				plm_uoms = append(plm_uoms, uom)
			}
			// log.Print(uom)
		}
	}

	log.Println("Return all uom: ", plm_uoms)
	return plm_uoms
}

// SearchUOMCount cout all uom
func SearchUOMCount(text string) int64 {
	log.Println("Count all uom")

	filter := bson.M{"$text": bson.M{"$search": text}}
	cnt, err := db.Collection("plm_uoms").Count(context.TODO(), filter)
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	}

	log.Println("Count all uom: ", cnt)
	return cnt
}
