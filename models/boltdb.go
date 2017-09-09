package models

import (
	"log"
	"time"

	"github.com/boltdb/bolt"
	"github.com/tonyalaribe/dsi-hackaton/constants"
)

// Client represents a client to the underlying BoltDB data store.
type Client struct {
	Path string   //Filename to the bolt db file
	DB   *bolt.DB //pointer to a bolt instance
}

// Open opens and initializes the BoltDB database.
func (c *Client) Open() error {
	// Open database file.
	db, err := bolt.Open(c.Path, 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Println(c.Path)
		log.Println(err)
		return err
	}
	c.DB = db

	// Start writable transaction.
	tx, err := c.DB.Begin(true)
	if err != nil {
		log.Println(err)
		return err
	}
	defer tx.Rollback()

	// Initialize top-level buckets.
	if _, err = tx.CreateBucketIfNotExists([]byte(constants.LOCATION_BUCKET)); err != nil {
		log.Println(err)
		return err
	}

	// Save transaction to disk.
	return tx.Commit()
}
