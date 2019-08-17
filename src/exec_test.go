// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// Test files package

package util

import (
	"testing"
)

func TestExecCmd(t *testing.T) {
	var err 	error
	var cmd		*ExecCmd

	t.Log("TestExecCmd()")

	cmd = NewExecCmd("")
	if cmd == nil {
		t.Errorf("NewExecCmd(\"\") failed to allocate!\n")
	}

	cmd = NewExecCmd("ls")
	if cmd == nil {
		t.Errorf("NewExecCmd(\"ls\") failed to allocate!\n")
	}
	out, err := cmd.RunWithOutput()
	if err != nil {
		t.Errorf("RunWithOutput(\"ls\") failed: %s\n", err.Error())
	}
	t.Logf("\tls output: %s", out)

	t.Log("\tend: TestExecCmd")
}

