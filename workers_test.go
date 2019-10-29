// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// Test files package

package util

import (
	"fmt"
	"testing"
)


func TestWorkQueue(t *testing.T) {
	var work 	*WorkQueue
	var n		int

	t.Log("TestWorkQueue()")

	// Set up a work queue.
	work = NewWorkQueue(
			func(a interface{}, cmn interface{}) {
				var t		int
				t = a.(int)
				fmt.Printf("\t work %d\n", t)
			},
			nil,
			2)

	// Now push some work on it.
	// Note - this may cause this goroutine to pause
	// if the work queue becomes full.
	for n = 1; n < 10; n++ {
		work.PushWork(n)
	}

	// Wait until all the work is complete.
	work.CloseAndWaitForCompletion()

	t.Log("\tend: TestWorkQueue")
}

