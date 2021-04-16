package controller

import (
	"fmt"
	"log"

	"github.com/VictorNapoles/teste-victor-conductor/config/http"
	"github.com/VictorNapoles/teste-victor-conductor/controller/response"
	"github.com/VictorNapoles/teste-victor-conductor/domain/conta"
	"github.com/VictorNapoles/teste-victor-conductor/domain/transacao"
	"github.com/VictorNapoles/teste-victor-conductor/service"
	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
)

type (
	IContaService interface {
		FindById(id string) conta.Conta
		FindAll() []conta.Conta
	}

	IContaTransacaoService interface {
		FindByConta(idConta string) []transacao.Transacao
		GenerateReport(idConta string) *gofpdf.Fpdf
	}

	ContaController struct {
		service          IContaService
		transacaoService IContaTransacaoService
	}
)

var (
	contaController *ContaController
)

func init() {

	contaController = &ContaController{
		service:          service.GetContaService(),
		transacaoService: service.GetTransacaoService(),
	}

	http.GetRouter().GET("/contas", contaController.GetAll)
	http.GetRouter().GET("/contas/:id", contaController.GetById)
	http.GetRouter().GET("/contas/:id/transacoes", contaController.GetTransacoesByConta)
	http.GetRouter().GET("/contas/:id/transacoes.pdf", contaController.GetTransacoesByContaReport)

	log.Println("Controller de conta criado com sucesso")
}

func GetContaController() *ContaController {
	return contaController
}

// @Summary Obter contas por id
// @Description Retorna a lista com todas as contas
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "Id da conta"
// @Success 200 {object} []response.ContaResponse"
// @Router /contas/{id} [get]
func (c *ContaController) GetById(context *gin.Context) {
	conta := c.service.FindById(context.Param("id"))

	response := &response.ContaResponse{ID: conta.ID, Status: conta.Status}
	context.JSON(200, response)
}

// @Summary Obter contas
// @Description Retorna a lista com todas as contas
// @Accept  json
// @Produce  json
// @Success 200 {object} []response.ContaResponse"
// @Router /contas [get]
func (c *ContaController) GetAll(context *gin.Context) {
	contas := c.service.FindAll()

	result := make([]response.ContaResponse, 0)

	for _, element := range contas {
		result = append(result, response.ContaResponse{ID: element.ID, Status: element.Status})
	}
	context.JSON(200, result)
}

// @Summary Obter transações da conta
// @Description Retorna a lista com todas as transações da conta
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "Id da conta"
// @Success 200 {object} []response.TransacaoResponse"
// @Router /contas/{id}/transacoes [get]
func (c *ContaController) GetTransacoesByConta(context *gin.Context) {
	transacoes := c.transacaoService.FindByConta(context.Param("id"))
	result := make([]response.TransacaoResponse, 0)

	for _, element := range transacoes {
		result = append(result, response.TransacaoResponse{ID: element.ID, ContaID: element.ContaID, Descricao: element.Descricao, Valor: element.Valor})
	}
	context.JSON(200, result)
}

// @Summary Obter contas
// @Description Download de relatório com todas as transações da conta
// @Accept  json
// @Produce  application/pdf
// @Param   id     path    string     true        "Id da conta"
// @Success 200 {object} byte"
// @Router /contas/{id}/transacoes.pdf [get]
func (c *ContaController) GetTransacoesByContaReport(context *gin.Context) {
	w := context.Writer
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%q", "report.pdf"))
	c.transacaoService.GenerateReport(context.Param("id")).Output(context.Writer)

}
