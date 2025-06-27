package httpx

import (
	"crm-backend/internal/models"
	"crm-backend/internal/respository"
	"crm-backend/internal/service"
	"crm-backend/pkg/util"
	"net/http"

	"gorm.io/gorm"
)

type Handler struct {
	UserService service.UserService
}

func InitHandle(db *gorm.DB) *Handler {
	userDao := &respository.UserDao{DB: db}
	userService := service.UserService{UserDao: userDao}
	return &Handler{UserService: userService}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	if ok := util.DecodeDataFromReq(w, r, user); !ok {
		return
	}

	err, resp := h.UserService.CreateUser(user)
	if err != nil {
		util.InternalErrHandler(w)
		return
	}

	util.JSON(w, http.StatusCreated, resp)
}
