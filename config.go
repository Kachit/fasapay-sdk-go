package fasapay

import "fmt"

//Config structure
type Config struct {
	ApiUri        string `json:"api_uri"`
	ApiKey        string `json:"api_key"`
	ApiSecretWord string `json:"api_secret_word"`
}

//IsSandbox check is sandbox environment
func (c *Config) IsSandbox() bool {
	return c.ApiUri != ProdAPIUrl
}

//IsValid check is valid config parameters
func (c *Config) IsValid() error {
	var err error
	if c.ApiUri == "" {
		err = fmt.Errorf(`parameter "api_uri" is empty`)
	} else if c.ApiKey == "" {
		err = fmt.Errorf(`parameter "api_key" is empty`)
	} else if c.ApiSecretWord == "" {
		err = fmt.Errorf(`parameter "api_secret_word" is empty`)
	}
	return err
}
