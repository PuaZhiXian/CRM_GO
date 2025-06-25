package respository

import (
	"crm-backend/internal/util"
	"crm-backend/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type userDaoInterface interface {
	CreateUser(user *models.User)
	FindUserById(id string) *models.User
	UpdateUser(user *models.User)
	DeleteUserById(id string)
}

type Handler struct {
	UserDao userDaoInterface
}

func (h *Handler) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Entering Handler")
		user := &models.User{}
		err := json.NewDecoder(r.Body).Decode(user)
		if err != nil {
			log.Println("[CreateUser] ", err)
			util.RequestErrWriter(w, util.ErrBadRequest)
			return
		}

		h.UserDao.CreateUser(user)
		util.JSON(w, http.StatusCreated, "User Created")
	}
}

func (h *Handler) Add(router *mux.Router) {
	r1 := router.Methods(http.MethodPost)
	r1.Path("/create-user")
	r1.Handler(h.CreateUser())
	r1.Name("user-create")
}
