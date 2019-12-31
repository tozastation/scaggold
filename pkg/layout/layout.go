package layout

import (
	"fmt"
	"github.com/tozastation/sgoffold/pkg/dir"
	"github.com/tozastation/sgoffold/pkg/file"
	"github.com/tozastation/sgoffold/pkg/parse"
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
	if err := _generateDockerfile(name); err != nil {
		return err
	}
	return nil
}

func _generateDockerfile(serviceName string) error {
	projectName, _ := dir.GetParentDirName()
	buildPath := "build/" + serviceName
	// Dockerfile
	parsedDockerfileDev, err := parse.DockerfileDev(projectName, serviceName)
	if err != nil {
		return err
	}
	if err := file.CreateFile(buildPath + "/Dockerfile.dev", parsedDockerfileDev); err != nil {
		return err
	}
	// Dockerfile
	parsedDockerfileProd, err := parse.DockerfileProd(projectName, serviceName)
	if err != nil {
		return err
	}
	if err := file.CreateFile(buildPath + "/Dockerfile.prod", parsedDockerfileProd); err != nil {
		return err
	}
	return nil
}