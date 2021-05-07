package shared_state

import (
	"context"
	"encoding/json"

	"github.com/reynencourt/rc-common-lib/v2/etcd"
	"github.com/reynencourt/rc-common-lib/v2/state"
)

func (c *Client) DeleteDNSEntry(ctx context.Context) error {
	_, err := c.Delete(ctx, "/host-entries")
	return err
}

func (c *Client) GetAllHostEntries(ctx context.Context) ([]HostEntry, error) {
	var allHostEntries []HostEntry
	resp, err := c.Get(ctx, KeyHostEntries)
	if err != nil {
		if err == etcd.ErrNoKeysFound {
			return allHostEntries, nil
		}
		return allHostEntries, err
	}
	err = json.Unmarshal(resp.Data, &allHostEntries)
	if err != nil {
		return allHostEntries, err
	}
	return allHostEntries, nil
}

func (c *Client) RemoveDNSEntry(ctx context.Context, serviceNames ...string) error {
	hostEntries, err := c.GetAllHostEntries(ctx)
	if err != nil {
		return err
	}
	var leftOut []HostEntry
	for _, hostEntry := range hostEntries {
		if notIn(serviceNames, hostEntry.HostName) {
			leftOut = append(leftOut, hostEntry)
		}
	}
	entries, err := json.Marshal(&leftOut)
	if err != nil {
		return err
	}
	e := &state.Entity{
		Key:  KeyHostEntries,
		Data: entries,
	}
	err = c.Put(ctx, e)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) AddDNSEntry(ctx context.Context, entry ...HostEntry) error {
	hostEntries, err := c.GetAllHostEntries(ctx)
	if err != nil {
		return err
	}
	finalEntries := dedup(append(hostEntries, entry...))
	entries, err := json.Marshal(&finalEntries)
	if err != nil {
		return err
	}
	e := &state.Entity{
		Key:  KeyHostEntries,
		Data: entries,
	}
	err = c.Put(ctx, e)
	if err != nil {
		return err
	}
	return nil
}
