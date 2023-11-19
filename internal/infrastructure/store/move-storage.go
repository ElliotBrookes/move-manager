package infrastructure

import "github.com/syndtr/goleveldb/leveldb"

// Contains functionality to return instance capable of storing info for current move

type moveStorage struct {
	readOnly bool
	db       *leveldb.DB
}

func (db moveStorage) Get(key []byte) ([]byte, error) {
	return nil, nil
}

func (db moveStorage) Put(key, value []byte) error {
	return nil
}

func (db moveStorage) Delete(key []byte) error {
	return nil
}
