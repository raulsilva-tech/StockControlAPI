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

type StockHandler struct {
	DAO database.StockDAO
}

func NewStockHandler(dao database.StockDAO) *StockHandler {
	return &StockHandler{DAO: dao}
}

func (h *StockHandler) CreateStock(w http.ResponseWriter, r *http.Request) {

	//getting body request
	var dto dto.CreateStockInput
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}

	//creating a new instance in memory
	record, err := entity.NewStock(dto.Id, dto.Description)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}

	//inserting into the database
	err = h.DAO.Create(record)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func (h *StockHandler) UpdateStock(w http.ResponseWriter, r *http.Request) {

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
	var record entity.Stock
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

func (h *StockHandler) DeleteStock(w http.ResponseWriter, r *http.Request) {

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
		json.NewEncoder(w).Encode(Error{Message: "record not found - " + err.Error()})
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

func (h *StockHandler) GetStock(w http.ResponseWriter, r *http.Request) {

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
		json.NewEncoder(w).Encode(Error{Message: "record not found - " + err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(found)
	w.WriteHeader(http.StatusOK)

}

func (h *StockHandler) GetAllStock(w http.ResponseWriter, r *http.Request) {

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
