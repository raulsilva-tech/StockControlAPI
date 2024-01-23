package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/raulsilva-tech/StockControlAPI/internal/dto"
	"github.com/raulsilva-tech/StockControlAPI/internal/entity"
	"github.com/raulsilva-tech/StockControlAPI/internal/infra/database"
)

type ProductHandler struct {
	DAO database.ProductDAO
}

func NewProductHandler(dao database.ProductDAO) *ProductHandler {
	return &ProductHandler{DAO: dao}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	//getting body request
	var dto dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}

	//creating a new instance in memory
	found, err := entity.NewProduct(dto.Id, dto.Description, entity.ProductType{Id: dto.ProductTypeId})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}

	//inserting into the database
	err = h.DAO.Create(found)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	param := chi.URLParam(r, "id")
	if param == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseInt(param, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}
	//getting body request
	var record entity.Product
	err = json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}

	_, err = h.DAO.FindById(int(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Error{Message: "record not found - " + err.Error()})
		return
	}

	record.UpdatedAt = time.Now()
	//updating record in database
	err = h.DAO.Update(&record)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {

	param := chi.URLParam(r, "id")
	if param == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseInt(param, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}

	found, err := h.DAO.FindById(int(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Error{Message: " record not found - " + err.Error()})
		return
	}

	//deleting record from the database
	err = h.DAO.Delete(found)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {

	param := chi.URLParam(r, "id")
	if param == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}

	found, err := h.DAO.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Error{Message: " record not found - " + err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(found)
	w.WriteHeader(http.StatusOK)

}

func (h *ProductHandler) GetAllProduct(w http.ResponseWriter, r *http.Request) {

	foundList, err := h.DAO.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foundList)
	w.WriteHeader(http.StatusOK)

}
