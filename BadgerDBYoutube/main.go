package main

import (
	"fmt"
	"log"
	"os"

	badger "github.com/dgraph-io/badger/v3"
)

func InsertRecords(db *badger.DB, key string, value string) {
	err := db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), []byte(value))
		return err
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Record Inserted")
	}
}
func DisplayRecords(db *badger.DB) {
	err := db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(v []byte) error {
				fmt.Printf("key=%s, value=%s\n", k, v)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

}
func DeleteRecords(db *badger.DB, key string) {
	err := db.Update(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			if string(k) == key {
				err := txn.Delete(k)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Record Deleted =", key)
				}
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	// Open the Badger database located in the /tmp/badger directory.
	// It will be created if it doesn't exist.
	db, err := badger.Open(badger.DefaultOptions("tmp/badger"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Your code here…
	for {
		var choice string
		fmt.Println("Enter your choice")
		fmt.Println("1. Enter value in BadgerDB")
		fmt.Println("2. Display all records")
		fmt.Println("3. Delete records ")
		fmt.Println("4. Exit")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			var key, value string
			fmt.Println("Enter key")
			fmt.Scanln(&key)
			fmt.Println("Enter value")
			fmt.Scanln(&value)
			InsertRecords(db, key, value)
		case "2":
			DisplayRecords(db)
		case "3":
			var key string
			fmt.Println("Enter key")
			fmt.Scanln(&key)
			DeleteRecords(db, key)
		default:
			os.Exit(0)
		}
	}

}
