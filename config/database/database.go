package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type (
	// IDatabase - Interface de operações com o banco de dados
	IDatabase interface {
		Insert(entity interface{})
		Update(entity interface{})
		Find(target interface{}, entity interface{}, inputs interface{})
		Delete(entity interface{}, id ...interface{})
	}

	// IDatabaseFramework - Interface com os métodos do GORM Framework
	IDatabaseFramework interface {
		Create(value interface{}) *gorm.DB
		Save(value interface{}) *gorm.DB
		Where(query interface{}, args ...interface{}) *gorm.DB
		Find(dest interface{}, conds ...interface{}) *gorm.DB
		Delete(value interface{}, conds ...interface{}) *gorm.DB
	}

	// Database - Implementação das operações com o banco de dados
	Database struct {
		db IDatabaseFramework
	}
)

var (
	instance Database
	//once     sync.Once
)

func init() {
	db, err := gorm.Open(sqlite.Open("conductor.db"), &gorm.Config{})
	fmt.Println("Conexão com o banco criada")

	if err != nil {
		panic("failed to connect database")
	}

	instance = Database{db: db}
}

func GetDatabase() *Database {
	/*once.Do(func() {
		load()
	})
	*/
	return &instance
}

// Insert - Inserção de registro no banco de dados
func (database *Database) Insert(entity interface{}) {
	database.db.Create(entity)
}

// Update - Alteração de registro no banco de dados
func (database *Database) Update(entity interface{}) {
	database.db.Save(entity)
}

// Find - Consulta de registros no banco de dados
func (database *Database) Find(target interface{}, entity interface{}, inputs interface{}) {
	database.db.Where(entity, inputs).Find(target)
}

// Delete - Exclusão de registro no banco de dados
func (database *Database) Delete(entity interface{}, id ...interface{}) {
	database.db.Delete(entity, id)
}
