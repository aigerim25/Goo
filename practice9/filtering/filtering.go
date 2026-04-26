package filtering

import (
	"math/rand"
	"net/http"
	"time"
)

func IsRetryable(resp *http.Response, err error) bool {
	if err != nil {
		return true
	}
	if resp == nil {
		return false
	}
	switch resp.StatusCode {
	case 429, 500, 502, 503, 504:
		return true
	case 401, 404:
		return false
	default:
		return false
	}
}

func CalculateBackOff(attempt int) time.Duration {
	baseDelay := 100 * time.Millisecond
	backoff := baseDelay * time.Duration(1<<attempt)
	jitter := time.Duration(rand.Int63n(int64(backoff)))
	return jitter
}
