package db

import (
	"context"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/mongo/options"

	"github.com/mongodb/mongo-go-driver/bson"

	"logika/plm/models"
)

// CreateProductType create new ptype
func CreateProductType(ptype models.ProductType) (string, error) {
	log.Println("Create a ptype:", ptype.Code)

	res, err := db.Collection("plm_types").InsertOne(context.TODO(), ptype)
	if err != nil {
		log.Print(err)
		return "", err
	}

	log.Println("Inserted ptype code: ", res.InsertedID)
	return fmt.Sprintf("%v", res.InsertedID), nil
	// fmt.Println("Inserted ptype: ", ptype.Code)
}

// UpdateProductType update ptype
func UpdateProductType(code string, ptype models.ProductType) (int64, error) {
	log.Println("Update ptype:", ptype.Code)

	filter := bson.M{"code": code}
	data := bson.D{
		{"$set", bson.D{
			{"name", ptype.Name},
		}},
		{"$currentDate", bson.D{
			{"lastModified", true},
		}},
	}

	res, err := db.Collection("plm_types").UpdateOne(context.TODO(), filter, data)
	if err != nil {
		log.Print(err)
		return 0, err
	}

	log.Println("Updated ptype code: ", res.UpsertedID)
	return res.UpsertedCount, nil
}

// GetProductType get uom
func GetProductType(code string) models.ProductType {
	log.Println("Get ptype:", code)

	filter := bson.M{"code": code}
	var ptype models.ProductType
	err := db.Collection("plm_types").FindOne(context.TODO(), filter).Decode(&ptype)
	if err != nil {
		log.Print(err)
	}

	log.Println("Get ptype: ", ptype.Code)
	return ptype
}

// GetAllProductType get ptype
func GetAllProductType() []models.ProductType {
	log.Println("Get all ptype")

	var plm_types []models.ProductType
	cur, err := db.Collection("plm_types").Find(context.TODO(), bson.D{})
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	} else {
		for cur.Next(context.TODO()) {
			var ptype models.ProductType
			err := cur.Decode(&ptype)
			if err != nil {
				log.Print(err)
			} else {
				plm_types = append(plm_types, ptype)
			}
		}
	}

	log.Println("Return all ptype: ", plm_types)
	return plm_types
}

// GetLimitProductType get ptype
func GetLimitProductType(offset int64, limit int64) []models.ProductType {
	log.Println("Get limit ptype")

	options := options.Find()
	options.SetLimit(limit)
	options.SetSkip(offset)

	var plm_types []models.ProductType
	cur, err := db.Collection("plm_types").Find(context.TODO(), bson.D{}, options)
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	} else {
		for cur.Next(context.TODO()) {
			var ptype models.ProductType
			err := cur.Decode(&ptype)
			if err != nil {
				log.Print(err)
			} else {
				plm_types = append(plm_types, ptype)
			}

		}
	}

	log.Println("Return all ptype: ", plm_types)
	return plm_types
}

// CountProductType cout all ptype
func CountProductType() int64 {
	log.Println("Count all ptype")

	cnt, err := db.Collection("plm_types").Count(context.TODO(), bson.D{})
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	}

	log.Println("Count all ptype: ", cnt)
	return cnt
}

// SearchProductType get ptype
func SearchProductType(text string, offset int64, limit int64) []models.ProductType {
	log.Println("Search ptype")

	options := options.Find()
	options.SetLimit(limit)
	options.SetSkip(offset)

	filter := bson.M{"$text": bson.M{"$search": text}}

	var plm_types []models.ProductType
	cur, err := db.Collection("plm_types").Find(context.TODO(), filter, options)
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	} else {
		for cur.Next(context.TODO()) {
			var ptype models.ProductType
			err := cur.Decode(&ptype)
			if err != nil {
				log.Print(err)
			} else {
				plm_types = append(plm_types, ptype)
			}
			// log.Print(ptype)
		}
	}

	log.Println("Return all ptype: ", plm_types)
	return plm_types
}

// SearchProductTypeCount cout all ptype
func SearchProductTypeCount(text string) int64 {
	log.Println("Count all ptype")

	filter := bson.M{"$text": bson.M{"$search": text}}
	cnt, err := db.Collection("plm_types").Count(context.TODO(), filter)
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	}

	log.Println("Count all ptype: ", cnt)
	return cnt
}
