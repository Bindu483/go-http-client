package controlleretcd

import (
	"context"
	"encoding/json"
	"github.com/reynencourt/rc-common-lib/v2/shared_state"
	"github.com/reynencourt/rc-common-lib/v2/state"
	"time"
)

type EtcdClientV2 struct {
	url      []string
	state    state.State
	nodeName string
}

func NewClientV2(url []string, nodeName string, s state.State) *EtcdClientV2 {
	return &EtcdClientV2{state: s, nodeName: nodeName, url: url}
}
func (c EtcdClientV2) GetState() state.State {
	return c.state
}

func (c *EtcdClientV2) UpdateHosts(oldEntries []shared_state.HostEntry, newEntries []shared_state.HostEntry) error {

	var updatedValues = make(map[string]string, 0)
	var finalEntries []shared_state.HostEntry

	for _, o := range oldEntries {
		updatedValues[o.HostName] = o.IP
	}

	for _, n := range newEntries {
		updatedValues[n.HostName] = n.IP
	}

	for hostname, ip := range updatedValues {
		finalEntries = append(finalEntries, shared_state.HostEntry{HostName: hostname, IP: ip})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	entries, err := json.Marshal(&finalEntries)
	if err != nil {
		return err
	}

	e := &state.Entity{
		Key:  shared_state.KeyHostEntries,
		Data: entries,
	}
	err = c.state.Put(ctx, e)
	if err != nil {
		return err
	}

	return nil
}
