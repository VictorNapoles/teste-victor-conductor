package crud

import "fmt"

type (

	// IDatabase - Interface de operações com o banco de dados
	ICrudRepository interface {
		Insert(entity interface{})
		Update(entity interface{})
		Find(target interface{}, entity interface{}, inputs ...interface{})
		FindAll(target interface{})
		Delete(entity interface{}, id ...interface{})
	}

	CrudService struct {
		repository ICrudRepository
	}
)

var (
	service *CrudService
)

func init() {
	service = &CrudService{repository: GetRepository()}
	fmt.Println("Serviço criado com sucesso")
}

func GetService() *CrudService {
	return service
}

// Insert - Inserção de registro no banco de dados
func (c *CrudService) Insert(entity interface{}) {
	c.repository.Insert(entity)
}

// Update - Alteração de registro no banco de dados
func (c *CrudService) Update(entity interface{}) {
	c.repository.Update(entity)
}

// Find - Consulta de registros no banco de dados
func (c *CrudService) Find(target interface{}, entity interface{}, inputs ...interface{}) {
	c.repository.Find(target, entity, inputs)
}

// FindAll - Consulta todos os registros no banco de dados
func (c *CrudService) FindAll(target interface{}) {
	c.repository.FindAll(target)
}

// Delete - Exclusão de registro no banco de dados
func (database *CrudService) Delete(entity interface{}, id ...interface{}) {
	database.repository.Delete(entity, id)
}
