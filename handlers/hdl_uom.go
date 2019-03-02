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

// CreateUOMEndpoint create a uom
func CreateUOMEndpoint(w http.ResponseWriter, r *http.Request) {
	var uom models.UnitOfMeasure
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&uom)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	uomExist := db.GetUOM(uom.Code)
	if uomExist.Code != "" {
		msg := "UOM already exist in the database"
		http.Error(w, msg, 400)
		return
	}

	db.CreateUOM(uom)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(uom)

}

// UpdateUOMEndpoint update a uom
func UpdateUOMEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	code := params["code"]

	var uom models.UnitOfMeasure
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&uom)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	db.UpdateUOM(code, uom)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(uom)

}

// DeleteUOMEndpoint get a ptype
func DeleteUOMEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	code := params["code"]

	// _ = json.NewDecoder(r.Body).Decode(&pcat)

	count := db.DeleteUOM(code)
	res := map[string]int64{"deleted": count}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}

// GetUOMEndpoint get a uom
func GetUOMEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	code := params["code"]

	var uom models.UnitOfMeasure
	// _ = json.NewDecoder(r.Body).Decode(&uom)

	uom = db.GetUOM(code)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(uom)

}

// GetAllUOMEndpoint get a uom
func GetAllUOMEndpoint(w http.ResponseWriter, r *http.Request) {

	var uoms []models.UnitOfMeasure
	uoms = db.GetAllUOM()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(uoms)

}

// GetPagingUOMEndpoint get a uom
func GetPagingUOMEndpoint(w http.ResponseWriter, r *http.Request) {
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

	count := db.CountUOM()
	scount := strconv.FormatInt(count, 10)

	page := float64(count) / float64(limit)
	page = math.Ceil(page)
	spage := strconv.FormatFloat(page, 'f', 0, 64)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Pagination-Count", scount)
	w.Header().Set("Pagination-Page", spage)
	w.Header().Set("Pagination-Limit", slimit)

	var uoms []models.UnitOfMeasure
	uoms = db.GetLimitUOM(offset, limit)
	json.NewEncoder(w).Encode(uoms)

}

// SearchUOMEndpoint search uom
func SearchUOMEndpoint(w http.ResponseWriter, r *http.Request) {
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

	var uoms []models.UnitOfMeasure
	uoms = db.SearchUOM(text, offset, limit)

	count := db.SearchUOMCount(text)
	scount := strconv.FormatInt(count, 10)

	page := float64(count) / float64(limit)
	page = math.Ceil(page)
	spage := strconv.FormatFloat(page, 'f', 0, 64)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Pagination-Count", scount)
	w.Header().Set("Pagination-Page", spage)
	w.Header().Set("Pagination-Limit", slimit)

	json.NewEncoder(w).Encode(uoms)

}

// CoutUOMEndpoint get a uom
func CoutUOMEndpoint(w http.ResponseWriter, r *http.Request) {

	var count int64
	count = db.CountUOM()
	res := map[string]int64{"count": count}
	json.NewEncoder(w).Encode(res)

}
