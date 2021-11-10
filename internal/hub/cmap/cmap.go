package cmap

import (
	"github.com/midoks/go-p2p-server/internal/client"
	"sync"
)

// A "thread" safe string to anything map.
type ConMapShared struct {
	items        map[string]*client.Client
	sync.RWMutex // Read Write mutex, guards access to internal map.
}

// Used by the Iter & IterBuffered functions to wrap two variables together over a channel,
type Tuple struct {
	Key string
	Val *client.Client
}

var SHARD_COUNT = 32

// A "thread" safe map of type string:Anything.
// To avoid lock bottlenecks this map is dived to several (SHARD_COUNT) map shards.
type ConMap []*ConMapShared

func New() ConMap {
	m := make(ConMap, SHARD_COUNT)
	for i := 0; i < SHARD_COUNT; i++ {
		m[i] = &ConMapShared{items: make(map[string]*client.Client)}
	}
	return m
}

// GetShard returns shard under given key
func (m ConMap) GetShard(key string) *ConMapShared {
	return m[uint(fnv32(key))%uint(SHARD_COUNT)]
}

func (m ConMap) MSet(data map[string]*client.Client) {
	for key, value := range data {
		shard := m.GetShard(key)
		shard.Lock()
		shard.items[key] = value
		shard.Unlock()
	}
}

// Sets the given value under the specified key.
func (m ConMap) Set(key string, value *client.Client) {
	// Get map shard.
	shard := m.GetShard(key)
	shard.Lock()
	shard.items[key] = value
	shard.Unlock()
}

// Get retrieves an element from map under given key.
func (m ConMap) Get(key string) (*client.Client, bool) {
	// Get shard
	shard := m.GetShard(key)
	shard.RLock()
	// Get item from shard.
	val, ok := shard.items[key]
	shard.RUnlock()
	return val, ok
}

// Looks up an item under specified key
func (m ConMap) Has(key string) bool {
	// Get shard
	shard := m.GetShard(key)
	shard.RLock()
	// See if element is within shard.
	_, ok := shard.items[key]
	shard.RUnlock()
	return ok
}

// Remove removes an element from the map.
func (m ConMap) Remove(key string) {
	// Try to get shard.
	shard := m.GetShard(key)
	shard.Lock()
	delete(shard.items, key)
	shard.Unlock()
}

// Clear removes all items from map.
func (m ConMap) Clear() {
	for item := range m.IterBuffered() {
		m.Remove(item.Key)
	}
}

func (m ConMap) CountPerMapNoLock() []int {
	var count = make([]int, SHARD_COUNT)
	for i := 0; i < SHARD_COUNT; i++ {
		shard := m[i]
		count[i] = len(shard.items)
	}
	return count
}

func (m ConMap) CountNoLock() int {
	count := 0
	for i := 0; i < SHARD_COUNT; i++ {
		shard := m[i]
		count += len(shard.items)
	}
	return count
}

// Returns a array of channels that contains elements in each shard,
// which likely takes a snapshot of `m`.
// It returns once the size of each buffered channel is determined,
// before all the channels are populated using goroutines.
func snapshot(m ConMap) (chans []chan Tuple) {
	chans = make([]chan Tuple, SHARD_COUNT)
	wg := sync.WaitGroup{}
	wg.Add(SHARD_COUNT)
	// Foreach shard.
	for index, shard := range m {
		go func(index int, shard *ConMapShared) {
			// Foreach key, value pair.
			shard.RLock()
			chans[index] = make(chan Tuple, len(shard.items))
			wg.Done()
			for key, val := range shard.items {
				chans[index] <- Tuple{key, val}
			}
			shard.RUnlock()
			close(chans[index])
		}(index, shard)
	}
	wg.Wait()
	return chans
}

// IterBuffered returns a buffered iterator which could be used in a for range loop.
func (m ConMap) IterBuffered() <-chan Tuple {
	chans := snapshot(m)
	total := 0
	for _, c := range chans {
		total += cap(c)
	}
	ch := make(chan Tuple, total)
	go fanIn(chans, ch)
	return ch
}

// fanIn reads elements from channels `chans` into channel `out`
func fanIn(chans []chan Tuple, out chan Tuple) {
	wg := sync.WaitGroup{}
	wg.Add(len(chans))
	for _, ch := range chans {
		go func(ch chan Tuple) {
			for t := range ch {
				out <- t
			}
			wg.Done()
		}(ch)
	}
	wg.Wait()
	close(out)
}

func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	for i := 0; i < len(key); i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}
