package database

import (
	"log"

	"github.com/boltdb/bolt"
)

type BoltStruct struct {
	BoltDB *bolt.DB
}

func (db *BoltStruct) BoltSetup(dbLocation string) {
	var err interface{}
	db.BoltDB, err = bolt.Open(dbLocation, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	db.boltSetupBucket("ec2")
}

func (db *BoltStruct) boltSetupBucket(bucket string) {
	db.BoltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			log.Fatalf("Error creating bucket %s - %s", bucket, err)
			return err
		}
		return nil
	})
}
