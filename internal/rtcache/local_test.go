package rtcache

import (
	"testing"
)

func TestLocalCache(t *testing.T) {
	rtCache := NewLocalCache()
	testCache(t, rtCache)
}

// func TestLocalConsumers(t *testing.T) {
// 	rtJobs := NewLocalJobs()
// 	rtCache := NewLocalCache()
// 	testConsumers(t, rtCache, rtJobs)
// }
