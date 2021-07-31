package utils

var AppParams *AppConfig

type AppConfig struct {
	HostAddr  string
	Port      string
	StoreName string
	StoreAddr string
	SecretKey string
	ApiBase  string
}

