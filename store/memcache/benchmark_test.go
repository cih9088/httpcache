//go:build httpcache_acceptance_benchmarks
// +build httpcache_acceptance_benchmarks

package memcache

import (
	"testing"

	"github.com/cih9088/httpcache/store/acceptance"
	"github.com/cih9088/httpcache/store/driver"
)

func BenchmarkMemCache(b *testing.B) {
	acceptance.RunB(b, acceptance.FactoryFunc(func() (driver.Conn, func()) {
		cache := Open()
		return cache, func() {}
	}))
}
