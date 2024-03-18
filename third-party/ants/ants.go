package ants

import (
	"errors"
	"math"
)

var (
	ErrorPoolClosed = errors.New("this pool has been closed")

	ErrInvalidPoolSize = errors.New("pool size invalid")

	ErrInvalidPoolExpiry = errors.New("pool expiry invalid")
)

const (
	DefaultAntsPoolSize = math.MaxInt32

	DefaultCleanIntervalTime = 5
)

var defaultAntsPool, _ = NewPool(DefaultAntsPoolSize)

func Submit(task f) error {
	return defaultAntsPool.Submit(task)
}

func Running() int {
	return defaultAntsPool.Running()
}

func Cap() int {
	return defaultAntsPool.Capacity()
}

func Free() int {
	return defaultAntsPool.Free()
}

func Release() {
	defaultAntsPool.Release()
}
