// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// Test files package

package util

import (
	"testing"
)


func TestSharedData(t *testing.T) {
	var sd 		*SharedData
	var b		bool

	t.Log("TestSharedData()")
	sd = &SharedData{}
	sd.Init()
	b = sd.Force()
	if b {
		t.Errorf("Force() should be false, but is true!\n")
	}
	t.Log("\tend: TestSharedData")
}

