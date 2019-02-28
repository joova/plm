package handlers

import (
	"encoding/json"
	"logika/plm/db"
	"logika/plm/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateCategoryEndpoint create a pcat
func CreateCategoryEndpoint(w http.ResponseWriter, r *http.Request) {
	var pcat models.ProductCategory
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&pcat)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	pcatExist := db.GetProductCategory(pcat.Code)
	if pcatExist.Code != "" {
		msg := "Category already exist in the database"
		http.Error(w, msg, 400)
		return
	}

	db.CreateProductCategory(pcat)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pcat)

}

// UpdateCategoryEndpoint update a pcat
func UpdateCategoryEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	code := params["code"]

	var pcat models.ProductCategory
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&pcat)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	db.UpdateProductCategory(code, pcat)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pcat)

}

// GetCategoryEndpoint get a pcat
func GetCategoryEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	code := params["code"]

	var pcat models.ProductCategory
	// _ = json.NewDecoder(r.Body).Decode(&pcat)

	pcat = db.GetProductCategory(code)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pcat)

}

// GetAllCategoryEndpoint get a pcat
func GetAllCategoryEndpoint(w http.ResponseWriter, r *http.Request) {

	var pcats []models.ProductCategory
	pcats = db.GetAllProductCategory()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pcats)

}

// GetPagingCategoryEndpoint get a pcat
func GetPagingCategoryEndpoint(w http.ResponseWriter, r *http.Request) {
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

	count := db.CountProductCategory()
	scount := strconv.FormatInt(count, 10)

	page := float64(count) / float64(limit)
	page = math.Ceil(page)
	spage := strconv.FormatFloat(page, 'f', 0, 64)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Pagination-Count", scount)
	w.Header().Set("Pagination-Page", spage)
	w.Header().Set("Pagination-Limit", slimit)

	var pcats []models.ProductCategory
	pcats = db.GetLimitProductCategory(offset, limit)
	json.NewEncoder(w).Encode(pcats)

}

// SearchCategoryEndpoint search pcat
func SearchCategoryEndpoint(w http.ResponseWriter, r *http.Request) {
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

	var pcats []models.ProductCategory
	pcats = db.SearchProductCategory(text, offset, limit)

	count := db.SearchProductCategoryCount(text)
	scount := strconv.FormatInt(count, 10)

	page := float64(count) / float64(limit)
	page = math.Ceil(page)
	spage := strconv.FormatFloat(page, 'f', 0, 64)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Pagination-Count", scount)
	w.Header().Set("Pagination-Page", spage)
	w.Header().Set("Pagination-Limit", slimit)

	json.NewEncoder(w).Encode(pcats)

}

// CoutCategoryEndpoint get a pcat
func CoutCategoryEndpoint(w http.ResponseWriter, r *http.Request) {

	var count int64
	count = db.CountProductCategory()
	res := map[string]int64{"count": count}
	json.NewEncoder(w).Encode(res)

}
