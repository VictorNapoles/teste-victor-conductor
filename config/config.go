package config

import (
	"github.com/VictorNapoles/teste-victor-conductor/adapter/crud"
	"github.com/VictorNapoles/teste-victor-conductor/config/database"
	"github.com/VictorNapoles/teste-victor-conductor/domain/conta"
	"github.com/VictorNapoles/teste-victor-conductor/domain/transacao"
	"github.com/google/uuid"
)

func init() {
	database.GetDatabase().GetORM().AutoMigrate(&conta.Conta{}, &transacao.Transacao{})
	service := crud.GetService()
	for i := 0; i < 15; i++ {

		id := uuid.New().String()
		descricoes := [3]string{"Apple Store", "Amazon Prime", "Spotify"}
		prices := [3]float64{30.00, 9.90, 35.50}

		service.Insert(&conta.Conta{ID: id, Status: "ATIVA"})

		for j := 0; j < len(descricoes); j++ {
			service.Insert(&transacao.Transacao{ID: uuid.New().String(), Descricao: descricoes[j], Valor: prices[j], ContaID: id})
		}
	}

}
