package crud

import (
	"testing"

	"github.com/google/uuid"
)

type (
	CrudRepositoryStub struct {
		doInsert   func(entity interface{})
		doUpdate   func(entity interface{})
		doFind     func(target interface{}, entity interface{}, inputs ...interface{})
		doFindAll  func(target interface{})
		doDelete   func(entity interface{}, id ...interface{})
		doFindById func(target interface{}, conds ...interface{})
	}
)

// Insert - Inserção de registro no banco de dados
func (repository *CrudRepositoryStub) Insert(entity interface{}) {
	repository.doInsert(entity)
}

// Update - Alteração de registro no banco de dados
func (repository *CrudRepositoryStub) Update(entity interface{}) {
	repository.doUpdate(entity)
}

// Find - Consulta de registros no banco de dados
func (repository *CrudRepositoryStub) Find(target interface{}, entity interface{}, inputs ...interface{}) {
	repository.doFind(target, entity, inputs)
}

// FindAll - Consulta de registros no banco de dados
func (repository *CrudRepositoryStub) FindAll(target interface{}) {
	repository.doFindAll(target)
}

// FindById - Consulta de registros no banco de dados por id
func (repository *CrudRepositoryStub) FindById(target interface{}, conds ...interface{}) {
	repository.doFindById(target, conds...)
}

// Delete - Exclusão de registro no banco de dados
func (repository *CrudRepositoryStub) Delete(entity interface{}, id ...interface{}) {
	repository.doDelete(entity, id)
}

func TestServiceInsert(t *testing.T) {
	isDatabaseCalled := false
	repository := &CrudRepositoryStub{
		doInsert: func(entity interface{}) {
			isDatabaseCalled = true
			if entity == nil {
				t.Error("Registro passado para o banco de dados é nulo")
			}
		},
	}

	service := &CrudService{repository: repository}

	service.Insert(uuid.New().String())

	if !isDatabaseCalled {
		t.Fatal("Método do framework de orm não foi chamado")
	}
}

func TestServiceUpdate(t *testing.T) {
	isDatabaseCalled := false
	repository := &CrudRepositoryStub{
		doUpdate: func(entity interface{}) {
			isDatabaseCalled = true
			if entity == nil {
				t.Error("Registro passado para o banco de dados é nulo")
			}
		},
	}

	service := &CrudService{repository: repository}

	service.Update(uuid.New().String())

	if !isDatabaseCalled {
		t.Fatal("Método do banco de dados não foi chamado")
	}
}

func TestServiceDelete(t *testing.T) {
	isDatabaseCalled := false
	repository := &CrudRepositoryStub{
		doDelete: func(entity interface{}, id ...interface{}) {
			isDatabaseCalled = true
			if entity == nil || id == nil {
				t.Error("Registro passado para o banco de dados é nulo")
			}
		},
	}

	service := &CrudService{repository: repository}

	service.Delete(uuid.New().String(), 1, 2)

	if !isDatabaseCalled {
		t.Fatal("Método do banco de dados não foi chamado")
	}
}
