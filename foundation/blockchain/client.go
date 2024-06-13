package blockchain

type Client struct {
	netUrl string
}

func NewClient(rawurl string) *Client {
	return &Client{netUrl: rawurl}
}
