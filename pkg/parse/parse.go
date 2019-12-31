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

func DockerfileDev(projectName, serviceName string) (string, error) {
	r, err := StatikFS.Open("/Dockerfile.dev")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	contents, err := ioutil.ReadAll(r)
	if err != nil {
		return "", nil
	}
	// 全部のJavaを置き換え
	replaced := strings.Replace(string(contents), "your-project-name", projectName, -1)
	replaced = strings.Replace(replaced, "your-service-name", serviceName, -1)
	return replaced, nil
}