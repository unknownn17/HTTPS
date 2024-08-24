package handler

import (
	"api/internal/kafka/adjust"
	"api/internal/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	A *adjust.Adjust
}

// Create godoc
// @Summary Create a new item
// @Description Creates a new item and broadcasts the creation event via Kafka.
// @Tags items
// @Accept  json
// @Produce  json
// @Param item body models.CreateItemRequest true "Item to create"
// @Success 200 {object} map[string]string
// @Failure 500 {object} models.Error
// @Router /items/create [post]
func (u *Handler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.CreateItemRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := u.A.Broadcast("create", req); err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(map[string]string{"Message": "Your request is in process"})
	time.Sleep(time.Second * 2)
	res, err := u.A.Create()
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(res)
}

// Get godoc
// @Summary Get an item by ID
// @Description Retrieves an item by its ID.
// @Tags items
// @Accept  json
// @Produce  json
// @Param id path string true "Item ID"
// @Success 200 {object} models.GeneralItem
// @Failure 500 {object} models.Error
// @Router /items/{id} [get]
func (u *Handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")
	res, err := u.A.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

// Gets godoc
// @Summary Get all items
// @Description Retrieves all items.
// @Tags items
// @Accept  json
// @Produce  json
// @Success 200 {array} models.GeneralItem
// @Failure 500 {object} models.Error
// @Router /items [get]
func (u *Handler) Gets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res, err := u.A.Gets()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

// Update godoc
// @Summary Update an item by ID
// @Description Updates an existing item by its ID and broadcasts the update event via Kafka.
// @Tags items
// @Accept  json
// @Produce  json
// @Param id path int true "Item ID"
// @Param item body models.GeneralItem true "Item to update"
// @Success 200 {object} map[string]string
// @Failure 500 {object} models.Error
// @Router /items/{id} [put]
func (u *Handler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var req models.GeneralItem

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.ID = int32(id)

	if err := u.A.Broadcast("update", req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"Message": fmt.Sprintf("Item has been updated with this id %v", id)})
}

// Delete godoc
// @Summary Delete an item by ID
// @Description Deletes an existing item by its ID and broadcasts the delete event via Kafka.
// @Tags items
// @Accept  json
// @Produce  json
// @Param id path int true "Item ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} models.Error
// @Router /items/{id} [delete]
func (u *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := u.A.Broadcast("update", models.GeneralItem{ID: int32(id)}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"Message": fmt.Sprintf("Item has been deleted with this id %v", id)})
}
