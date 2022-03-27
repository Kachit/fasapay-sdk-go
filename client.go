package fasapay

import "net/http"

//Client struct
type Client struct {
	transport *Transport
	config    *Config
}

//NewClientFromConfig Create new client from config
func NewClientFromConfig(config *Config, cl *http.Client) (*Client, error) {
	err := config.IsValid()
	if err != nil {
		return nil, err
	}
	if cl == nil {
		cl = &http.Client{}
	}
	transport := NewHttpTransport(config, cl)
	return &Client{transport, config}, nil
}

//Accounts resource method
func (c *Client) Accounts() *AccountsResource {
	return &AccountsResource{ResourceAbstract: NewResourceAbstract(c.transport, c.config)}
}

//Transfers resource method
func (c *Client) Transfers() *TransfersResource {
	return &TransfersResource{ResourceAbstract: NewResourceAbstract(c.transport, c.config)}
}
