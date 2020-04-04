package config

// "github.com/ralali/sdk-go-api/libs"

// Configs ...
type Configs struct {
	Environtment string
	Service      string
}

// NewConfig ...
func NewConfig() *Configs {
	return &Configs{}
}
