package controller

import (
	"fmt"
	"log"

	"github.com/VictorNapoles/teste-victor-conductor/config/http"
	"github.com/VictorNapoles/teste-victor-conductor/domain/conta"
	"github.com/VictorNapoles/teste-victor-conductor/domain/transacao"
	"github.com/VictorNapoles/teste-victor-conductor/service"
	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
)

type (
	IContaService interface {
		FindById(id string) *conta.Conta
		FindAll() *[]conta.Conta
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

func (c *ContaController) GetById(context *gin.Context) {
	context.JSON(200, c.service.FindById(context.Param("id")))
}

func (c *ContaController) GetAll(context *gin.Context) {
	context.JSON(200, c.service.FindAll())
}

func (c *ContaController) GetTransacoesByConta(context *gin.Context) {
	context.JSON(200, c.transacaoService.FindByConta(context.Param("id")))
}

func (c *ContaController) GetTransacoesByContaReport(context *gin.Context) {
	w := context.Writer
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%q", "report.pdf"))
	c.transacaoService.GenerateReport(context.Param("id")).Output(context.Writer)

}
