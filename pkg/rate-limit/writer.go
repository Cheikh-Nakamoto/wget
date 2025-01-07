package ratelimit

import (
	"io"
	"time"
)

/*type rateLimitWriter struct {
	limiter *rate.Limiter
}

func (rlw *rateLimitWriter) Write(p []byte) (int, error) {
	n := len(p)
	err := rlw.limiter.WaitN(nil, n)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func parseRateLimit(rateLimit string) (int, error) {
	var multiplier int
	var value int
	switch {
	case rateLimit[len(rateLimit)-1] == 'k':
		multiplier = 1024
	case rateLimit[len(rateLimit)-1] == 'M':
		multiplier = 1024 * 1024
	default:
		return 0, fmt.Errorf("invalid rate limit format: %s", rateLimit)
	}

	fmt.Sscanf(rateLimit[:len(rateLimit)-1], "%d", &value)
	return value * multiplier, nil
}*/

type RateLimitedReader struct {
	Reader  io.Reader
	Limiter *RateLimiter
}
type RateLimiter struct {
	rate   int // Débit en octel par seconde
	burst  int // Nombre maximal de octel dans le seau
	tokens int // Nombre actuel de octel dans le seau
}

func NewRateLimitedReader(reader io.Reader, rateLimit int) *RateLimitedReader {
	limiter := NewRateLimiter(rateLimit, rateLimit)
	return &RateLimitedReader{
		Reader:  reader,
		Limiter: limiter,
	}
}

func NewRateLimiter(rate, burst int) *RateLimiter {
	return &RateLimiter{
		rate:   rate,
		burst:  burst,
		tokens: burst,
	}
}

func (r *RateLimitedReader) Read(p []byte) (int, error) {
	n, err := r.Reader.Read(p)
	if err != nil {
		return n, err
	}

	if r.Limiter != nil {
		//Calcule le temps nécessaire pour lire les données et réserve des jetons
		timeRequired := r.Limiter.Reserve(n)
		time.Sleep(timeRequired)
	}

	return n, nil
}

func (r *RateLimiter) Reserve(bytes int) time.Duration {
	requiredTokens := float64(bytes)
	timeRequired := time.Duration(requiredTokens / float64(r.rate) * float64(time.Second))
	return timeRequired
}
