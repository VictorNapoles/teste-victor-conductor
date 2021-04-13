package database

import (
	"testing"

	"gorm.io/gorm"
)

type (
	DatabaseStub struct {
		doInsert func(entity interface{})
		doUpdate func(entity interface{})
		doDelete func(entity interface{}, id ...interface{})
	}

	GormStub struct {
		*gorm.DB
		doCreate func(value interface{}) *gorm.DB
		doSave   func(value interface{}) *gorm.DB
		doDelete func(value interface{}, conds ...interface{}) *gorm.DB
	}

	Entity struct {
		Name string
	}
)

// ######################################## DatabaseStub ########################################

func (d *DatabaseStub) Insert(entity *interface{}) {
	d.doInsert(entity)
}

func (d *DatabaseStub) Update(entity *interface{}) {
	d.doUpdate(entity)
}

func (d *DatabaseStub) Delete(entity interface{}, id ...interface{}) {
	d.doDelete(entity, id)
}

// ######################################## GormStub ########################################

func (g *GormStub) Create(value interface{}) *gorm.DB {
	return g.doCreate(value)
}

func (g *GormStub) Save(value interface{}) *gorm.DB {
	return g.doSave(value)
}

func (g *GormStub) Delete(value interface{}, conds ...interface{}) *gorm.DB {
	return g.doDelete(value, conds)
}

// ####################################################################################################

func TestInsert(t *testing.T) {
	isGormCalled := false

	gormStub := &GormStub{
		doCreate: func(value interface{}) *gorm.DB {
			isGormCalled = true
			if value == nil {
				t.Error("Registro passado para o banco de dados é nulo")
			}
			return &gorm.DB{}
		},
	}

	database := &Database{orm: gormStub}
	entity := &Entity{Name: "Entity Name"}
	database.Insert(entity)

	if !isGormCalled {
		t.Error("Método do framework de orm não foi chamado")
	}
}

func TestUpdate(t *testing.T) {
	isGormCalled := false

	gormStub := &GormStub{
		doSave: func(value interface{}) *gorm.DB {
			if value == nil {
				t.Error("Registro passado para o banco de dados é nulo")
			}
			isGormCalled = true
			return &gorm.DB{}
		},
	}

	database := &Database{orm: gormStub}

	database.Update(&Entity{Name: "Entity Name"})

	if !isGormCalled {
		t.Error("Método do framework de orm não foi chamado")
	}
}

func TestDelete(t *testing.T) {
	isGormCalled := false

	gormStub := &GormStub{
		doDelete: func(value interface{}, conds ...interface{}) *gorm.DB {
			isGormCalled = true
			if value == nil || conds == nil {
				t.Error("Registro passado para o banco de dados é nulo")
			}
			return &gorm.DB{}
		},
	}

	database := &Database{orm: gormStub}

	database.Delete(&Entity{}, 1, 2, 3)

	if !isGormCalled {
		t.Fatal("Método do framework de orm não foi chamado")
	}
}
