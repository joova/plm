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

// CreateTypeEndpoint create a ptype
func CreateTypeEndpoint(w http.ResponseWriter, r *http.Request) {
	var ptype models.ProductType
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&ptype)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	ptypeExist := db.GetProductType(ptype.Code)
	if ptypeExist.Code != "" {
		msg := "Type already exist in the database"
		http.Error(w, msg, 400)
		return
	}

	db.CreateProductType(ptype)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ptype)

}

// UpdateTypeEndpoint update a ptype
func UpdateTypeEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	code := params["code"]

	var ptype models.ProductType
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&ptype)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	db.UpdateProductType(code, ptype)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ptype)

}

// DeleteTypeEndpoint get a ptype
func DeleteTypeEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	code := params["code"]

	// _ = json.NewDecoder(r.Body).Decode(&pcat)

	count := db.DeleteProductType(code)
	res := map[string]int64{"deleted": count}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}

// GetTypeEndpoint get a ptype
func GetTypeEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	code := params["code"]

	var ptype models.ProductType
	// _ = json.NewDecoder(r.Body).Decode(&ptype)

	ptype = db.GetProductType(code)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ptype)

}

// GetAllTypeEndpoint get a ptype
func GetAllTypeEndpoint(w http.ResponseWriter, r *http.Request) {

	var ptypes []models.ProductType
	ptypes = db.GetAllProductType()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ptypes)

}

// GetPagingTypeEndpoint get a ptype
func GetPagingTypeEndpoint(w http.ResponseWriter, r *http.Request) {
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

	count := db.CountProductType()
	scount := strconv.FormatInt(count, 10)

	page := float64(count) / float64(limit)
	page = math.Ceil(page)
	spage := strconv.FormatFloat(page, 'f', 0, 64)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Pagination-Count", scount)
	w.Header().Set("Pagination-Page", spage)
	w.Header().Set("Pagination-Limit", slimit)

	var ptypes []models.ProductType
	ptypes = db.GetLimitProductType(offset, limit)
	json.NewEncoder(w).Encode(ptypes)

}

// SearchTypeEndpoint search ptype
func SearchTypeEndpoint(w http.ResponseWriter, r *http.Request) {
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

	var ptypes []models.ProductType
	ptypes = db.SearchProductType(text, offset, limit)

	count := db.SearchProductTypeCount(text)
	scount := strconv.FormatInt(count, 10)

	page := float64(count) / float64(limit)
	page = math.Ceil(page)
	spage := strconv.FormatFloat(page, 'f', 0, 64)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Pagination-Count", scount)
	w.Header().Set("Pagination-Page", spage)
	w.Header().Set("Pagination-Limit", slimit)

	json.NewEncoder(w).Encode(ptypes)

}

// CoutTypeEndpoint get a ptype
func CoutTypeEndpoint(w http.ResponseWriter, r *http.Request) {

	var count int64
	count = db.CountProductType()
	res := map[string]int64{"count": count}
	json.NewEncoder(w).Encode(res)

}
