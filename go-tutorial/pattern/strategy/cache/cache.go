package cache

import (
	"fmt"
	"sort"
	"time"
)

type Cache struct {
	//storage      map[string]string
	storage      map[string]map[string]time.Time
	evictionAlgo EvictionAlgo
	capacity     int
	maxCapacity  int
}

func InitCache(e EvictionAlgo) *Cache {
	//storage := make(map[string]string)
	storage := make(map[string]map[string]time.Time)
	return &Cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *Cache) SetEviction(e EvictionAlgo) {
	c.evictionAlgo = e
}

// func (c *Cache) Add(key, value string) {
func (c *Cache) Add(key string, value map[string]time.Time) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) AddDefault(key, value string) {
	c.Add(key, map[string]time.Time{value: time.Now()})
}

// func (c *Cache) Get(key string) string {
func (c *Cache) Get(key string) map[string]time.Time {
	value := c.storage[key]
	delete(c.storage, key)
	return value
}

func (c *Cache) GetAll() map[string]map[string]time.Time {
	return c.storage
}

func (c *Cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

type EvictionAlgo interface {
	evict(c *Cache)
}

type Fifo struct {
}

func (f *Fifo) evict(c *Cache) {
	fmt.Println("Evicting by fifo strategy, first in first out")
	//get time and key
	tempTime := make(map[time.Time]string)

	for key, value := range c.storage {
		for _, value := range value {
			tempTime[value] = key
		}
	}
	//order range times
	keys := []time.Time{}
	for key := range tempTime {
		keys = append(keys, key)
	}

	//sort times
	sort.Sort(SliceTime(keys))

	//get key of the lowes time
	keyForDelete := tempTime[keys[0]]

	//remove key of storage where is the lowest time
	delete(c.storage, keyForDelete)
}

func (f *Lru) evict(c *Cache) {
	fmt.Println("Evicting by lru strategy, lease recently use")
}

type Lfu struct {
}

func (f *Lfu) evict(c *Cache) {
	fmt.Println("Evicting by lru strategy, lease frequently use")
}
