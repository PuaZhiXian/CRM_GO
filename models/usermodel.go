package models

import "time"

type Entity struct {
	Id          string    `json:"id"`
	CreatedDate time.Time `json:"createdDate"`
	UpdatedDate time.Time `josn:"updatedDate"`
}

type User struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	CreatedDate time.Time `json:"createdDate"`
	UpdatedDate time.Time `json:"updatedDate"`
}
