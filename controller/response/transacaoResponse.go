package response

type (
	TransacaoResponse struct {
		ID        string  `json:"id"`
		ContaID   string  `json:"id_conta"`
		Descricao string  `json:"descricao"`
		Valor     float64 `json:"valor"`
	}
)
