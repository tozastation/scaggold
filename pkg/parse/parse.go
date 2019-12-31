package parse

import (
	"github.com/rakyll/statik/fs"
	_ "github.com/tozastation/sgoffold/statik"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	StatikFS http.FileSystem
)

func init() {
	var err error
	StatikFS, err = fs.New()
	if err != nil {
		log.Fatal(err)
	}
}

func DockerfileProd(projectName, serviceName string) (string, error) {
	r, err := StatikFS.Open("/Dockerfile.prod")
	if err != nil {
		return "", err
	}
	defer r.Close()
	contents, err := ioutil.ReadAll(r)
	if err != nil {
		return "", nil
	}
	replaced := strings.Replace(string(contents), "your-project-name", projectName, -1)
	replaced = strings.Replace(replaced, "your-service-name", serviceName, -1)
	replaced = replaced + "\n"
	return replaced, nil
}

func DockerfileDev(projectName, serviceName string) (string, error) {
	r, err := StatikFS.Open("/Dockerfile.dev")
	if err != nil {
		return "", err
	}
	defer r.Close()
	contents, err := ioutil.ReadAll(r)
	if err != nil {
		return "", nil
	}
	replaced := strings.Replace(string(contents), "your-project-name", projectName, -1)
	replaced = strings.Replace(replaced, "your-service-name", serviceName, -1)
	replaced = replaced + "\n"
	return replaced, nil
}