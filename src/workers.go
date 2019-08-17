// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// Miscellaneous utility functions

// Note that these functions are made unique by attaching them to a structure
// which must be instanced before they can be used.

// Some functions were taken from https://blog.kowalczyk.info/book/go-cookbook.html
// which was declared public domain at the time that the functions were taken.

package util

import (
	"runtime"
)

//============================================================================
//                            		Workers
//============================================================================

// WorkQueue represents a one use only structure to execute multiple
// go routines easily.
type WorkQueue struct {
	queue		chan interface{}
	ack			chan bool
	done		chan bool
	size		int
}

func (w *WorkQueue) CloseQueue() {
	close(w.queue)
}

func (w *WorkQueue) CloseAndWaitForCompletion() {
	close(w.queue)
	<-w.done
	close(w.ack)
	close(w.done)
}

func (w *WorkQueue) Complete() {
	w.done <- true
}

func (w *WorkQueue) PushWork(i interface{}) {
	w.queue <- i
}

func NewWorkQueue(task func(a interface{}, cmn interface{}), cmn interface{}, s int) *WorkQueue {

	// Set up the Work Queue.
	wq := &WorkQueue{}
	if s > 0 {
		wq.size = s
	} else {
		wq.size = runtime.NumCPU()
	}

	// Now set up the actual queues needed.
	wq.queue = make(chan interface{})
	wq.ack = make(chan bool)
	wq.done = make(chan bool)
	for i := 0; i < wq.size; i++ {
		go func() {
			for {
				v, ok := <-wq.queue
				if ok {
					task(v, cmn)
				} else {
					wq.ack <- true
					return
				}
			}
		}()
	}
	go func() {
		for i := 0; i < wq.size; i++ {
			<-wq.ack
		}
		wq.Complete()
	}()

	return wq
}

