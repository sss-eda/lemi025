package serial

// Client TODO
type Client struct{}

// NewClient TODO
func NewClient() (*Client, error) {
	return &Client{}, nil
}

// ReadConfig TODO
func (client *Client) ReadConfig(
	id string,
) error {
	return nil
}
