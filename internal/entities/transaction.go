package entities

type Transaction struct {
	DefaultModel
	TxHash string                 `json:"tx_hash" gorm:"type:VARCHAR(255) NOT NULL;uniqueIndex"`
	From   string                 `json:"from" gorm:"type:VARCHAR(255) NOT NULL"`
	To     string                 `json:"to" gorm:"type:VARCHAR(255) NOT NULL"`
	Value  string                 `json:"value" gorm:"type:VARCHAR(255)"`
	Status string                 `json:"status" gorm:"type:VARCHAR(255) NOT NULL"`
	Params map[string]interface{} `json:"params" gorm:"type:JSON"`
}

type TransactionDto struct {
	Transaction
}

func (t TransactionDto) ToEntity() (interface{}, error) {
	return &t.Transaction, nil
}

func (t TransactionDto) IsValid() bool {
	return true
}

func NewTxDtO(t Transaction) IEntity {
	return &TransactionDto{
		Transaction: t,
	}
}
