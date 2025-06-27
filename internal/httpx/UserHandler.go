package httpx

import (
	"crm-backend/internal/models"
	"crm-backend/internal/respository"
	"crm-backend/internal/service"
	"crm-backend/pkg/util"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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

func (h *Handler) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &models.User{}
		if ok := decodeDataFromReq(w, r, user); !ok {
			return
		}
		err, resp := h.UserService.CreateUser(user)
		if err != nil {
			util.InternalErrHandler(w)
			return
		}
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
