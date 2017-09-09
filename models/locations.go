package models

import (
	"encoding/json"
	"log"

	"github.com/boltdb/bolt"
	"github.com/tonyalaribe/dsi-hackaton/constants"
)

type Location struct {
	LocationID   int
	LocationName string
	Good         int
	Bad          int
	GoodBool     bool
	Latitude     float64
	Longitude    float64
}

func Create(db *bolt.DB, location Location) error {

	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	b := tx.Bucket([]byte(constants.LOCATION_BUCKET))

	number, err := b.NextSequence()
	if err != nil {
		return err
	}

	location.LocationID = int(number)
	locationJSON, err := json.Marshal(location)
	if err != nil {
		log.Println(err)
	}
	err = b.Put(itob(int(number)), locationJSON)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func UpdateLocation(db *bolt.DB, newLocation Location) error {
	log.Println(newLocation)
	tx, err := db.Begin(true)
	if err != nil {
		log.Println(err)
	}
	defer tx.Rollback()

	b := tx.Bucket([]byte(constants.LOCATION_BUCKET))

	oldLocationByte := b.Get(itob(newLocation.LocationID))
	oldLocation := Location{}
	err = json.Unmarshal(oldLocationByte, &oldLocation)
	if err != nil {
		log.Println(err)
		//return err
	}

	log.Printf("old location %+v", oldLocation)
	if newLocation.GoodBool {
		oldLocation.Good++
	} else {
		oldLocation.Bad++
	}
	log.Printf("new old location %+v", oldLocation)

	oldLocationJSON, err := json.Marshal(oldLocation)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(oldLocationJSON))
	err = b.Put(itob(newLocation.LocationID), oldLocationJSON)
	if err != nil {
		log.Println(err)
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil

}

func GetAll(db *bolt.DB) ([]Location, error) {
	tx, err := db.Begin(false)
	if err != nil {
		return []Location{}, err
	}
	defer tx.Rollback()

	b := tx.Bucket([]byte(constants.LOCATION_BUCKET))

	c := b.Cursor()

	locations := []Location{}

	for k, v := c.Last(); k != nil; k, v = c.Prev() {
		x := Location{}
		err = json.Unmarshal(v, &x)
		if err != nil {
			log.Println(err)
		}
		locations = append(locations, x)
	}

	return locations, nil

}
