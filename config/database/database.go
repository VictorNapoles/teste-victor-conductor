package database

import (
	"fmt"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type (
	// IDatabase - Interface de operações com o banco de dados
	IDatabase interface {
		Insert(entity *interface{})
		Update(entity *interface{})
		Find(dest *[]interface{}, entity *interface{}, inputs interface{})
		Delete(entity interface{}, id ...interface{})
	}

	// Database - Implementação das operações com o banco de dados
	Database struct {
		db *gorm.DB
	}
)

var (
	instance Database
	once     sync.Once
)

func load() {
	db, err := gorm.Open(sqlite.Open("conductor.db"), &gorm.Config{})
	fmt.Println("Conexão com o banco criada")

	if err != nil {
		panic("failed to connect database")
	}

	instance = Database{db: db}
}

func GetDatabase() IDatabase {
	once.Do(func() {
		load()
	})
	return &instance
}

// Insert - Inserção de registro no banco de dados
func (database *Database) Insert(entity *interface{}) {
	database.db.Create(entity)
}

// Update - Alteração de registro no banco de dados
func (database *Database) Update(entity *interface{}) {
	database.db.Save(entity)
}

// Find - Consulta de registros no banco de dados
func (database *Database) Find(dest *[]interface{}, entity *interface{}, inputs interface{}) {
	database.db.Where(entity, inputs).Find(dest)
}

// Delete - Exclusão de registro no banco de dados
func (database *Database) Delete(entity interface{}, id ...interface{}) {
	database.db.Delete(entity, id)
}
