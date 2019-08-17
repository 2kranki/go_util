// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// Miscellaneous utility functions

// Note that these functions are made unique by attaching them to a structure
// which must be instanced before they can be used.

// Some functions were taken from https://blog.kowalczyk.info/book/go-cookbook.html
// which was declared public domain at the time that the functions were taken.

package util

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

//============================================================================
//                             		Path
//============================================================================

// Path provides a centralized
type Path struct {
	str       	string
}

// Absolute returns the absolute file path for
// this path.
func (p *Path) Absolute( ) string {
	path := p.Clean()
	path, _ = filepath.Abs(path)
	return path
}

// Append a subdirectory or file name[.file extension] to the path.
// If the string is empty, then '/' will be appended.
func (p *Path) Append(s string) *Path {
	pth := Path{}
	pth.str = p.str + string(os.PathSeparator) + s
	pth.str = filepath.Clean(pth.str)
	return &pth
}

// Base returns the last component of the path. If the
// path is empty, "." is returned.
func (p *Path) Base( ) string {
	b := filepath.Base(p.str)
	return b
}

// Chmod changes the mode of the named file to the given mode.
func (p *Path) Chmod(mode os.FileMode) error {
	var err error

	b := p.Clean()
	if len(b) > 0 {
		err = os.Chmod(b, mode)
	}

	return err
}

// Clean cleans up the file path. It returns the absolute
// file path if needed.
func (p *Path) Clean( ) string {
	var path string

	if strings.HasPrefix(p.str, "~") {
		p.str = NewHomeDir().String() + string(os.PathSeparator) + p.str[1:]
	}
	p.str = os.ExpandEnv(p.str)
	p.str = filepath.Clean(p.str)
	path, _ = filepath.Abs(p.str)

	return path
}

// Copy creates a new copy of the path.
func (p *Path) Copy( ) *Path {
	pth := Path{}
	pth.str = p.str
	return &pth
}

// CreateDir assumes that this path represents a
// directory and creates it along with any parent
// directories needed as well.
func (p *Path) CreateDir( ) error {
	var err error

	b := p.Clean()
	if len(b) > 0 {
		err = os.MkdirAll(b, 0777)
	}

	return err
}

// DeleteFile assumes that this path represents a
// file and deletes it.
func (p *Path) DeleteFile( ) error {
	var err error

	pth := p.Clean()
	if len(pth) > 0 {
		fi, err := os.Lstat(pth)
		if err != nil {
			err = fmt.Errorf("Error: DeleteFile(): %s is not a file!\n", pth)
		} else {
			if fi.Mode().IsRegular() {
				err = os.Remove(pth)
			} else {
				err = fmt.Errorf("Error: DeleteFile(): %s is not a file!\n", pth)
			}
		}
	}

	return err
}

// Expand replaces ${var} or $var in the given path based on the
// mapping function returning a new path.
func (p *Path) Expand(mapping func(string) string) *Path {
	pth := &Path{}
	pth.str = os.Expand(p.str, mapping)
	pth.str = filepath.Clean(pth.str)
	return pth
}

// IsPathDir cleans up the supplied file path
// and then checks the cleaned file path to see
// if it is an existing standard directory.
func (p *Path) IsPathDir( ) bool {
	var err error
	var pth string

	pth = p.Clean( )
	fi, err := os.Lstat(pth)
	if err != nil {
		return false
	}
	if fi.Mode().IsDir() {
		return true
	}
	return false
}

// IsPathRegularFile cleans up the supplied file path
// and then checks the cleaned file path to see
// if it is an existing standard file.
func (p *Path) IsPathRegularFile( ) bool {
	var err error
	var pth string

	pth = p.Clean()
	fi, err := os.Lstat(pth)
	if err != nil {
		return false
	}
	if fi.Mode().IsRegular() {
		return true
	}
	return false
}

func (p *Path) Mode( ) os.FileMode {
	var mode	os.FileMode

	si, err := os.Stat(p.Absolute())
	if err == nil {
		mode = si.Mode()
	}
	return mode
}

func (p *Path) ModTime( ) time.Time {
	var mod		time.Time

	si, err := os.Stat(p.Absolute())
	if err == nil {
		mod = si.ModTime()
	}
	return mod
}

// RemoveDir assumes that this path represents a
// directory and deletes it along with any parent
// directories that it can as well.
func (p *Path) RemoveDir( ) error {
	var err error

	b := p.Clean()
	if len(b) > 0 {
		err = os.RemoveAll(b)
	}

	return err
}

func (p *Path) SetStr(s string) {
	p.str = s
}

// Size returns length in bytes for regular files.
func (p *Path) Size( ) int64 {
	var size	int64

	si, err := os.Stat(p.Absolute())
	if err == nil {
		size = si.Size()
	}
	return size
}

func (p *Path) String( ) string {
	return p.str
}

// NewHomeDir returns the current working directory as a Path.
func NewHomeDir() *Path {
	p := Path{}

	// user.Current() returns nil if cross-compiled e.g. on mac for linux.
	// So, our backup is to get the home directory from the environment.
	if usr, _ := user.Current(); usr != nil {
		p.str = usr.HomeDir
	} else {
		p.str = os.Getenv("HOME")
	}

	return &p
}

func NewPath(s string) *Path {
	p := Path{}
	p.str = s
	return &p
}

// NewWorkDir returns the current working directory as a Path.
func NewWorkDir() *Path {
	p := Path{}
	p.str, _ = os.Getwd()
	return &p
}

// NewTempDir returns the temporary directory as a Path.
func NewTempDir() *Path {
	p := Path{}
	p.str = os.TempDir()
	return &p
}


