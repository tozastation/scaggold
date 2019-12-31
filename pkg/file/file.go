package file

import (
	"io/ioutil"
	"os"
)

func CreateFile(filePath, body string) error {
	fp, err := os.Create(filePath)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(filePath, []byte(body), 0777); err!=nil {
		return err
	}
	defer fp.Close()
	return nil
}
