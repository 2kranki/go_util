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

	work = NewWorkQueue(
			func(a interface{}, cmn interface{}) {
				var t		int
				t = a.(int)
				fmt.Printf("\t work %d\n", t)
			},
			nil,
			2)
	for n = 1; n < 10; n++ {
		work.PushWork(n)
	}
	work.CloseAndWaitForCompletion()

	t.Log("\tend: TestWorkQueue")
}

