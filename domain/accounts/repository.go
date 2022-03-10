package accounts

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nuts-foundation/nuts-pgo-demo/domain/types"
	"os"
	"sort"
	"sync"
)

type Repository interface {
	AddAccount(account types.Account) (*types.Account, error)
	FindByUserName(username string) (*types.Account, error)
	All() ([]types.Account, error)
}

type jsonFileRepo struct {
	filepath string
	mutex    sync.RWMutex
	// records is a cache
	records []types.Account
}

func NewJsonFileRepository(filepath string) *jsonFileRepo {
	f, err := os.OpenFile(filepath, os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	repo := jsonFileRepo{
		filepath: filepath,
		mutex:    sync.RWMutex{},
		records:  make([]types.Account, 0),
	}

	if err := repo.readAll(); err != nil {
		panic(err)
	}

	return &repo
}

func (db *jsonFileRepo) AddAccount(account types.Account) (*types.Account, error) {
	db.mutex.Lock()
	db.mutex.Unlock()

	if err := db.readAll(); err != nil {
		return nil, err
	}
	for _, a := range db.records {
		if a.Username == account.Username {
			return nil, errors.New("account already exists")
		}
	}

	db.records = append(db.records, account)

	return &account, db.writeAll()
}

func (db *jsonFileRepo) FindByUserName(username string) (*types.Account, error) {
	if len(db.records) == 0 {
		if err := db.readAll(); err != nil {
			return nil, err
		}
	}

	db.mutex.RLock()
	defer db.mutex.RUnlock()

	for _, a := range db.records {
		if a.Username == username {
			// Hazardous to return a pointer, but this is a demo.
			return &a, nil
		}
	}

	// Not found
	return nil, types.NotFoundErr
}

func (db *jsonFileRepo) readAll() error {
	bytes, err := os.ReadFile(db.filepath)
	if err != nil {
		return fmt.Errorf("unable to read db from file: %w", err)
	}

	if len(bytes) == 0 {
		return nil
	}

	if err = json.Unmarshal(bytes, &db.records); err != nil {
		return fmt.Errorf("unable to unmarshall db from file: %w", err)
	}

	return nil
}

// WriteAll writes all records to the file, truncating the file if it exists
func (db *jsonFileRepo) writeAll() error {
	bytes, err := json.Marshal(db.records)
	if err != nil {
		return fmt.Errorf("unable to marshall db records to json: %w", err)
	}

	if err = os.WriteFile(db.filepath, bytes, 0666); err != nil {
		return fmt.Errorf("unable to write db to file: %w", err)
	}
	return nil
}

// All returns every record sorted by username
func (db *jsonFileRepo) All() ([]types.Account, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	if err := db.readAll(); err != nil {
		return nil, err
	}

	v := make([]types.Account, len(db.records))

	for idx, value := range db.records {
		v[idx] = value
	}
	sort.SliceStable(v, func(i, j int) bool {
		return v[i].Username < v[j].Username
	})
	return v, nil
}
