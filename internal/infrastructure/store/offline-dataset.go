package infrastructure

import (
	"ElliotBrookes/move-manager/internal/domain"
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/syndtr/goleveldb/leveldb"
)

type offlineDatasetStorage struct{}

// Return an instance of the offline dataset

func ReturnOfflineDataset() (*leveldb.DB, error) {
	//
	// Open the current db, if there isnt one, initialise it
	//

	fmt.Println("Loading Postcode Data...")

	// Checking if db already exists
	_, err := os.Stat("dataset.db")
	if err != nil {
		exists := os.IsExist(err)
		if !exists {
			// Open dataset and seed a new bolt db
			return initLevelDB(), nil

		} else {
			// Error out if error not about file existance
			return nil, err
		}

	}

	// If one does exist
	db, err := leveldb.OpenFile("dataset.db", nil)
	if err != nil {
		log.Fatal("Unable to Open Existing Bolt DB")
	}

	return db, nil
}

func initLevelDB() *leveldb.DB {
	// Open & Create new database
	db, err := leveldb.OpenFile("dataset.db", nil)
	if err != nil {
		log.Fatal("Unable to Init New Bolt DB")
	}

	// Create and fill a batch change with dataset
	initBatch := new(leveldb.Batch)
	scannerReadCsv(initBatch)

	// Write the batched changes
	err = db.Write(initBatch, nil)

	return db
}

func scannerReadCsv(batch *leveldb.Batch) error {
	file, err := os.Open("ukpostcodes.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	// init Reader and go through Initial column line
	reader := bufio.NewScanner(file)
	_ = reader.Scan()
	if err != nil {
		return errors.New("cant read column line of dataset")
	}

	for reader.Scan() {
		// Read a row and print a . to signify
		text := reader.Text()
		row := strings.Split(text, ",")
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// Prepare and Insert all rows of data
		postcode, coordinates, err := prepareRow(row)
		if err != nil {
			return err
		}

		batch.Put(postcode, coordinates)
		if err != nil {
			return err
		}

	}

	return nil
}

func prepareRow(row []string) ([]byte, []byte, error) {
	postcode := []byte(row[1])
	latitude, err := strconv.ParseFloat(row[2], 64)
	if err != nil {
		fmt.Println(row)
	}

	longitude, err := strconv.ParseFloat(row[3], 64)
	if err != nil {
		fmt.Println(row)
	}

	coords := domain.Coordinates{
		Long: longitude,
		Lat:  latitude,
	}

	marshalledBytes, err := json.Marshal(coords)
	if err != nil {
		log.Fatal(err)
	}

	return postcode, marshalledBytes, nil
}
