package inmemory_storage

import (
	"fmt"
	"sync"
	"time"
)

type DataStorage struct {
	sync.Mutex
	cache map[string][]string
}

func NewDataStorage() *DataStorage {
	return &DataStorage{
		cache: make(map[string][]string, 10),
	}
}

func (ds *DataStorage) cleaner(key, val string) {
	time.Sleep(time.Minute)
	if ds.Exist(key, val) {
		ds.DeleteCode(key, val)
	}
}

func (ds *DataStorage) add(key, val string) {
	ds.cache[key] = append(ds.cache[key], val)
}

func (ds *DataStorage) check(key, val string) bool {
	if valIn, ok := ds.cache[key]; ok {
		for _, curVal := range valIn {
			if curVal == val {
				return true
			}
		}
	}
	return false
}

func (ds *DataStorage) delete(key, val string) {
	strs, ok := ds.cache[key]
	if !ok {
		return
	}

	if len(strs) == 1 {
		delete(ds.cache, key)
		return
	}

	newCache := make([]string, 0, len(ds.cache[key])-1)

	for _, s := range strs {
		if s != val {
			newCache = append(newCache, s)
		}
	}

	ds.cache[key] = newCache
}

func (ds *DataStorage) AddCode(key string, val string) {
	ds.Lock()
	ds.add(key, val)
	ds.Unlock()
	go ds.cleaner(key, val)
	fmt.Println(ds.cache)
}

func (ds *DataStorage) Exist(key, val string) bool {
	ds.Lock()
	defer ds.Unlock()
	return ds.check(key, val)
}

func (ds *DataStorage) DeleteCode(key, val string) {
	ds.Lock()
	ds.delete(key, val)
	ds.Unlock()
}
