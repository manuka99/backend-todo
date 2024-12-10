package utils

import (
	"time"
)

// FetchRequest represents a request for fetching a specific type of data.
type FetchRequest[K comparable, T any] struct {
	ID       K
	Response chan T
	Error    chan error
}

// StartDataFetcher starts the data-fetching worker to retrieve data and optionally cache it.
func StartDataFetcher[K comparable, T any](
	fetchChan chan FetchRequest[K, T],
	fetchDataFunc func(K) (T, error), // Accepts a function that fetches data for a specific type
	cache *Cache[K, T], // Cache is optional
	cacheTTL time.Duration, // TTL for cache entries
) {
	go func() {
		for req := range fetchChan {
			// Check the cache before fetching data, if cache is provided
			if cache != nil {
				if data, found := cache.Get(req.ID); found {
					req.Response <- data
					continue
				}
			}

			// Fetch data from the provided fetchDataFunc
			data, err := fetchDataFunc(req.ID)
			if err != nil {
				req.Error <- err
				continue
			}

			// If a cache is provided, add the data to the cache
			if cache != nil {
				// Set default TTL if none is provided
				if cacheTTL == 0 {
					cacheTTL = 5 * time.Minute
				}
				cache.Set(req.ID, data, cacheTTL)
			}

			// Send the result back to the caller
			req.Response <- data
		}
	}()
}
