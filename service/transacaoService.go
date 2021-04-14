package service

import (
	"log"

	"github.com/VictorNapoles/teste-victor-conductor/adapter/crud"
	"github.com/VictorNapoles/teste-victor-conductor/domain/transacao"
	"github.com/jung-kurt/gofpdf"
)

type (
	TransacaoService struct {
		crud.CrudService
	}
)

var transacaoService *TransacaoService

func init() {
	transacaoService = &TransacaoService{}

	log.Println("Serviço de transação criado com sucesso")
}
func GetTransacaoService() *TransacaoService {
	return transacaoService
}

func (t *TransacaoService) FindByConta(idConta string) []transacao.Transacao {
	var transacoes []transacao.Transacao
	t.CrudService.Find(&transacoes, &transacao.Transacao{ContaID: idConta}, "conta_id")
	return transacoes
}

func (t *TransacaoService) GenerateReport(idConta string) *gofpdf.Fpdf {
	//transacoes := t.FindByConta(idConta)

	xCellConta, xCellTransacao, xCellValor, yTitle, yHeader := 10, 90, 175, 50, 70
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.ImageOptions(
		"assert/logo_conductor.png",
		80, 20,
		60, 0,
		false,
		gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true},
		0,
		"",
	)
	pdf.Text(float64(80), float64(yTitle), "RELATÓRIO")

	pdf.Text(float64(xCellConta), float64(yHeader), "CONTA")
	pdf.Text(float64(xCellTransacao), float64(yHeader), "TRANSAÇÃO")
	pdf.Text(float64(xCellValor), float64(yHeader), "VALOR")
	//pdf.SetFont()

	return pdf
}
