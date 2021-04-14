package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type (

	// IDatabaseFramework - Interface com os métodos do GORM Framework
	IOrmFramework interface {
		Create(value interface{}) *gorm.DB
		Save(value interface{}) *gorm.DB
		Where(query interface{}, args ...interface{}) *gorm.DB
		Find(dest interface{}, conds ...interface{}) *gorm.DB
		First(dest interface{}, conds ...interface{}) *gorm.DB
		Delete(value interface{}, conds ...interface{}) *gorm.DB
		AutoMigrate(dst ...interface{}) error
	}

	// Database - Implementação das operações com o banco de dados
	Database struct {
		orm IOrmFramework
	}
)

var (
	instance Database
)

func init() {
	db, err := gorm.Open(sqlite.Open("conductor.db"), &gorm.Config{})

	log.Println("Conexão com o banco criada com sucesso")

	if err != nil {
		panic("failed to connect database")
	}

	instance = Database{orm: db}
}

func GetDatabase() *Database {
	return &instance
}

// Insert - Inserção de registro no banco de dados
func (database *Database) Insert(entity interface{}) {
	database.orm.Create(entity)
}

// Update - Alteração de registro no banco de dados
func (database *Database) Update(entity interface{}) {
	database.orm.Save(entity)
}

// Find - Consulta de registros no banco de dados
func (database *Database) Find(target interface{}, entity interface{}, inputs ...interface{}) {
	database.orm.Where(entity, inputs).Find(target)
}

// FindAll - Consulta de registros no banco de dados
func (database *Database) FindAll(target interface{}) {
	database.orm.Find(target)
}

// FindById - Consulta de registros no banco de dados por id
func (database *Database) FindById(target interface{}, conds ...interface{}) {
	database.orm.First(target, conds)
}

// Delete - Exclusão de registro no banco de dados
func (database *Database) Delete(entity interface{}, id ...interface{}) {
	database.orm.Delete(entity, id)
}

// GetORM - Retorna a instância do orm utilizado.
func (database *Database) GetORM() IOrmFramework {
	return database.orm
}
