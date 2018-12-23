package badgerDB

import (
	"../config"
	"github.com/dgraph-io/badger"
)

var databasePath string

func init() {
	if config.DataPath != "" {
		databasePath = config.DataPath
	} else {
		databasePath = "./"
	}
}

//OpenDB Open basger database at specific location
func OpenDB() (db *badger.DB, err error) {
	// Open the Badger database located in the /tmp/badger directory.
	// It will be created if it doesn't exist.
	opts := badger.DefaultOptions
	opts.Dir = databasePath
	opts.ValueDir = databasePath
	db, err = badger.Open(opts)
	if err != nil {
		return nil, err
	}
	return db, err
}

//Get get a value according to key from database
func Get(key []byte) (value []byte) {
	//open DB
	db, err := OpenDB()
	if err != nil {
		return nil //func Get
	}
	defer db.Close()
	//get
	getErr := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		value, err = item.ValueCopy(value)
		if err != nil {
			return err
		}
		return nil
	})
	//result of opeartion get, If transaction view has something wrong return nil
	if getErr != nil {
		return nil //func Get
	}
	return value
}

func Set(key, value []byte) error {
	//open DB
	db, err := OpenDB()
	if err != nil {
		return err //func Set
	}
	defer db.Close()
	// set
	err = db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, value)
	})
	//result of opeartion set
	if err != nil {
		return err
	}
	return nil
}

//Delete erase the key
func Delete(key []byte) error {
	//open DB
	db, err := OpenDB()
	if err != nil {
		return err //func Set
	}
	defer db.Close()
	// delete
	err = db.Update(func(txn *badger.Txn) error {
		return txn.Delete(key)
	})
	//result of operation delete
	if err != nil {
		return err
	}
	return nil
}

func GbCollect() {
	db, _ := OpenDB()
	db.RunValueLogGC(0.7)
	defer db.Close()
}
