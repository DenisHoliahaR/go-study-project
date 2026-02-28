package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/DenisHoliahaR/go-beautyhub/internal/domain"
	"github.com/DenisHoliahaR/go-beautyhub/internal/service"
	"github.com/DenisHoliahaR/go-beautyhub/internal/transport/http/dto"
	"github.com/go-chi/chi/v5"
)

type handler struct {
	service *services.UserService
	logger  *slog.Logger
}

func NewUserHandler(service *services.UserService, logger *slog.Logger) *handler {
	return &handler{
		service: service,
		logger:  logger,
	}
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	var req dto.CreateUserRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		h.logger.Warn("invalid create user request", "error", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	user := domain.User{
		FirstName:  req.FirstName,
		SecondName: req.SecondName,
		Email:      req.Email,
		Phone:      req.Phone,
	}

	createdUser, err := h.service.CreateUser(r.Context(), &user, req.Password)
	if err != nil {
		h.logger.Error("failed to create user", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	Write(w, http.StatusCreated, dto.CreateUserResponse{
		ID:         createdUser.ID,
		FirstName:  createdUser.FirstName,
		SecondName: createdUser.SecondName,
		Email:      createdUser.Email,
		Phone:      createdUser.Phone,
		CreatedAt:  createdUser.CreatedAt,
	})
}

func (h *handler) GetUserById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		slog.Error("Request with invalid Id", "error", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUserById(r.Context(), id)
	if err != nil {
		slog.Error("Error getting user by ID", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Write(w, http.StatusOK, user)
}

func (h *handler) GetUserList(w http.ResponseWriter, r *http.Request) {
	userList, err := h.service.GetUserList(r.Context())
	if err != nil {
		h.logger.Error("error getting user list", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	response := dto.GetUserListResponse{
		Users: make([]dto.GetUserResponse, 0, len(userList)),
	}

	for _, user := range userList {
		response.Users = append(response.Users, dto.GetUserResponse{
			ID:         user.ID,
			FirstName:  user.FirstName,
			SecondName: user.SecondName,
			Email:      user.Email,
			Phone:      user.Phone,
			CreatedAt:  user.CreatedAt,
		})
	}

	Write(w, http.StatusOK, response)
}


func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	var req dto.UpdateUserRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		h.logger.Warn("invalid update user request", "error", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	user := domain.User{
		FirstName:  req.FirstName,
		SecondName: req.SecondName,
		Email:      req.Email,
		Phone:      req.Phone,
	}

	updatedUser, err := h.service.UpdateUser(r.Context(), &user, req.Password)
	if err != nil {
		h.logger.Error("failed to create user", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	Write(w, http.StatusCreated, dto.CreateUserResponse{
		ID:         updatedUser.ID,
		FirstName:  updatedUser.FirstName,
		SecondName: updatedUser.SecondName,
		Email:      updatedUser.Email,
		Phone:      updatedUser.Phone,
		CreatedAt:  updatedUser.CreatedAt,
	})
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		slog.Error("Request with invalid Id", "error", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteUser(r.Context(), id); err != nil {
		slog.Error("Error deleting user by ID", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Write(w, http.StatusOK, nil)
}
