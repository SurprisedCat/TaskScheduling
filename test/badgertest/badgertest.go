package badgertest

import (
	"log"

	"github.com/dgraph-io/badger"
)

//OpenDB Open basger database at specific location
func OpenDB(path string) (err error) {
	if path == "" {
		path = "/tmp/badger"
	}
	// Open the Badger database located in the /tmp/badger directory.
	// It will be created if it doesn't exist.
	opts := badger.DefaultOptions
	opts.Dir = path
	opts.ValueDir = path
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return err
}
