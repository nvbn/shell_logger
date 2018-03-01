package shell

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

// Creates the database if it doesn't exist, otherwise opens the database and creates a bucket for key:value pairs
func SetupDatabase(path string) error {
	DBPath = path
	db := openConn()
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Mappings"))
		if err != nil {
			fmt.Errorf("Failed to create bucket: %v", err)
			return err
		}
		return nil
	})
	return nil
}

func openConn() *bolt.DB {
	db, err := bolt.Open(DBPath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// Gets the JSON Object from the database
func GetJsonValues(db *bolt.DB, key []byte) ([]byte, error) {
	var jsonValue []byte
	if db == nil {
		fmt.Errorf("DB is null")
	}

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Mappings"))
		if err != nil {
			fmt.Errorf("Failed to create bucket: %v", err)
			return err
		}
		return nil
	})

	err := db.View(func(tx *bolt.Tx) error {
		jsonValue = tx.Bucket([]byte("Mappings")).Get(key)
		return nil
	})
	if err != nil {
		fmt.Errorf("Got json retrieval error: %v", err)
	}
	return jsonValue, err
}

// Gets the list of good commands from the database
func GetGoodCommands(key []byte) ([][]byte, error) {
	db := openConn()
	defer db.Close()

	var vals [][]byte
	fmt.Println()
	jsonObject, err := GetJsonValues(db, key)
	if err != nil {
		fmt.Errorf("Failed to retrieve values: %v", err)
		return nil, err
	}
	json.Unmarshal(jsonObject, &vals)
	if err != nil {
		fmt.Errorf("Failed to unmarshal: %v", err)
	}
	return vals, err
}

// Inserts the key:value pair of correctCommand:incorrectCommand into the database
func Insert(correct []byte, incorrect []byte) error {
	firstWord := GetFirstCommand(correct)
	if firstWord == nil {
		return nil
	}
	correctCommands, err := GetGoodCommands(firstWord)
	if err != nil {
		return err
	}
	db := openConn()
	defer db.Close()
	if correctCommands == nil {
		fmt.Println("Correct commands is nil")
		err := db.Update(func(tx *bolt.Tx) error {
			correctCommand := [1][]byte{correct}
			jsonObject, err := json.Marshal(correctCommand)
			if err != nil {
				fmt.Errorf("Failed to marshal: %v", err)
				return err
			}
			err = tx.Bucket([]byte("Mappings")).Put(firstWord, jsonObject)
			if err != nil {
				fmt.Errorf("Failed to insert values: %v", err)
				return err
			}
			return nil
		})
		return err
	} else {
		jsonObject, err := GetJsonValues(db, firstWord)
		if err != nil {
			return err
		}
		err = db.Update(func(tx *bolt.Tx) error {
			var vals [][]byte
			err = json.Unmarshal(jsonObject, &vals)
			if err != nil {
				fmt.Errorf("Failed to unmarshal object: %v", err)
				return err
			}
			vals = append(vals, correct)
			newJsonObject, err := json.Marshal(vals)
			if err != nil {
				fmt.Errorf("Failed to marshal object: %v", err)
				return err
			}
			err = tx.Bucket([]byte("Mappings")).Put(firstWord, newJsonObject)
			if err != nil {
				return fmt.Errorf("Could not set value: %v", err)
				return err
			}
			return nil
		})
		return err
	}

}
