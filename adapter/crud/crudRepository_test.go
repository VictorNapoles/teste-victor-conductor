package crud

import (
	"testing"

	"github.com/google/uuid"
)

type (
	DatabaseStub struct {
		doInsert  func(entity interface{})
		doUpdate  func(entity interface{})
		doFind    func(target interface{}, entity interface{}, inputs ...interface{})
		doFindAll func(target interface{})
		doDelete  func(entity interface{}, id ...interface{})
	}
)

// Insert - Inserção de registro no banco de dados
func (database *DatabaseStub) Insert(entity interface{}) {
	database.doInsert(entity)
}

// Update - Alteração de registro no banco de dados
func (database *DatabaseStub) Update(entity interface{}) {
	database.doUpdate(entity)
}

// Find - Consulta de registros no banco de dados
func (database *DatabaseStub) Find(target interface{}, entity interface{}, inputs ...interface{}) {
	database.doFind(target, entity, inputs)
}

// FindAll - Consulta de registros no banco de dados
func (database *DatabaseStub) FindAll(target interface{}) {
	database.doFindAll(target)
}

// Delete - Exclusão de registro no banco de dados
func (database *DatabaseStub) Delete(entity interface{}, id ...interface{}) {
	database.doDelete(entity, id)
}

func TestInsert(t *testing.T) {
	isDatabaseCalled := false
	db := &DatabaseStub{
		doInsert: func(entity interface{}) {
			isDatabaseCalled = true
			if entity == nil {
				t.Error("Registro passado para o banco de dados é nulo")
			}
		},
	}

	repository := &CrudRepository{db: db}

	repository.Insert(uuid.New().String())

	if !isDatabaseCalled {
		t.Fatal("Método do framework de orm não foi chamado")
	}
}

func TestUpdate(t *testing.T) {
	isDatabaseCalled := false
	db := &DatabaseStub{
		doUpdate: func(entity interface{}) {
			isDatabaseCalled = true
			if entity == nil {
				t.Error("Registro passado para o banco de dados é nulo")
			}
		},
	}

	repository := &CrudRepository{db: db}

	repository.Update(uuid.New().String())

	if !isDatabaseCalled {
		t.Fatal("Método do banco de dados não foi chamado")
	}
}

func TestDelete(t *testing.T) {
	isDatabaseCalled := false
	db := &DatabaseStub{
		doDelete: func(entity interface{}, id ...interface{}) {
			isDatabaseCalled = true
			if entity == nil || id == nil {
				t.Error("Registro passado para o banco de dados é nulo")
			}
		},
	}

	repository := &CrudRepository{db: db}

	repository.Delete(uuid.New().String(), 1, 2)

	if !isDatabaseCalled {
		t.Fatal("Método do banco de dados não foi chamado")
	}
}
