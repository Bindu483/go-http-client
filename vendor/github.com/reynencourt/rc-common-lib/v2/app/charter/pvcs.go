package charter

import (
	"github.com/pkg/errors"
	appsV1 "k8s.io/api/apps/v1"
	"k8s.io/api/apps/v1beta1"
	"k8s.io/api/apps/v1beta2"
	v1 "k8s.io/api/core/v1"

	"github.com/reynencourt/rc-common-lib/v2/proto/app"
)

// PVCs parses all the chart PVCs and creates a map with them
func (c *RcHelmChart) PVCs() ([]*app.AppPVC, error) {
	solPvcs := make([]*app.AppPVC, 0)

	// Parse chart objects
	for filePath, objects := range c.DeploymentManifest.Chart {
		for _, object := range objects {
			switch object.(type) {
			case *v1beta1.StatefulSet:
				statefulSet, ok := object.(*v1beta1.StatefulSet)
				if !ok {
					return nil, errors.New("failed to convert app statefulSet")
				}

				for pvcIndex, pvc := range statefulSet.Spec.VolumeClaimTemplates {
					// Parse access modes
					accessModes := make([]int64, 0)
					for _, accessMode := range pvc.Spec.AccessModes {
						accessModes = append(accessModes, convertKubernetesAccessMode(string(accessMode)))
					}

					// Parse storage class name
					storageClassName := ""
					if pvc.Spec.StorageClassName != nil {
						storageClassName = *pvc.Spec.StorageClassName
					}

					// Add the PVC into the list
					solPvcs = append(solPvcs, &app.AppPVC{
						Name:         GeneratePVCUniqueName(filePath, pvcIndex),
						AccessModes:  accessModes,
						StorageClass: storageClassName,
					})
				}
			case *v1beta2.StatefulSet:
				statefulSet, ok := object.(*v1beta2.StatefulSet)
				if !ok {
					return nil, errors.New("failed to convert app statefulSet")
				}

				for pvcIndex, pvc := range statefulSet.Spec.VolumeClaimTemplates {
					// Parse access modes
					accessModes := make([]int64, 0)
					for _, accessMode := range pvc.Spec.AccessModes {
						accessModes = append(accessModes, convertKubernetesAccessMode(string(accessMode)))
					}

					// Parse storage class name
					storageClassName := ""
					if pvc.Spec.StorageClassName != nil {
						storageClassName = *pvc.Spec.StorageClassName
					}

					// Add the PVC into the list
					solPvcs = append(solPvcs, &app.AppPVC{
						Name:         GeneratePVCUniqueName(filePath, pvcIndex),
						AccessModes:  accessModes,
						StorageClass: storageClassName,
					})
				}
			case *v1.PersistentVolumeClaim:
				pvc, ok := object.(*v1.PersistentVolumeClaim)
				if !ok {
					return nil, errors.New("failed to convert app pvc")
				}

				// Parse access modes
				accessModes := make([]int64, 0)
				for _, accessMode := range pvc.Spec.AccessModes {
					accessModes = append(accessModes, convertKubernetesAccessMode(string(accessMode)))
				}

				// Parse storage class name
				storageClassName := ""
				if pvc.Spec.StorageClassName != nil {
					storageClassName = *pvc.Spec.StorageClassName
				}

				// Add the PVC into the list
				solPvcs = append(solPvcs, &app.AppPVC{
					Name:         GeneratePVCUniqueName(filePath, 0), // Static PVC
					AccessModes:  accessModes,
					StorageClass: storageClassName,
				})

			case *appsV1.StatefulSet:
				statefulSet, ok := object.(*appsV1.StatefulSet)
				if !ok {
					return nil, errors.New("failed to convert app statefulSet")
				}

				for pvcIndex, pvc := range statefulSet.Spec.VolumeClaimTemplates {
					// Parse access modes
					accessModes := make([]int64, 0)
					for _, accessMode := range pvc.Spec.AccessModes {
						accessModes = append(accessModes, convertKubernetesAccessMode(string(accessMode)))
					}

					// Parse storage class name
					storageClassName := ""
					if pvc.Spec.StorageClassName != nil {
						storageClassName = *pvc.Spec.StorageClassName
					}

					// Add the PVC into the list
					solPvcs = append(solPvcs, &app.AppPVC{
						Name:         GeneratePVCUniqueName(filePath, pvcIndex),
						AccessModes:  accessModes,
						StorageClass: storageClassName,
					})
				}
			}
		}
	}

	// Create the solution pvcs object
	return solPvcs, nil
}

func convertKubernetesAccessMode(accessMode string) int64 {
	switch accessMode {
	case "ReadWriteOnce":
		return 0
	case "ReadWriteMany":
		return 1
	case "ReadOnlyMany":
		return 2
	}

	return -1
}
