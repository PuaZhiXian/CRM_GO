package handlers

import (
	"crm-backend/internal/util"
	"crm-backend/models"
	"crm-backend/respository"
	"crm-backend/service"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	UserService service.UserService
}

func InitHandle(db *sql.DB) *Handler {
	userDao := &respository.UserDao{DB: db}
	userService := service.UserService{UserDao: userDao}
	return &Handler{UserService: userService}
}

func (h *Handler) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &models.User{}
		if ok := decodeDataFromReq(w, r, user); !ok {
			return
		}
		resp := h.UserService.CreateUser(user)
		util.JSON(w, http.StatusCreated, resp)
	}
}

func decodeDataFromReq(w http.ResponseWriter, r *http.Request, dataSchema any) bool {
	if err := json.NewDecoder(r.Body).Decode(dataSchema); err != nil {
		util.RequestErrWriter(w, util.ErrBadRequest)
		return false
	}
	return true
}

func (h *Handler) Add(router *mux.Router) {
	r1 := router.Methods(http.MethodPost)
	r1.Path("/create-user")
	r1.Handler(h.CreateUser())
	r1.Name("user-create")
}
