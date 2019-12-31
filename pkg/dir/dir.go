package dir

import (
	"os"
	"path/filepath"
)

func GetParentDirName() (string, error) {
	wd,err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Base(wd), nil
}
