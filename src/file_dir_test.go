// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// Test files package

package util

import (
	"os/exec"
	"testing"
)

type jsonData struct {
	Debug   bool   `json:"debug,omitempty"`
	Force   bool   `json:"force,omitempty"`
	Noop    bool   `json:"noop,omitempty"`
	Quiet   bool   `json:"quiet,omitempty"`
	Cmd     string `json:"cmd,omitempty"`
	Defines string `json:"defines,omitempty"`
	Outdir  string `json:"outdir,omitempty"`
}

func TestFileCompare(t *testing.T) {

	t.Log("TestFileCompare()")

	src := NewPath("./exec.go")
	dst := NewPath("./exec.go")
	if !FileCompareEqual(src, dst) {
		t.Errorf("FileCompare(%s,%s) failed comparison\n", src, dst)
	}

	src = NewPath("./misc.go")
	dst = NewPath("./exec.go")
	if FileCompareEqual(src, dst) {
		t.Errorf("FileCompare(%s,%s) failed comparison\n", src, dst)
	}

	t.Log("\tend: TestFileCompare")
}

func TestCopyFile(t *testing.T) {
	var err 	error

	t.Log("TestCopyFile()")

	src := NewPath("../test").Append("test.exec.json.txt")
	dst := NewTempDir().Append("testout.txt")
	err = CopyFile(src, dst)
	if err != nil {
		t.Errorf("CopyFile(%s,%s) failed: %s\n", src.String(), dst.String(), err.Error())
	}

	if !FileCompareEqual(src,dst) {
		t.Errorf("CopyFile(%s,%s) failed comparison\n", src.String(), dst.String())
	}

	err = dst.DeleteFile()
	if err != nil {
		t.Errorf("DeleteFile(%s) failed: %s\n", dst.String(), err.Error())
	}

	t.Log("\tend: TestCopyFile")
}

func TestCopyDir(t *testing.T) {
	var err error

	t.Log("TestCopyDir()")

	src  := NewPath("../test")
	dst  := NewTempDir().Append("test2")
	dst2 := NewTempDir().Append("test3")

	err = dst.RemoveDir()
	if err != nil {
		t.Logf("\tError: Deleting %s: %s\n", dst.String(), err.Error())
	}
	err = dst2.RemoveDir()
	if err != nil {
		t.Logf("\tError: Deleting %s: %s\n", dst2.String(), err.Error())
	}

	t.Logf("\tcopying %s -> %s\n", src, dst)
	err = CopyDir(src, dst)
	if err != nil {
		t.Fatalf("CopyDir(%s,%s) failed: %s\n", src.String(), dst.String(), err.Error())
	}

	cmd := exec.Command("diff", src.Absolute(), dst.Absolute())
	err = cmd.Run()
	if err != nil {
		t.Fatalf("CopyDir(%s,%s) comparison failed: %s\n", src, dst, err)
	}

	dst.RemoveDir()

	dst3 := dst2.Append("../test")
	dst2 =  dst2.Append("")
	t.Logf("\tcopying %s -> %s\n", src.String(), dst2.String())
	err = CopyDir(src, dst3)
	if err != nil {
		t.Fatalf("CopyDir(%s,%s) failed: %s\n", src.String(), dst3.String(), err.Error())
	}

	cmd = exec.Command("diff", src.Absolute(), dst3.Absolute())
	err = cmd.Run()
	if err != nil {
		t.Fatalf("CopyDir(%s,%s) comparison failed: %s\n", src.String(), dst3.String(), err)
	}

	dst2.RemoveDir()

	t.Log("\tend: TestCopyDir")
}

func TestReadJson(t *testing.T) {
	var jsonOut interface{}
	var wrk interface{}
	var err error

	t.Log("TestReadJson()")

	if jsonOut, err = ReadJsonFile("../test/test.exec.json.txt"); err != nil {
		t.Errorf("ReadJson(test.exec.json.txt) failed: %s\n", err)
	}
	m := jsonOut.(map[string]interface{})
	if wrk = m["debug"]; wrk == nil {
		t.Errorf("ReadJson(test.exec.json.txt) missing 'debug'\n")
	}
	if wrk = m["debug_not_there"]; wrk != nil {
		t.Errorf("ReadJson(test.exec.json.txt) missing 'debug'\n")
	}
	wrk = m["cmd"]
	if wrk.(string) != "sqlapp" {
		t.Errorf("ReadJson(test.exec.json.txt) missing 'cmd'\n")
	}

	t.Log("\tend: TestReadJson")
}

func TestReadJsonFileToData(t *testing.T) {
	var jsonOut = jsonData{}
	var err error

	t.Log("TestReadJsonFileToData()")

	jsonOut = jsonData{}
	t.Log("&jsonOut:", &jsonOut)
	err = ReadJsonFileToData("../test/test.exec.json.txt", &jsonOut)
	if err != nil {
		t.Errorf("ReadJsonToData(test.exec.json.txt) failed: %s\n", err)
	}
	t.Log("test jsonOut:", jsonOut)
	if jsonOut.Cmd != "sqlapp" {
		t.Errorf("ReadJsonToData(test.exec.json.txt) missing or invalid 'cmd'\n")
	}
	if jsonOut.Outdir != "./test" {
		t.Errorf("ReadJson(test.exec.json.txt) missing or invalid 'outdir'\n")
	}
	t.Log("\tend: TestReadJsonToData")
}


