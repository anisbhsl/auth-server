package utils

var AppParams *AppConfig

type AppConfig struct {
	HostAddr       string
	Port           string
	PrivateKeyPath string
	PublicKeyPath  string
	SecretKey      string
	ApiBase        string
}
