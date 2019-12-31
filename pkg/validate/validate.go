package validate

import (
	"github.com/tozastation/sgoffold/pkg/layout"
	"os"
)

func IsArgumentExist(args []string) bool {
	if len(args) == 0 {
		return false
	}
	return true
}

func IsGolangStandardsProjectLayoutDirExist() bool {
	for _, directory := range layout.Directories {
		_, err := os.Stat(directory)
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}