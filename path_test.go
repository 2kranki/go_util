// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// Test files package

package util

import (
	"os"
	"testing"
)

func TestDir(t *testing.T) {
	var test 	func (string,string)

	t.Log("TestDir()")
	test = func(src, dst string) {
		p := NewPath(src)
		if p.Dir() != dst {
			t.Errorf("Dir(%s) = %s, but should be %s!\n", src, p.Dir(), dst)
		}
	}

	test("/a/b/c.d", "/a/b")
	test("/a/b/c", "/a/b")
	test("a/b/c", "a/b")
	test("/a/", "/a")
	test("a/", "a")
	test("/", "/")
	test("..", "..")
	test("", ".")

	t.Log("\tend: TestDir")
}

func TestIsPathDir(t *testing.T) {
	var path	*Path

	t.Log("TestIsPathDir()")

	path = NewPath("./exec.go")
	if path.IsPathDir() {
		t.Errorf("IsPathDir(%s) failed!\n", path.String())
	}
	t.Logf("\t%s absolute: %s\n", path.String(), path.Absolute())

	path = NewPath("./test")
	if !path.IsPathDir() {
		t.Errorf("IsPathDir(%s) failed!\n", path.String())
	}
	t.Logf("\t%s absolute: %s\n", path.String(), path.Absolute())

	t.Log("\tend: TestIsPathDir")
}

func TestIsPathRegularFile(t *testing.T) {
	var path 	*Path
	var err 	error

	t.Log("TestIsPathRegularFile()")

	path = NewPath("./exec.go")
	if !path.IsPathRegularFile() {
		t.Errorf("IsPathRegularFile(%s) failed: %s\n", path.String(), err.Error())
	}
	t.Logf("\t%s Absolute: %s\n", path.String(), path.Absolute())

	path = NewPath("./xyzzy.go")
	if path.IsPathRegularFile() {
		t.Errorf("IsPathRegularFile(%s) failed: %s\n", path.String(), err.Error())
	}
	t.Logf("\t%s Absolute: %s\n", path.String(), path.Absolute())

	t.Log("\tend: TestIsPathRegularFile")
}

func TestPath(t *testing.T) {
	var err			error
	var expected	string
	var input		string
	var path 		*Path
	var pth			string
	homeDir := NewHomeDir()
	curDir, err := os.Getwd()
	if err != nil {
		t.Errorf("Error: Getting Current Directory: %s\n", err)
	}

	t.Log("TestPathClean()")

	input = "./exec.go"
	expected = curDir + "/exec.go"
	path = NewPath(input)
	pth = path.Clean()
	t.Logf("\t%s => %s\n", input, pth)
	if pth != expected {
		t.Errorf("PathClean Got: %s  Expected: %s\n", pth, expected)
	}

	input = "./xyzzy.go"
	expected = curDir + "/xyzzy.go"
	path = NewPath(input)
	pth = path.Clean()
	t.Logf("\t%s => %s\n", input, pth)
	if pth != expected {
		t.Errorf("PathClean Got: %s  Expected: %s\n", pth, expected)
	}

	input = "~"
	expected = homeDir.String()
	path = NewPath(input)
	path.Clean()
	pth = path.Absolute()
	t.Logf("\t%s => %s\n", input, pth)
	if pth != expected {
		t.Errorf("PathClean Got: %s  Expected: %s\n", pth, expected)
	}

	input = "~/.ssh"
	expected = homeDir.String() + "/.ssh"
	path = NewPath(input)
	pth = path.Clean()
	t.Logf("\t%s => %s\n", input, pth)
	if pth != expected {
		t.Errorf("PathClean Got: %s  Expected: %s\n", pth, expected)
	}

	input = "./test3"
	path = NewPath(input)
	if err = path.CreateDir(); err != nil {
		t.Fatalf("FATAL: create ./test3 failed: %s\n", err.Error())
	}
	if !path.IsPathDir() {
		t.Fatalf("FATAL: create ./test3 failed!\n")
	}
	if err = path.RemoveDir(); err != nil {
		t.Fatalf("FATAL: remove ./test3 failed: %s\n", err.Error())
	}
	if path.IsPathDir() {
		t.Fatalf("FATAL: remove ./test3 failed!\n")
	}

	t.Logf("\t%s => %s\n", input, pth)
	if pth != expected {
		t.Errorf("PathClean Got: %s  Expected: %s\n", pth, expected)
	}

	pwd := NewCurrentWorkDir()
	t.Logf("PWD: %s\n", pwd.String())

	t.Log("\tend: TestPathClean")
}


