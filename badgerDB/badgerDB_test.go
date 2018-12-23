package badgerDB

import (
	"fmt"
	"testing"
)

func Test_badger(t *testing.T) {
	err := Set([]byte("cx"), []byte("chenxin"))
	if err != nil {
		t.Error(err)
	}
	value := Get([]byte("cx"))
	if value == nil {
		t.Error("Get error")
	} else {
		fmt.Printf("%s\n", value)
		err = Set([]byte("cx"), []byte("lijuan"))
		value = Get([]byte("cx"))
		fmt.Printf("%s\n", value)
		if err != nil {
			t.Error(err)
		}
		err := Delete([]byte("cx"))
		if err != nil {
			t.Error(err)
		}
	}
	GbCollect()
}
