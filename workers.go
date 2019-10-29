// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// Work Queue Classes

// These classes allow work to be distributed to a limited number of goroutines.
// The primary design was for a main goroutine to create the work queue and
// populate it with data, then wait for the data to be processed. As the work
// in the queue is accomplished, the workers end. Right now, there is nothing
// to restart them if a partial shutdown occurs. This is set up for a situation
// where the queue is front loaded and then allowed to perform the needed work
// and end.  It is not designed for ongoing work. That would be a different
// design.

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
	queue		chan interface{}	// Work Queue - Each element is data that will be
	//								// handed to a worker goroutine.
	ack			chan bool			//
	done		chan bool			//
	size		int					// Number of Concurrent Workers allowed
}

// CloseQueue closes the Work Queue which has the side-effect that
// no more work will be allowed to be added.
func (w *WorkQueue) CloseQueue() {
	close(w.queue)
}

// CloseAndWaitForCompletion is normally called by the goroutine
// that created the Work Queue. It
func (w *WorkQueue) CloseAndWaitForCompletion() {
	close(w.queue)
	<-w.done			// Wait for completion.
	close(w.ack)
	close(w.done)
}

// Complete should only be called internally when all worker data
// has been completed.
func (w *WorkQueue) Complete() {
	w.done <- true
}

// PushWork is called by anyone to add more work to the Work Queue.
// If the Work Queue is full, then the goroutine that calls this
// method will be paused.
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

	// Set up the work goroutines per the size restrictions.
	for i := 0; i < wq.size; i++ {
		go func() {
			for {
				v, ok := <-wq.queue		// Get the next work data.
				if ok {
					task(v, cmn)		// Process the data.
				} else {				// Channel is closed and drained.
					wq.ack <- true		// Let this work goroutine end.
					return
				}
			}
		}()
	}

	// Set up another goroutine which calls Complete() when
	// all of the worker data has been completed.
	go func() {
		// Wait for all worker goroutines to end.
		for i := 0; i < wq.size; i++ {
			<-wq.ack
		}
		wq.Complete()
	}()

	return wq
}

