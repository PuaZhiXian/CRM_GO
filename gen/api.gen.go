// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// CreateUserRespFailed defines model for CreateUserRespFailed.
type CreateUserRespFailed struct {
	Age         *uint32 `json:"age,omitempty"`
	Name        *string `json:"name,omitempty"`
	Nationality *string `json:"nationality,omitempty"`
	Reason      *string `json:"reason,omitempty"`
	Residential *string `json:"residential,omitempty"`
}

// CreateUserRespSuccess defines model for CreateUserRespSuccess.
type CreateUserRespSuccess struct {
	Age         *uint32 `json:"age,omitempty"`
	Name        *string `json:"name,omitempty"`
	Nationality *string `json:"nationality,omitempty"`
	Residential *string `json:"residential,omitempty"`
}

// CreateUserRespWrapper defines model for CreateUserRespWrapper.
type CreateUserRespWrapper struct {
	FailedUser  *[]CreateUserRespFailed  `json:"failedUser,omitempty"`
	SuccessUser *[]CreateUserRespSuccess `json:"successUser,omitempty"`
}

// User defines model for User.
type User struct {
	Age         *int32  `json:"age,omitempty"`
	Name        *string `json:"name,omitempty"`
	Nationality *string `json:"nationality,omitempty"`
	Residential *string `json:"residential,omitempty"`
	Row         *int32  `json:"row,omitempty"`
}

// CreateUserByBulkJSONBody defines parameters for CreateUserByBulk.
type CreateUserByBulkJSONBody = []User

// CreateUserJSONRequestBody defines body for CreateUser for application/json ContentType.
type CreateUserJSONRequestBody = User

// CreateUserByBulkJSONRequestBody defines body for CreateUserByBulk for application/json ContentType.
type CreateUserByBulkJSONRequestBody = CreateUserByBulkJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create a new user
	// (POST /create-user)
	CreateUser(w http.ResponseWriter, r *http.Request)
	// Create multiple new user
	// (POST /create-user-by-bulk)
	CreateUserByBulk(w http.ResponseWriter, r *http.Request)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// Create a new user
// (POST /create-user)
func (_ Unimplemented) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Create multiple new user
// (POST /create-user-by-bulk)
func (_ Unimplemented) CreateUserByBulk(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// CreateUser operation middleware
func (siw *ServerInterfaceWrapper) CreateUser(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateUser(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// CreateUserByBulk operation middleware
func (siw *ServerInterfaceWrapper) CreateUserByBulk(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateUserByBulk(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/create-user", wrapper.CreateUser)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/create-user-by-bulk", wrapper.CreateUserByBulk)
	})

	return r
}
