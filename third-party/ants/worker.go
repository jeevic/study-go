package ants

import "time"

type Worker struct {
	//the work own pool
	pool *Pool

	//task is a job queue
	task chan f

	// recycleTime will be update when putting a worker back into queue.
	recycleTime time.Time
}

// run worker
func (w *Worker) run() {
	go func() {
		//loop listen
		for f := range w.task {
			//if a job is nil exit and runing decr
			if f == nil {
				w.pool.decRunning()
				return
			}

			f()

			//回收复用
			w.pool.putWorker(w)
		}
	}()
}

func (w *Worker) stop() {
	w.sendTask(nil)
}

func (w *Worker) sendTask(task f) {
	w.task <- task
}
