package dataloaders

import (
	"sync"
	"time"
)

type LoaderManyConfig[E any] struct {
	// Fetch is a method that provides the data for the loader
	Fetch func(keys []int, fields []string) ([][]E, []error)

	// Wait is how long wait before sending a batch
	Wait time.Duration

	// MaxBatch will limit the maximum number of keys to send in one batch, 0 = not limit
	MaxBatch int
}

// LoaderMany batches and caches requests
type LoaderMany[E any] struct {
	// this method provides the data for the loader
	fetch func(keys []int, fields []string) ([][]E, []error)

	// how long to done before sending a batch
	wait time.Duration

	// this will limit the maximum number of keys to send in one batch, 0 = no limit
	maxBatch int

	// INTERNAL

	// lazily created cache
	cache map[int][]E

	// the current batch. keys will continue to be collected until timeout is hit,
	// then everything will be sent to the fetch method and out to the listeners
	batch *loaderManyBatch[E]

	// mutex to prevent races
	mu sync.Mutex
}

type loaderManyBatch[E any] struct {
	keys    []int
	fields  []string
	data    [][]E
	error   []error
	closing bool
	done    chan struct{}
}

// Load a TariffOutput by key, batching and caching will be applied automatically
func (l *LoaderMany[E]) Load(key int, fields []string) ([]E, error) {
	return l.LoadThunk(key, fields)()
}

// LoadThunk returns a function that when called will block waiting for a TariffOutput.
// This method should be used if you want one goroutine to make requests to many
// different data loaders without blocking until the thunk is called.
func (l *LoaderMany[E]) LoadThunk(key int, fields []string) func() ([]E, error) {
	l.mu.Lock()
	if it, ok := l.cache[key]; ok {
		l.mu.Unlock()
		return func() ([]E, error) {
			return it, nil
		}
	}
	if l.batch == nil {
		l.batch = &loaderManyBatch[E]{
			done:   make(chan struct{}),
			fields: fields,
		}
	}
	batch := l.batch
	pos := batch.keyIndex(l, key)
	l.mu.Unlock()

	return func() ([]E, error) {
		<-batch.done

		var data []E
		if pos < len(batch.data) {
			data = batch.data[pos]
		}

		var err error
		// its convenient to be able to return a single error for everything
		if len(batch.error) == 1 {
			err = batch.error[0]
		} else if batch.error != nil {
			err = batch.error[pos]
		}

		if err == nil {
			l.mu.Lock()
			l.unsafeSet(key, data)
			l.mu.Unlock()
		}

		return data, err
	}
}

// LoadAll fetches many keys at once. It will be broken into appropriate sized
// sub batches depending on how the loader is configured
func (l *LoaderMany[E]) LoadAll(keys []int, fields []string) ([][]E, []error) {
	results := make([]func() ([]E, error), len(keys))

	for i, key := range keys {
		results[i] = l.LoadThunk(key, fields)
	}

	tariff := make([][]E, len(keys))
	errors := make([]error, len(keys))
	for i, thunk := range results {
		tariff[i], errors[i] = thunk()
	}
	return tariff, errors
}

// LoadAllThunk returns a function that when called will block waiting for a TariffOutput.
// This method should be used if you want one goroutine to make requests to many
// different data loaders without blocking until the thunk is called.
func (l *LoaderMany[E]) LoadAllThunk(keys []int, fields []string) func() ([][]E, []error) {
	results := make([]func() ([]E, error), len(keys))
	for i, key := range keys {
		results[i] = l.LoadThunk(key, fields)
	}
	return func() ([][]E, []error) {
		tariff := make([][]E, len(keys))
		errors := make([]error, len(keys))
		for i, thunk := range results {
			tariff[i], errors[i] = thunk()
		}
		return tariff, errors
	}
}

// Prime the cache with the provided key and value. If the key already exists, no change is made
// and false is returned.
// (To forcefully prime the cache, clear the key first with loader.clear(key).prime(key, value).)
func (l *LoaderMany[E]) Prime(key int, value []E) bool {
	l.mu.Lock()
	var found bool
	if _, found = l.cache[key]; !found {
		// make a copy when writing to the cache, its easy to pass a pointer in from a loop var
		// and end up with the whole cache pointing to the same value.
		cpy := make([]E, len(value))
		copy(cpy, value)
		l.unsafeSet(key, cpy)
	}
	l.mu.Unlock()
	return !found
}

// Clear the value at key from the cache, if it exists
func (l *LoaderMany[E]) Clear(key int) {
	l.mu.Lock()
	delete(l.cache, key)
	l.mu.Unlock()
}

func (l *LoaderMany[E]) unsafeSet(key int, value []E) {
	if l.cache == nil {
		l.cache = map[int][]E{}
	}
	l.cache[key] = value
}

// keyIndex will return the location of the key in the batch, if its not found
// it will add the key to the batch
func (b *loaderManyBatch[E]) keyIndex(l *LoaderMany[E], key int) int {
	for i, existingKey := range b.keys {
		if key == existingKey {
			return i
		}
	}

	pos := len(b.keys)
	b.keys = append(b.keys, key)
	if pos == 0 {
		go b.startTimer(l)
	}

	if l.maxBatch != 0 && pos >= l.maxBatch-1 {
		if !b.closing {
			b.closing = true
			l.batch = nil
			go b.end(l)
		}
	}

	return pos
}

func (b *loaderManyBatch[E]) startTimer(l *LoaderMany[E]) {
	time.Sleep(l.wait)
	l.mu.Lock()

	// we must have hit a batch limit and are already finalizing this batch
	if b.closing {
		l.mu.Unlock()
		return
	}

	l.batch = nil
	l.mu.Unlock()

	b.end(l)
}

func (b *loaderManyBatch[E]) end(l *LoaderMany[E]) {
	b.data, b.error = l.fetch(b.keys, b.fields)
	close(b.done)
}
