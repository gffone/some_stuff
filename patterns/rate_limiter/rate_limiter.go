package rate_limiter

import "time"

type RateLimiter struct {
	leakyBucketCh chan struct{}

	closeCh     chan struct{}
	closeDoneCh chan struct{}
}

func NewLeakyBucketLimiter(limit int, period time.Duration) RateLimiter {
	limiter := RateLimiter{
		leakyBucketCh: make(chan struct{}, limit),
		closeCh:       make(chan struct{}),
		closeDoneCh:   make(chan struct{}),
	}

	leakInterval := period.Nanoseconds() / int64(limit)

	go limiter.startPeriodLeak(time.Duration(leakInterval))
	return limiter
}

func (l *RateLimiter) startPeriodLeak(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer func() {
		ticker.Stop()
		close(l.leakyBucketCh)
	}()

	for {
		select {
		case <-l.closeCh:
			return
		case <-ticker.C:
			select {
			case <-l.leakyBucketCh:
			default:
			}
		}
	}

}

func (l *RateLimiter) Allow() bool {
	select {
	case l.leakyBucketCh <- struct{}{}:
		return true
	default:
		return false
	}
}

func (l *RateLimiter) Shutdown() {
	close(l.closeCh)
	<-l.closeDoneCh
}
