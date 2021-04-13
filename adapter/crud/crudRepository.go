package crud

import (
	"fmt"

	"github.com/VictorNapoles/teste-victor-conductor/config/database"
)

type (

	// IDatabase - Interface de operações com o banco de dados
	IDatabase interface {
		Insert(entity interface{})
		Update(entity interface{})
		Find(target interface{}, entity interface{}, inputs ...interface{})
		FindAll(target interface{})
		Delete(entity interface{}, id ...interface{})
	}

	CrudRepository struct {
		db IDatabase
	}
)

var (
	repository *CrudRepository
)

func init() {
	repository = &CrudRepository{db: database.GetDatabase()}
	fmt.Println("Repositório criado com sucesso")
}

func GetRepository() *CrudRepository {
	return repository
}

// Insert - Inserção de registro no banco de dados
func (c *CrudRepository) Insert(entity interface{}) {
	c.db.Insert(entity)
}

// Update - Alteração de registro no banco de dados
func (c *CrudRepository) Update(entity interface{}) {
	c.db.Update(entity)
}

// Find - Consulta de registros no banco de dados
func (c *CrudRepository) Find(target interface{}, entity interface{}, inputs ...interface{}) {
	c.db.Find(target, entity, inputs)
}

// FindAll - Consulta todos os registros no banco de dados
func (c *CrudRepository) FindAll(target interface{}) {
	c.db.FindAll(target)
}

// Delete - Exclusão de registro no banco de dados
func (database *CrudRepository) Delete(entity interface{}, id ...interface{}) {
	database.db.Delete(entity, id)
}
