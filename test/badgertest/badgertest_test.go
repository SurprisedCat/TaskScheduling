package badgertest

import (
	"fmt"
	"testing"
)

func Test_OpenDB(t *testing.T) {
	err := OpenDB("")
	if err != nil {
		t.Error("OpenDB failed")
	}
	fmt.Println("OpenDB test")

}
