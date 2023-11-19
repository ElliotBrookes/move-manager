package infrastructure

import (
	"ElliotBrookes/move-manager/internal/adapters/store"
	"errors"
)

const (
	DataSet = iota
	MoveData
)

func NewLevelDB(iType int) (store.KVStore, error) {
	switch iType {
	case DataSet:
		return offlineDatasetStorage{}, nil
	case MoveData:
		return moveStorage{}, nil
	}

	return nil, errors.New("invalid leveldb type provided")
}
