package handlers

import (
	"encoding/json"
	"logika/plm/db"
	"logika/plm/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// CreateProductEndpoint create a product
func CreateProductEndpoint(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	productExist := db.GetProductByCode(product.Code)
	if productExist.Code != "" {
		msg := "Product already exist in the database"
		http.Error(w, msg, 400)
		return
	}

	product.ID = primitive.NewObjectID()

	db.CreateProduct(product)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)

}

// UpdateProductEndpoint update a product
func UpdateProductEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var product models.Product
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	db.UpdateProduct(oid, product)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)

}

// DeleteProductEndpoint get a ptype
func DeleteProductEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// _ = json.NewDecoder(r.Body).Decode(&pcat)

	count := db.DeleteProduct(oid)
	res := map[string]int64{"deleted": count}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}

// GetProductByNameEndpoint get a product
func GetProductByNameEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	var product models.Product
	_ = json.NewDecoder(r.Body).Decode(&product)

	product = db.GetProductByName(name)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)

}

// GetProductEndpoint get a product
func GetProductEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var product models.Product
	_ = json.NewDecoder(r.Body).Decode(&product)

	product = db.GetProduct(oid)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)

}

// GetAllProductEndpoint get a product
func GetAllProductEndpoint(w http.ResponseWriter, r *http.Request) {

	var products []models.Product
	products = db.GetAllProduct()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)

}

// GetPagingProductEndpoint get a product
func GetPagingProductEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	slimit := params["limit"]
	soffset := params["offset"]

	// parser limit to int
	limit, err := strconv.ParseInt(slimit, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// parser offset to int
	offset, err := strconv.ParseInt(soffset, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	count := db.CountProduct()
	scount := strconv.FormatInt(count, 10)

	page := float64(count) / float64(limit)
	page = math.Ceil(page)
	spage := strconv.FormatFloat(page, 'f', 0, 64)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Pagination-Count", scount)
	w.Header().Set("Pagination-Page", spage)
	w.Header().Set("Pagination-Limit", slimit)

	var products []models.Product
	products = db.GetLimitProduct(offset, limit)
	json.NewEncoder(w).Encode(products)

}

// SearchProductEndpoint search product
func SearchProductEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	text := params["text"]
	slimit := params["limit"]
	soffset := params["offset"]

	// parser limit to int
	limit, err := strconv.ParseInt(slimit, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// parser offset to int
	offset, err := strconv.ParseInt(soffset, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var products []models.Product
	products = db.SearchProduct(text, offset, limit)

	count := db.SearchProductCount(text)
	scount := strconv.FormatInt(count, 10)

	page := float64(count) / float64(limit)
	page = math.Ceil(page)
	spage := strconv.FormatFloat(page, 'f', 0, 64)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Pagination-Count", scount)
	w.Header().Set("Pagination-Page", spage)
	w.Header().Set("Pagination-Limit", slimit)

	json.NewEncoder(w).Encode(products)

}

// CoutProductEndpoint get a product
func CoutProductEndpoint(w http.ResponseWriter, r *http.Request) {

	var count int64
	count = db.CountProduct()
	res := map[string]int64{"count": count}
	json.NewEncoder(w).Encode(res)

}
