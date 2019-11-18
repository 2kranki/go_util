// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// Test files package

package util

import (
	"log"
	"testing"
)

func TestExecArgs(t *testing.T) {
	var err 	error
	var cmd		*ExecCmd

	t.Log("TestExecCmd()")

	cmd = NewExecArgs("")
	if cmd == nil {
		t.Errorf("NewExecCmd(\"\") failed to allocate!\n")
	}

	cmd = NewExecArgs("ls")
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

func TestParseCommand(t *testing.T) {
	var err 	error
	var args	[]string
	cmd01 := "./program_name --debug -v 1st_Arg 2nd_Arg"
	cmd02 := "./program_name --debug='true' -v 1st_Arg \"2nd_Arg\""

	t.Log("TestParseCommand01()")

	args, err = ParseCommandLine("")
	if err != nil {
		t.Fatalf("ParseCommand(\"\") error: %s!\n", err.Error())
	}
	if len(args) != 0 {
		t.Fatalf("ParseCommand(\"\") error: len: %d!\n", len(args))
	}

	log.Printf("%s\n", cmd01)
	log.Printf("          1         2         3         4\n")
	log.Printf("01234567890123456789012345678901234567890123456789\n")
	log.Printf("cmd len = %d\n", len(cmd01))
	args, err = ParseCommandLine(cmd01)
	if err != nil {
		t.Fatalf("ParseCommand(cmd01) error: %s!\n", err.Error())
	}
	t.Logf("\treturned args: %v", args)
	if len(args) != 5 {
		t.Fatalf("ParseCommand(cmd01) error: len: %d!\n", len(args))
	}
	if args[0] != "./program_name" {
		t.Errorf("ParseCommand(cmd01) arg0: %s!\n", args[0])
	}
	if args[1] != "--debug" {
		t.Errorf("ParseCommand(cmd01) arg1: %s!\n", args[1])
	}
	if args[2] != "-v" {
		t.Errorf("ParseCommand(cmd01) arg2: %s!\n", args[2])
	}
	if args[3] != "1st_Arg" {
		t.Errorf("ParseCommand(cmd01) arg3: %s!\n", args[3])
	}
	if args[4] != "2nd_Arg" {
		t.Errorf("ParseCommand(cmd01) arg4: %s!\n", args[4])
	}

	log.Printf("%s\n", cmd02)
	log.Printf("          1         2         3         4\n")
	log.Printf("01234567890123456789012345678901234567890123456789\n")
	log.Printf("cmd len = %d\n", len(cmd02))
	args, err = ParseCommandLine(cmd02)
	if err != nil {
		t.Fatalf("ParseCommand(cmd01) error: %s!\n", err.Error())
	}
	t.Logf("\treturned args: %v", args)
	if len(args) != 5 {
		t.Fatalf("ParseCommand(cmd01) error: len: %d!\n", len(args))
	}
	if args[0] != "./program_name" {
		t.Errorf("ParseCommand(cmd01) arg0: %s!\n", args[0])
	}
	if args[1] != "--debug='true'" {
		t.Errorf("ParseCommand(cmd01) arg1: %s!\n", args[1])
	}
	if args[2] != "-v" {
		t.Errorf("ParseCommand(cmd01) arg2: %s!\n", args[2])
	}
	if args[3] != "1st_Arg" {
		t.Errorf("ParseCommand(cmd01) arg3: %s!\n", args[3])
	}
	if args[4] != "\"2nd_Arg\"" {
		t.Errorf("ParseCommand(cmd01) arg4: %s!\n", args[4])
	}

	t.Log("\tend: TestParseCommand01")
}

