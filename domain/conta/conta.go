package conta

type (
	Conta struct {
		ID     string `gorm:"primaryKey"`
		Status string
	}
)
