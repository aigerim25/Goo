package filtering

import (
	"context"
	"net/http"
	"time"
)

func ExecutePayment(
	ctx context.Context,
	client *http.Client,
	req *http.Request,
	maxAttempts int,
) (*http.Response, error) {
	for attempt := 0; attempt < maxAttempts; attempt++ {
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}
		reqWithCtx := req.Clone(ctx)
		resp, err := client.Do(reqWithCtx)

		if !IsRetryable(resp, err) {
			return resp, err
		}
		if attempt == maxAttempts-1 {
			return resp, err
		}
		delay := CalculateBackOff(attempt)

		select {
		case <-time.After(delay):
			continue
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
	return nil, ctx.Err()
}
