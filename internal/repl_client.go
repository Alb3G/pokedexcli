package internal

import (
	"net/http"
	"time"
)

type Client struct {
	cache      Cache
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	client := Client{
		cache:      NewCache(cacheInterval),
		httpClient: http.Client{Timeout: timeout},
	}

	return client
}
