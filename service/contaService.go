package service

import (
	"log"

	"github.com/VictorNapoles/teste-victor-conductor/adapter/crud"
	"github.com/VictorNapoles/teste-victor-conductor/domain/conta"
)

type (
	ContaService struct {
		crud.CrudService
	}
)

var contaService *ContaService

func init() {
	contaService = &ContaService{}

	log.Println("Serviço de conta criado com sucesso")
}
func GetContaService() *ContaService {
	return contaService
}

func (c *ContaService) FindById(id string) *conta.Conta {
	var result conta.Conta
	c.CrudService.FindById(&result, id)
	return &result
}

func (c *ContaService) FindAll() *[]conta.Conta {
	var contas []conta.Conta
	c.CrudService.FindAll(&contas)
	return &contas
}
