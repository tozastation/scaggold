package layout

import (
	"fmt"
	"os"
)

var (
	Directories = []string{"build","cmd","configs","deployments","docs","internal","pkg","scripts","tests"}
	ServiceDirectories = []string{"build","cmd","configs","internal"}
)

func GenerateGolangStandardsProjectLayout() error {
	for _, directory := range Directories {
		if err := os.Mkdir(directory, 0777); err != nil {
			return fmt.Errorf("generate golang-standards/project-layout error: %w",err)
		}
	}
	return nil
}

func GenerateService(name string) error {
	for _, directory := range ServiceDirectories {
		combinedPath := directory + "/" + name
		if err := os.MkdirAll(combinedPath, 0777); err != nil {
			return fmt.Errorf("generate golang-standards/project-layout error: %w",err)
		}
	}
	return nil
}