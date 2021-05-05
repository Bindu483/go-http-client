package charter

import (
	"github.com/pkg/errors"
	appsV1 "k8s.io/api/apps/v1"
	"k8s.io/api/apps/v1beta1"
	"k8s.io/api/apps/v1beta2"
	v1 "k8s.io/api/core/v1"
)

// PVCs parses all the chart PVCs and creates a map with them
func (c *RcHelmChart) ApplyPvcsMapping(pvcsMapping map[string]string) error {
	for filePath, objects := range c.DeploymentManifest.Chart {
		for objectKey, object := range objects {
			switch object.(type) {
			case *v1.PersistentVolumeClaim:
				pvc, ok := object.(*v1.PersistentVolumeClaim)
				if !ok {
					return errors.New("failed to convert app pvc")
				}

				storageClassName, ok := pvcsMapping[GeneratePVCUniqueName(filePath, 0)] // Static PVC
				if !ok {
					// No mapping found for the giving PVC
					continue
				}

				// Update the PVC's storage class name
				pvc.Spec.StorageClassName = &storageClassName
				c.DeploymentManifest.Chart[filePath][objectKey] = pvc
			case *v1beta1.StatefulSet:
				statefulSet, ok := object.(*v1beta1.StatefulSet)
				if !ok {
					return errors.New("failed to convert app statefulSet")
				}

				for pvcIndex, pvc := range statefulSet.Spec.VolumeClaimTemplates {
					storageClassName, ok := pvcsMapping[GeneratePVCUniqueName(filePath, pvcIndex)]
					if !ok {
						// No mapping found for the giving PVC
						continue
					}

					// Update the PVC's storage class name
					pvc.Spec.StorageClassName = &storageClassName
					statefulSet.Spec.VolumeClaimTemplates[pvcIndex] = pvc
				}

				// Update the statefulset with the changed PVC
				c.DeploymentManifest.Chart[filePath][objectKey] = statefulSet
			case *v1beta2.StatefulSet:
				statefulSet, ok := object.(*v1beta2.StatefulSet)
				if !ok {
					return errors.New("failed to convert app statefulSet")
				}

				for pvcIndex, pvc := range statefulSet.Spec.VolumeClaimTemplates {
					storageClassName, ok := pvcsMapping[GeneratePVCUniqueName(filePath, pvcIndex)]
					if !ok {
						// No mapping found for the giving PVC
						continue
					}

					// Update the PVC's storage class name
					pvc.Spec.StorageClassName = &storageClassName
					statefulSet.Spec.VolumeClaimTemplates[pvcIndex] = pvc
				}

				// Update the statefulset with the changed PVC
				c.DeploymentManifest.Chart[filePath][objectKey] = statefulSet
			case *appsV1.StatefulSet:
				statefulSet, ok := object.(*appsV1.StatefulSet)
				if !ok {
					return errors.New("failed to convert app statefulSet")
				}

				for pvcIndex, pvc := range statefulSet.Spec.VolumeClaimTemplates {
					storageClassName, ok := pvcsMapping[GeneratePVCUniqueName(filePath, pvcIndex)]
					if !ok {
						// No mapping found for the giving PVC
						continue
					}

					// Update the PVC's storage class name
					pvc.Spec.StorageClassName = &storageClassName
					statefulSet.Spec.VolumeClaimTemplates[pvcIndex] = pvc
				}

				// Update the statefulset with the changed PVC
				c.DeploymentManifest.Chart[filePath][objectKey] = statefulSet
			}
		}
	}

	return nil
}
