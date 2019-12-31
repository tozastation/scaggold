package parse

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestParseDockerfileDev(t *testing.T) {
	actual, err := DockerfileDev("test", "test")
	if err != nil {
		t.Error(err)
	}
	//
	f, err := os.Open("./expect/Dockerfile")
	if err != nil{
		t.Error(err)
	}
	defer f.Close()
	// 一気に全部読み取り
	b, err := ioutil.ReadAll(f)
	if actual != string(b) {
		fmt.Printf("actual:%v\n,expect:%v\n",actual,string(b))
		return
	}
}