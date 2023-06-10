package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
)

var db *sql.DB
var dbOnce sync.Once
var dbErr error

// Подключение к БД
func connectionDb() (*sql.DB, error) {
	dbOnce.Do(func() {
		db, dbErr = sql.Open("mysql", "root:@/pet_projects") //newuser:password@/pet_projects
		if dbErr != nil {
			log.Printf("Ошибка подключения к базе данных: %v", dbErr)
			return
		}
	})

	return db, dbErr
}
