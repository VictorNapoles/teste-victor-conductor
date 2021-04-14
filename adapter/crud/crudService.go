package crud

import "log"

type (

	// IDatabase - Interface de operações com o banco de dados
	ICrudRepository interface {
		Insert(entity interface{})
		Update(entity interface{})
		Find(target interface{}, entity interface{}, inputs ...interface{})
		FindAll(target interface{})
		FindById(target interface{}, conds ...interface{})
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
	log.Println("Serviço criado com sucesso")
}

func GetService() *CrudService {
	return service
}

func (s *CrudService) Repository() ICrudRepository {
	if s.repository == nil {
		s.repository = GetRepository()
	}
	return s.repository
}

// Insert - Inserção de registro no banco de dados
func (c *CrudService) Insert(entity interface{}) {
	c.Repository().Insert(entity)
}

// Update - Alteração de registro no banco de dados
func (c *CrudService) Update(entity interface{}) {
	c.Repository().Update(entity)
}

// Find - Consulta de registros no banco de dados
func (c *CrudService) Find(target interface{}, entity interface{}, inputs ...interface{}) {
	c.Repository().Find(target, entity, inputs)
}

// FindAll - Consulta todos os registros no banco de dados
func (c *CrudService) FindAll(target interface{}) {
	c.Repository().FindAll(target)
}

// FindById - Consulta de registros no banco de dados por id
func (c *CrudService) FindById(target interface{}, conds ...interface{}) {
	c.Repository().FindById(target, conds...)
}

// Delete - Exclusão de registro no banco de dados
func (c *CrudService) Delete(entity interface{}, id ...interface{}) {
	c.Repository().Delete(entity, id)
}
