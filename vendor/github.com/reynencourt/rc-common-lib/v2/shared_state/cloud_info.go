package shared_state

import "context"

func (c *Client) GetCloudInfo(ctx context.Context) ([]byte, error) {
	resp, err := c.Get(ctx, KeyCloudInfo)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
