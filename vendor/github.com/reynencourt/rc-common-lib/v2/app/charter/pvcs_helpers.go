package charter

import (
	"fmt"
)

func GeneratePVCUniqueName(filePath string, pvcIndex int) string {
	return fmt.Sprintf("%s-%d", filePath, pvcIndex)
}
