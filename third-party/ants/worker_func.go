package ants

import "time"

type WorkerWithFunc struct {
	//the work own pool
	pool *PoolWithFunc

	args chan interface{}

	// recycleTime will be update when putting a worker back into queue.
	recycleTime time.Time
}

// run worker
func (w *WorkerWithFunc) run() {
	go func() {
		//loop listen
		for args := range w.args {
			//if a job is nil exit and runing decr
			if args == nil {
				w.pool.decRunning()
				return
			}

			w.pool.poolFunc(args)

			//回收复用
			w.pool.putWorker(w)
		}
	}()
}
