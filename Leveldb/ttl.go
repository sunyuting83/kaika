package LevelDB

import (
	"encoding/json"
	"kaika/database"
	"strings"
	"time"

	leveldb_errors "github.com/syndtr/goleveldb/leveldb/errors"
)

type CacheType struct {
	Data    []byte `json:"data"`
	Created int64  `json:"created"`
	Expires int64  `json:"expires"`
}

func Get(key string) ([]byte, error) {

	data, err := LevelDB.Get([]byte(key), nil)

	if err != nil && err != leveldb_errors.ErrNotFound {
		return nil, err
	}

	if data == nil {
		return nil, nil
	}

	var cache CacheType
	err = json.Unmarshal(data, &cache)

	if err != nil {
		return nil, nil
	}

	secs := time.Now().Unix()

	if cache.Expires > 0 && cache.Expires <= secs {
		LevelDB.Delete([]byte(key), nil)
		return nil, nil
	}

	return cache.Data, nil
}

func CardGet(key, computer string, login bool) ([]byte, error) {

	data, err := LevelDB.Get([]byte(key), nil)

	if err != nil && err != leveldb_errors.ErrNotFound {
		return nil, err
	}

	if data == nil {
		return nil, nil
	}

	var cache CacheType
	err = json.Unmarshal(data, &cache)

	if err != nil {
		return nil, nil
	}
	id := strings.Split(string(cache.Data), "----")[1]

	secs := time.Now().Unix()

	if cache.Expires > 0 && cache.Expires <= secs {

		var card database.Card
		card.DeleteCard(id)
		LevelDB.Delete([]byte(key), nil)
		return nil, nil
	}

	if login {
		NewData := []byte(strings.Join([]string{computer, id}, "----"))
		CardSet(key, NewData, cache.Created, cache.Expires)
	}

	return cache.Data, nil
}

func CardSet(key string, value []byte, created int64, expires int64) {
	cache := CacheType{Data: value, Created: created, Expires: expires}
	jsonString, _ := json.Marshal(cache)
	LevelDB.Put([]byte(key), jsonString, nil)
}

func Set(key string, value string, expires int64) {
	cache := CacheType{Data: []byte(value), Created: time.Now().Unix(), Expires: 0}

	if expires > 0 {
		cache.Expires = cache.Created + expires
	}
	jsonString, _ := json.Marshal(cache)

	LevelDB.Put([]byte(key), jsonString, nil)
}
