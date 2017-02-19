package database

import (
    "log"

    "github.com/boltdb/bolt"
)

func BoltSetup(dbLocation string) {
    boltDb, err := bolt.Open(dbLocation, 0600, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer boltDb.Close()
    boltSetupBucket(boltDb, "ec2")
}

func boltSetupBucket(db *bolt.DB, bucket string) {
    db.Update(func(tx *bolt.Tx) error {
        _, err := tx.CreateBucketIfNotExists([]byte(bucket))
        if err != nil {
            log.Fatalf("Error creating bucket %s - %s", bucket, err)
            return err
        }
        return nil
    }) 
}
