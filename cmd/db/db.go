package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"os"

	driver "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"

	"gorm.io/gorm"
)

func loadEnv() *driver.Config {
	if err := godotenv.Load("../../configs/.env"); err != nil {
		log.Fatal("Loading env file err : ", err)
	}
	cfg := driver.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = os.Getenv("Net")
	cfg.Addr = os.Getenv("Addr")
	cfg.DBName = os.Getenv("DBName")

	return cfg
}

func Open(ctx context.Context, stmts []string) (*gorm.DB, error) {
	fmt.Println("Connecting Database .....")

	cfg := loadEnv()
	mysqlDB, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("mysql.Open failed: %w", err)
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: mysqlDB,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("gorm.Open failed: %w", err)
	}

	if err := mysqlDB.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("database ping failed: %w", err)
	}

	fmt.Println("Connected Database")
	return gormDB, err
}

func Close() {}
