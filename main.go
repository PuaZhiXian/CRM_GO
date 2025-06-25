package main

import (
	"crm-backend/db"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/schema"
)

type CoinEnq struct {
	Username string
}
type CoinResp struct {
	Username string
	Amt      int64
}

func mainV2() {
	var r *chi.Mux = chi.NewRouter()

	Handler(r)
	fmt.Println("Application Started")

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Fatalln(err)
	}
}

func Handler(r *chi.Mux) {
	r.Use(middleware.StripSlashes)
	r.Route("/account", func(router chi.Router) {
		router.Use(Authorization)
		router.Get("/coins", RealFunc)
	})
}

var ErrUnauthorized = errors.New("unauthorized")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var authHeader = r.Header.Get("Authorization")

		const bearerPrefix = "Bearer "

		var token string
		if strings.HasPrefix(authHeader, bearerPrefix) {
			token = strings.TrimPrefix(authHeader, bearerPrefix)
			fmt.Println("Token is:", token)
		} else {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		if username == "" || token == "" {
			RequestErrWriter(w, ErrUnauthorized)
			return
		}

		var database *db.DbInterface
		var err error
		database, err = db.NewDb()
		if err != nil {
			InternalErrorHandler(w)
			return
		}

		loginDtl := (*database).GetUserLoginDtl(username)
		if loginDtl == nil || loginDtl.AuthToken != token {
			RequestErrWriter(w, ErrUnauthorized)
		}

		next.ServeHTTP(w, r)
	})
}

func RealFunc(w http.ResponseWriter, r *http.Request) {
	var params = CoinEnq{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())
	if err != nil {
		log.Fatalln(err)
		return
	}

	var database *db.DbInterface
	database, err = db.NewDb()
	if err != nil {
		InternalErrorHandler(w)
		return
	}

	(*database).GetUserCoins(params.Username)

	coin := (*database).GetUserCoins(params.Username)

	var resp = CoinResp{
		Username: coin.Username,
		Amt:      coin.Coins,
	}

	w.Header().Set("Content-Type", "applicantion/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Fatalln(err)
		return
	}
}

type Error struct {
	Code    int
	Message string
}


func writeErr(w http.ResponseWriter, msg string, code int) {
	resp := Error{
		Code:    code,
		Message: msg,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrWriter = func(w http.ResponseWriter, err error) {
		writeErr(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeErr(w, "Unexpected Err", http.StatusInternalServerError)
	}
)
