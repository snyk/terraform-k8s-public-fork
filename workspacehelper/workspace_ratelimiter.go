package workspacehelper

import (
	"time"
	"golang.org/x/time/rate"

	rl "k8s.io/client-go/util/workqueue"

)

// WorkspaceRateLimiter is a constructor for a rate limiter for a workqueue. It has
// both overall and per-item rate limiting. The overall is a token bucket and the per-item is exponential
func WorkspaceRateLimiter(baseDelayMultiplier int, maxDelayMultiplier int, startingTokens int, tokensRefilRate int) rl.RateLimiter {
	return rl.NewMaxOfRateLimiter(
		rl.NewItemExponentialFailureRateLimiter(
			time.Duration(baseDelayMultiplier)*time.Second, 
			time.Duration(maxDelayMultiplier)*time.Minute),
		// tokensRefilRate qps, startingTokens bucket size.  
		// This is only for retry speed and its only the overall factor (not per item)
		&rl.BucketRateLimiter{Limiter: rate.NewLimiter(rate.Limit(tokensRefilRate), startingTokens)},
	)
}
