package util

import (
	"hash/fnv"
	"strconv"
	"sync"
)

// Number of shards (adjust based on hardware and workload)
const shardCount = 32

// A single shard with its own map and RWMutex
type shard struct {
	sync.RWMutex
	data map[int]int
}

// ShardedMap structure
type ShardedMap struct {
	shards []*shard
}

// NewShardedMap initializes the sharded map
func NewShardedMap() *ShardedMap {
	sm := &ShardedMap{
		shards: make([]*shard, shardCount),
	}
	for i := 0; i < shardCount; i++ {
		sm.shards[i] = &shard{
			data: make(map[int]int),
		}
	}
	return sm
}

// hashKey determines which shard a key belongs to
func (sm *ShardedMap) hashKey(key int) int {
	h := fnv.New32a()
	h.Write([]byte(strconv.Itoa(key)))
	return int(h.Sum32()) % shardCount
}

// Set adds or updates a value in the map
func (sm *ShardedMap) Set(key, value int) {
	shard := sm.shards[sm.hashKey(key)]
	shard.Lock()
	defer shard.Unlock()
	shard.data[key] = value
}

// Get retrieves a value from the map
func (sm *ShardedMap) Get(key int) (int, bool) {
	shard := sm.shards[sm.hashKey(key)]
	shard.RLock()
	defer shard.RUnlock()
	value, exists := shard.data[key]
	return value, exists
}

// Delete removes a key-value pair from the map
func (sm *ShardedMap) Delete(key int) {
	shard := sm.shards[sm.hashKey(key)]
	shard.Lock()
	defer shard.Unlock()
	delete(shard.data, key)
}

// Exists checks if a key exists in the map
func (sm *ShardedMap) Exists(key int) bool {
	shard := sm.shards[sm.hashKey(key)]
	shard.RLock()
	defer shard.RUnlock()
	_, exists := shard.data[key]
	return exists
}

// Keys returns all keys in the map (optional feature)
func (sm *ShardedMap) Keys() []int {
	keys := []int{}
	wg := sync.WaitGroup{}
	mu := sync.Mutex{} // Protects the keys slice

	wg.Add(shardCount)
	for _, s := range sm.shards {
		go func(sh *shard) {
			defer wg.Done()
			sh.RLock()
			defer sh.RUnlock()
			for key := range sh.data {
				mu.Lock()
				keys = append(keys, key)
				mu.Unlock()
			}
		}(s)
	}
	wg.Wait()
	return keys
}

// Count returns the total number of key-value pairs in the map
func (sm *ShardedMap) Count() int {
	totalCount := 0
	wg := sync.WaitGroup{}
	mu := sync.Mutex{} // Protects the totalCount variable

	wg.Add(shardCount)
	for _, s := range sm.shards {
		go func(sh *shard) {
			defer wg.Done()
			sh.RLock()
			defer sh.RUnlock()
			count := len(sh.data)
			mu.Lock()
			totalCount += count
			mu.Unlock()
		}(s)
	}
	wg.Wait()
	return totalCount
}

var muSD sync.Mutex

func SafeDecrement(total *int32) {
	muSD.Lock()
	if *total > 0 {
		*total--
	}
	muSD.Unlock()
}
