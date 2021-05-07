package controlleretcd

//type ControllerEtcdClient struct {
//	url      []string
//	queue    *recipe.Queue
//	client   *clientv3.Client
//	nodeName string
//}
//
//func New(url []string, nodeName string) (*ControllerEtcdClient, error) {
//
//	controllerETCDClient, err := rc_etcd.New(url)
//	if err != nil {
//		return nil, err
//	}
//
//	return &ControllerEtcdClient{
//		url:      url,
//		client:   controllerETCDClient.GetClient(),
//		nodeName: nodeName,
//	}, nil
//}
//
//func (c *ControllerEtcdClient) UpdateHosts(oldEntries []rc_etcd.HostEntry, newEntries []rc_etcd.HostEntry) error {
//
//	var updatedValues = make(map[string]string, 0)
//	var finalEntries []rc_etcd.HostEntry
//
//	for _, o := range oldEntries {
//		updatedValues[o.HostName] = o.IP
//	}
//
//	for _, n := range newEntries {
//		updatedValues[n.HostName] = n.IP
//	}
//
//	for hostname, ip := range updatedValues {
//		finalEntries = append(finalEntries, rc_etcd.HostEntry{HostName: hostname, IP: ip})
//	}
//
//	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
//	defer cancel()
//
//	entries, err := json.Marshal(&finalEntries)
//	if err != nil {
//		return err
//	}
//
//	_, err = c.client.Put(ctx, "/host-entries", string(entries))
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
