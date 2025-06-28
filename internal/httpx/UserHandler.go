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
	countryDao := &respository.CountryDao{DB: db}

	userService := service.UserService{
		UserDao:    userDao,
		CountryDao: countryDao,
	}
	return &Handler{UserService: userService}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	if ok := util.DecodeDataFromReq(w, r, user); !ok {
		return
	}

	resp, err := h.UserService.CreateUser(user)
	if err != nil {
		util.InternalErrHandler(w)
		return
	}

	util.JSON(w, http.StatusCreated, resp)
}

func (h *Handler) CreateUserByBulk(w http.ResponseWriter, r *http.Request) {
	var user []models.User
	if ok := util.DecodeDataFromReq(w, r, &user); !ok {
		return
	}

	respWrapper := h.UserService.CreateUserByBulk(&user)
	util.JSON(w, http.StatusCreated, respWrapper)
}
