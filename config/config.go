package config

import (
	"github.com/VictorNapoles/teste-victor-conductor/adapter/crud"
	"github.com/VictorNapoles/teste-victor-conductor/config/database"
	"github.com/VictorNapoles/teste-victor-conductor/domain/conta"
	"github.com/google/uuid"
)

func init() {
	database.GetDatabase().GetORM().AutoMigrate(&conta.Conta{})
	service := crud.GetService()
	for i := 0; i < 15; i++ {
		service.Insert(&conta.Conta{ID: uuid.New().String(), Status: "ATIVA"})
	}

}
