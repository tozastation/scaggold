package dir

import (
	"fmt"
	"testing"
)

func TestGetParentDirName(t *testing.T) {
	actual, err := GetParentDirName()
	if err != nil {
		t.Error(err)
	}
	expect := "dir"
	if actual != expect {
		fmt.Printf("actual:%v\n,expect:%v\n",actual,expect)
		return
	}
}