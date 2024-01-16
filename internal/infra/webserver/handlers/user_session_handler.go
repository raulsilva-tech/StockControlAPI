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

type UserSessionHandler struct {
	DAO database.UserSessionDAO
}

func NewUserSessionHandler(dao database.UserSessionDAO) *UserSessionHandler {
	return &UserSessionHandler{DAO: dao}
}

func (h *UserSessionHandler) CreateUserSession(w http.ResponseWriter, r *http.Request) {

	//getting body request
	var dto dto.CreateUserSessionInput
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}

	//creating a new instance in memory
	found, err := entity.NewUserSession(dto.Id, entity.User{Id: dto.UserId}, dto.StartedAt, dto.FinishedAt)
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

func (h *UserSessionHandler) UpdateUserSession(w http.ResponseWriter, r *http.Request) {

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
	var record entity.UserSession
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

func (h *UserSessionHandler) DeleteUserSession(w http.ResponseWriter, r *http.Request) {

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

func (h *UserSessionHandler) GetUserSession(w http.ResponseWriter, r *http.Request) {

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

func (h *UserSessionHandler) GetAllUserSession(w http.ResponseWriter, r *http.Request) {

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
