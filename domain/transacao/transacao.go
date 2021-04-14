package transacao

import "github.com/VictorNapoles/teste-victor-conductor/domain/conta"

type (
	Transacao struct {
		ID        string
		Descricao string
		Valor     float32
		ContaID   string
		Conta     conta.Conta
	}
)
