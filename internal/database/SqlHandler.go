package database

import (
	"vk_test/internal/config"
	"vk_test/internal/database/interfaces"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SqlHandler struct {
	db *gorm.DB
}

func NewSqlHandler() interfaces.SqlHandler {
	dsn := config.DatabaseUrl
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.db = db
	return sqlHandler
}

func (handler *SqlHandler) Create(obj interface{}) {
	handler.db.Create(obj)
}

func (handler *SqlHandler) FindAll(obj interface{}) {
	handler.db.Find(obj)
}

func (handler *SqlHandler) DeleteById(obj interface{}, id string) {
	handler.db.Delete(obj, id)
}

func (handler *SqlHandler) SelectById(obj interface{}, id string) {
	handler.db.Select(obj, id)
}

func (handler *SqlHandler) Where(object interface{}, args ...interface{}) (tx *gorm.DB) {
	return handler.db.Where(object, args)
}

func (handler *SqlHandler) Preload(query string, args ...interface{}) (tx *gorm.DB) {
	return handler.db.Preload(query, args)
}
