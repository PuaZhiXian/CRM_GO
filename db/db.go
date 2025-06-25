package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

type LoginDetails struct {
	Username  string
	AuthToken string
}

type CoinDtl struct {
	Coins    int64
	Username string
}

type DbInterface interface {
	GetUserLoginDtl(username string) *LoginDetails
	GetUserCoins(username string) *CoinDtl
	SetupDb() error
}

func NewDb() (*DbInterface, error) {
	var db DbInterface = &fakeDb{}
	var err error = db.SetupDb()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return &db, nil
}

func Open(ctx context.Context, stmts []string) (*sql.DB, error) {
	//TODO GET FROM FUNC OR DEFINE SETTING
	fmt.Println("Connecting Database .....")
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "recordings"

	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected Database")
	return db, err
}

func Close(){}