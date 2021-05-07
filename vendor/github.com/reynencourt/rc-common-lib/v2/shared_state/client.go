package shared_state

import "github.com/reynencourt/rc-common-lib/v2/state"

type Client struct {
	state.State
}

func New(s state.State) *Client {
	return &Client{s}
}

func (c *Client) GetState() state.State {
	return c.State
}
