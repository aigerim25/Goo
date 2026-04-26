package filtering

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func ExecuteWithLogs(ctx context.Context, client *http.Client, req *http.Request, maxAttempts int) (*http.Response, error) {
	for attempt := 0; attempt < maxAttempts; attempt++ {
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}
		reqWithCtx := req.Clone(ctx)
		resp, err := client.Do(reqWithCtx)
		if !IsRetryable(resp, err) {
			fmt.Printf("Attempt %d: Success!\n", attempt+1)
			return resp, err
		}
		if attempt == maxAttempts-1 {
			return resp, err
		}
		delay := CalculateBackOff(attempt)
		fmt.Printf("Attempt %d failed: waiting %d ms...\n", attempt+1, delay.Milliseconds())
		select {
		case <-time.After(delay):
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
	return nil, ctx.Err()
}
