package pkg

type AppName string

const (
	// APIAppName Visible name of the application.
	APIAppName = AppName("GTP_Multi-chain")
	// MySqlConnectorAppName Visible name of the application.
	MySqlConnectorAppName = AppName("GTPM_MysqlConnector")
)

func (name AppName) ToString() string {
	return string(name)
}
