package entities

type Token struct {
	DefaultModel
	Name            string `json:"name" gorm:"size:128"`
	Symbol          string `json:"symbol" gorm:"size:16"`
	Decimals        int64  `json:"decimals"`
	NetworkName     string `json:"network_name" gorm:"size:32"`
	ContractAddress string `json:"contract_address" gorm:"size:512;uniqueIndex"`
}
