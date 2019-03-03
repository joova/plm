package db

import (
	"context"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/mongo/options"

	"github.com/mongodb/mongo-go-driver/bson"

	"logika/plm/models"
)

// CreateProductCategory create new category
func CreateProductCategory(category models.ProductCategory) (string, error) {
	log.Println("Create a category:", category.Code)

	res, err := db.Collection("plm_categories").InsertOne(context.TODO(), category)
	if err != nil {
		log.Print(err)
		return "", err
	}

	log.Println("Inserted category code: ", res.InsertedID)
	return fmt.Sprintf("%v", res.InsertedID), nil
	// fmt.Println("Inserted category: ", category.Code)
}

// UpdateProductCategory update category
func UpdateProductCategory(code string, category models.ProductCategory) (int64, error) {
	log.Println("Update category:", category.Code)

	filter := bson.M{"code": code}
	data := bson.D{
		{"$set", bson.D{
			{"name", category.Name},
		}},
		{"$currentDate", bson.D{
			{"lastModified", true},
		}},
	}

	res, err := db.Collection("plm_categories").UpdateOne(context.TODO(), filter, data)
	if err != nil {
		log.Print(err)
		return 0, err
	}

	log.Println("Updated category code: ", res.UpsertedID)
	return res.UpsertedCount, nil
}

// GetProductCategory get category
func GetProductCategory(code string) models.ProductCategory {
	log.Println("Get category:", code)

	filter := bson.M{"code": code}
	var category models.ProductCategory
	err := db.Collection("plm_categories").FindOne(context.TODO(), filter).Decode(&category)
	if err != nil {
		log.Print(err)
	}

	log.Println("Get category: ", category.Code)
	return category
}

// DeleteProductCategory get category
func DeleteProductCategory(code string) int64 {
	log.Println("Delete category:", code)

	filter := bson.M{"code": code}
	res, err := db.Collection("plm_categories").DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Print(err)
	}

	log.Println("Delete Count : ", res.DeletedCount)
	return res.DeletedCount
}

// GetAllProductCategory get category
func GetAllProductCategory() []models.ProductCategory {
	log.Println("Get all category")

	var categories []models.ProductCategory
	cur, err := db.Collection("plm_categories").Find(context.TODO(), bson.D{})
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	} else {
		for cur.Next(context.TODO()) {
			var category models.ProductCategory
			err := cur.Decode(&category)
			if err != nil {
				log.Print(err)
			} else {
				categories = append(categories, category)
			}
		}
	}

	log.Println("Return all category: ", categories)
	return categories
}

// GetLimitProductCategory get category
func GetLimitProductCategory(offset int64, limit int64) []models.ProductCategory {
	log.Println("Get limit category")

	options := options.Find()
	options.SetLimit(limit)
	options.SetSkip(offset)

	var categories []models.ProductCategory
	cur, err := db.Collection("plm_categories").Find(context.TODO(), bson.D{}, options)
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	} else {
		for cur.Next(context.TODO()) {
			var category models.ProductCategory
			err := cur.Decode(&category)
			if err != nil {
				log.Print(err)
			} else {
				categories = append(categories, category)
			}

		}
	}

	log.Println("Return all category: ", categories)
	return categories
}

// CountProductCategory cout all category
func CountProductCategory() int64 {
	log.Println("Count all category")

	cnt, err := db.Collection("plm_categories").Count(context.TODO(), bson.D{})
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	}

	log.Println("Count all category: ", cnt)
	return cnt
}

// SearchProductCategory get category
func SearchProductCategory(text string, offset int64, limit int64) []models.ProductCategory {
	log.Println("Search category")

	options := options.Find()
	options.SetLimit(limit)
	options.SetSkip(offset)

	filter := bson.M{"$text": bson.M{"$search": text}}

	var categories []models.ProductCategory
	cur, err := db.Collection("plm_categories").Find(context.TODO(), filter, options)
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	} else {
		for cur.Next(context.TODO()) {
			var category models.ProductCategory
			err := cur.Decode(&category)
			if err != nil {
				log.Print(err)
			} else {
				categories = append(categories, category)
			}
			// log.Print(category)
		}
	}

	log.Println("Return all category: ", categories)
	return categories
}

// SearchProductCategoryCount cout all category
func SearchProductCategoryCount(text string) int64 {
	log.Println("Count all category")

	filter := bson.M{"$text": bson.M{"$search": text}}
	cnt, err := db.Collection("plm_categories").Count(context.TODO(), filter)
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	}

	log.Println("Count all category: ", cnt)
	return cnt
}
