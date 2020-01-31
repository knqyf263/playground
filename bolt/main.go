package main

import (
	"fmt"
	"log"

	bolt "go.etcd.io/bbolt"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
func run() error {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte("test")); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("test"))
		if err := bucket.Put([]byte("keyA"), nil); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("test"))
		b := bucket.Get([]byte("keyA"))
		if b == nil {
			fmt.Println("nil")
		}
		fmt.Println(b)
		b = bucket.Get([]byte("keyB"))
		if b == nil {
			fmt.Println("nil")
		}
		fmt.Println(b)
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
