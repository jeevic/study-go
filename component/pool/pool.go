// Example provided with help from Fatih Arslan and Gabriel Aszalos.
// Package pool manages a user defined set of resources.
package pool

import (
	"errors"
	"io"
	"log"
	"sync"
	"sync/atomic"
)

// Pool manages a set of resources that can be shared safely by
// multiple goroutines. The resource being managed must implement
// the io.Closer interface.
type Pool struct {
	m         sync.Mutex
	Resources chan io.Closer
	capacity  int32
	running   int32
	factory   func() (io.Closer, error)
	closed    bool
}

// ErrPoolClosed is returned when an Acquire returns on a
// closed pool.
var ErrPoolClosed = errors.New("Pool has been closed.")

// New creates a pool that manages resources. A pool requires a
// function that can allocate a new resource and the size of
// the pool.
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Size value too small.")
	}

	return &Pool{
		factory:   fn,
		capacity:  int32(size),
		running:   int32(0),
		Resources: make(chan io.Closer, size),
	}, nil
}

// Acquire retrieves a resource	from the pool.
func (p *Pool) Acquire() (io.Closer, error) {
	p.m.Lock()
	defer p.m.Unlock()

	for {
		select {
		// Check for a free resource.
		case r, ok := <-p.Resources:
			log.Println("Acquire:", "Shared Resource")
			if !ok {
				return nil, ErrPoolClosed
			}
			atomic.AddInt32(&p.running, 1)
			return r, nil

			// Provide a new resource since there are none available.
		default:
			if p.running < p.capacity {
				log.Println("Acquire:", "New Resource")
				c, e := p.factory()
				if e != nil {

				}
				atomic.AddInt32(&p.running, 1)
				return c, e
			}
		}
	}

}

// Release places a new resource onto the pool.
func (p *Pool) Release(r io.Closer) {
	// Secure this operation with the Close operation.
	// If the pool is closed, discard the resource.
	if p.closed {
		r.Close()
		return
	}

	select {
	// Attempt to place the new resource on the queue.
	case p.Resources <- r:
		atomic.AddInt32(&p.running, -1)
		log.Println("Release:", "In Queue")

	// If the queue is already at cap we close the resource.
	default:
		atomic.AddInt32(&p.running, -1)
		log.Println("Release:", "Closing")
		r.Close()
	}
	//log.Println("chan-len*************",len(p.Resources))
	//log.Println("chan-len*************",<-p.Resources)
}

// Close will shutdown the pool and close all existing resources.
func (p *Pool) Close() {
	// Secure this operation with the Release operation.
	p.m.Lock()
	defer p.m.Unlock()

	// If the pool is already close, don't do anything.
	if p.closed {
		return
	}

	// Set the pool as closed.
	p.closed = true

	// Close the channel before we drain the channel of its
	// resources. If we don't do this, we will have a deadlock.
	close(p.Resources)
	atomic.StoreInt32(&p.running, 0)
	// Close the resources
	for r := range p.Resources {
		r.Close()
	}
}
