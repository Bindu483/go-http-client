package resourceManager

import (
	"fmt"
	"strings"
)

func (x *Storage) AccessModes() []StorageAccessMode {
	switch x.Type {
	case StorageType_NFS:
		return []StorageAccessMode{StorageAccessMode_RWO, StorageAccessMode_RWX}
	case StorageType_RBD:
		return []StorageAccessMode{StorageAccessMode_RWO}
	case StorageType_CEPHFS:
		return []StorageAccessMode{StorageAccessMode_RWX}
	}

	// TODO: Check default return
	return []StorageAccessMode{}
}

// ProvisionerName generate a unique provisioner name
func (x *Storage) ProvisionerName() string {
	provisionerName := strings.Split(x.Id, "-")[0]
	provisionerName = fmt.Sprintf("%s-provisioner-%s", strings.ToLower(x.Type.String()), provisionerName)

	return provisionerName
}
