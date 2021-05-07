package shared_state

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/reynencourt/rc-common-lib/v2/proto/resource"
	"github.com/reynencourt/rc-common-lib/v2/state"
)

func (c *Client) PutClusterResourceInfo(ctx context.Context, info *resource.ClusterResourceInfo) error {
	k := fmt.Sprintf(KeyClusterResourceInfo, info.ClusterId)
	data, err := json.Marshal(info)
	if err != nil {
		return err
	}
	err = c.Put(ctx, &state.Entity{
		Key:  k,
		Data: data,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GetClusterResourceInfo(ctx context.Context, clusterId string) (*resource.ClusterResourceInfo, error) {
	k := fmt.Sprintf(KeyClusterResourceInfo, clusterId)
	resp, err := c.Get(ctx, k)
	if err != nil {
		return nil, err
	}

	var info resource.ClusterResourceInfo
	err = json.Unmarshal(resp.Data, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (c *Client) ReserveClusterResource(ctx context.Context, clusterId string, resources *resource.Resources) error {
	clusterResource, err := c.GetClusterResourceInfo(ctx, clusterId)
	if err != nil {
		return err
	}
	clusterResource.Used.Add(resources)
	err = c.PutClusterResourceInfo(ctx, clusterResource)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) UnReserveClusterResource(ctx context.Context, clusterId string, resources *resource.Resources) error {
	clusterResource, err := c.GetClusterResourceInfo(ctx, clusterId)
	if err != nil {
		return err
	}
	clusterResource.Used.Sub(resources)
	err = c.PutClusterResourceInfo(ctx, clusterResource)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) ListClusterResource(ctx context.Context) ([]*resource.ClusterResourceInfo, error) {
	var resourceList []*resource.ClusterResourceInfo

	resp, err := c.List(ctx, KeyClusterResourceInfoPrefix)
	if err != nil {
		return nil, err
	}

	for _, entity := range resp {
		if len(strings.Split(entity.Key, "/")) != 5 {
			return nil, err
		}
		var cr resource.ClusterResourceInfo
		err = json.Unmarshal(entity.Data, &cr)
		if err != nil {
			return nil, err
		}
		resourceList = append(resourceList, &cr)
	}
	return resourceList, nil
}

func (c *Client) DeleteResourceInfoForCluster(ctx context.Context, clusterId string) error {
	k := fmt.Sprintf(KeyClusterResourceInfo, clusterId)
	_, err := c.Delete(ctx, k)
	if err != nil {
		return err
	}
	return nil
}
