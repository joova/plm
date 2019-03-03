package db

import (
	"context"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo/options"

	"github.com/mongodb/mongo-go-driver/bson"

	"logika/plm/models"
)

// CreateProduct create new product
func CreateProduct(product models.Product) (string, error) {
	log.Println("Create a product:", product.Name)

	res, err := db.Collection("products").InsertOne(context.TODO(), product)
	if err != nil {
		log.Print(err)
		return "", err
	}

	log.Println("Inserted product id: ", res.InsertedID)
	return fmt.Sprintf("%v", res.InsertedID), nil
	// fmt.Println("Inserted user: ", user.Username)
}

// UpdateProduct update product
func UpdateProduct(id primitive.ObjectID, product models.Product) (int64, error) {
	log.Println("Update product:", product.Name)

	filter := bson.M{"_id": id}
	data := bson.D{
		{"$set", bson.D{
			{"code", product.Code},
			{"name", product.Name},
			{"description", product.Description},
			{"uom", product.UOM},
			{"type", product.Type},
			{"category", product.Category},
		}},
		{"$currentDate", bson.D{
			{"lastModified", true},
		}},
	}

	res, err := db.Collection("products").UpdateOne(context.TODO(), filter, data)
	if err != nil {
		log.Print(err)
		return 0, err
	}

	log.Println("Updated res id: ", res.UpsertedID)
	return res.UpsertedCount, nil
}

// GetProduct get product
func GetProduct(id primitive.ObjectID) models.Product {
	log.Println("Get product:", id)

	filter := bson.M{"_id": id}
	var product models.Product
	err := db.Collection("products").FindOne(context.TODO(), filter).Decode(&product)
	if err != nil {
		log.Print(err)
	}

	log.Println("Get product: ", product.Name)
	return product
}

// DeleteProduct delete product
func DeleteProduct(id primitive.ObjectID) int64 {
	log.Println("Delete product: ", id)

	filter := bson.M{"_id": id}
	res, err := db.Collection("products").DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Print(err)
	}

	log.Println("Delete Count : ", res.DeletedCount)
	return res.DeletedCount
}

// GetProductByCode get product by product code
func GetProductByCode(code string) models.Product {
	log.Println("Get product:", code)

	filter := bson.M{"code": code}
	var product models.Product
	err := db.Collection("products").FindOne(context.TODO(), filter).Decode(&product)
	if err != nil {
		log.Print(err)
	}

	log.Println("Get product: ", product.Code)
	return product
}

// GetProductByName get product by product name
func GetProductByName(name string) models.Product {
	log.Println("Get product:", name)

	filter := bson.M{"name": name}
	var product models.Product
	err := db.Collection("products").FindOne(context.TODO(), filter).Decode(&product)
	if err != nil {
		log.Print(err)
	}

	log.Println("Get product: ", product.Name)
	return product
}

// GetAllProduct get all product
func GetAllProduct() []models.Product {
	log.Println("Get all product")

	var products []models.Product
	cur, err := db.Collection("products").Find(context.TODO(), bson.D{})
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	} else {
		for cur.Next(context.TODO()) {
			var product models.Product
			err := cur.Decode(&product)
			if err != nil {
				log.Print(err)
			} else {
				products = append(products, product)
			}

		}
	}

	log.Println("Return all product: ", products)
	return products
}

// GetLimitProduct limited product
func GetLimitProduct(offset int64, limit int64) []models.Product {
	log.Println("Get limit product")

	options := options.Find()
	options.SetLimit(limit)
	options.SetSkip(offset)

	var products []models.Product
	cur, err := db.Collection("products").Find(context.TODO(), bson.D{}, options)
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	} else {
		for cur.Next(context.TODO()) {
			var product models.Product
			err := cur.Decode(&product)
			if err != nil {
				log.Print(err)
			} else {
				products = append(products, product)
			}

		}
	}

	log.Println("Return all product ")
	return products
}

// CountProduct cout all product
func CountProduct() int64 {
	log.Println("Count all product")

	cnt, err := db.Collection("products").Count(context.TODO(), bson.D{})
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	}

	log.Println("Count all product: ", cnt)
	return cnt
}

// SearchProduct search product
func SearchProduct(text string, offset int64, limit int64) []models.Product {
	log.Println("Search product")

	options := options.Find()
	options.SetLimit(limit)
	options.SetSkip(offset)

	filter := bson.M{"$text": bson.M{"$search": text}}

	var products []models.Product
	cur, err := db.Collection("products").Find(context.TODO(), filter, options)
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	} else {
		for cur.Next(context.TODO()) {
			var product models.Product
			err := cur.Decode(&product)
			if err != nil {
				log.Print(err)
			} else {
				products = append(products, product)
			}
			// log.Print(product)
		}
	}

	log.Println("Return all product")
	return products
}

// SearchProductCount cout product
func SearchProductCount(text string) int64 {
	log.Println("Count search product")

	filter := bson.M{"$text": bson.M{"$search": text}}
	cnt, err := db.Collection("products").Count(context.TODO(), filter)
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	}

	log.Println("Count search product: ", cnt)
	return cnt
}
