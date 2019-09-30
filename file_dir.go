// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// Miscellaneous utility functions

// Note that these functions are made unique by attaching them to a structure
// which must be instanced before they can be used.

// Some functions were taken from https://blog.kowalczyk.info/book/go-cookbook.html
// which was declared public domain at the time that the functions were taken.

package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/2kranki/jsonpreprocess"
)

//============================================================================
//                        File and Directory Functions 
//============================================================================

//----------------------------------------------------------------------------
//                             CopyDir
//----------------------------------------------------------------------------

// CopyDir copies from the given directory (src) and all of its files to the
// destination (dst).
func CopyDir(src, dst *Path) error {
	var err 	error

	//log.Printf("CopyDir: base: %s  last: %c\n", pathIn.Base(), dst[len(dst)-1])
	if dst.String()[len(dst.String())-1] == os.PathSeparator {
		dst = dst.Append(src.Base())
	}
	//log.Printf("CopyDir: %s -> %s\n", pathIn.String(), pathOut.String())

	if src.IsPathRegularFile() {
		return CopyFile(src, dst)
	}

	if !src.IsPathDir() {
		return fmt.Errorf("Error: CopyDir: %s is not a file or directory!\n", src.String())
	}

	si, err := os.Stat(src.Absolute())
	if err != nil {
		return err
	}
	mode := si.Mode() & 03777

	log.Printf("CopyDir: MkdirAll %s %o\n", dst.Absolute(), mode)
	err = os.MkdirAll(dst.Absolute(), mode)
	if err != nil {
		return err
	}
	if !dst.IsPathDir() {
		return fmt.Errorf("Error: %s could not be found!", dst.Absolute())
	}

	dir, err := os.Open(src.Absolute())
	if err != nil {
		return err
	}

	entries, err := dir.Readdir(-1)
	if err != nil {
		dir.Close()
		return err
	}
	dir.Close()

	for _, fi := range entries {
		srcNew := src.Append(fi.Name())
		dstNew := dst.Append(fi.Name())

		if fi.Mode().IsDir() {
			log.Printf("CopyDir: Dir: %s -> %s\n", srcNew.String(), dstNew.String())
			err = CopyDir(srcNew, dstNew)
			if err != nil {
				return err
			}
		} else if fi.Mode().IsRegular() {
			log.Printf("CopyDir: File: %s -> %s\n", srcNew.String(), dstNew.String())
			err = CopyFile(srcNew, dstNew)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

//----------------------------------------------------------------------------
//                             CopyFile
//----------------------------------------------------------------------------

// CopyFile copies a file given by its path (src) creating
// an output file given its path (dst)
func CopyFile(src, dst *Path) error {
	var err 	error

	log.Printf("CopyFile: %s -> %s\n", src.Absolute(), dst.Absolute())

	// Clean up the input file path and check for its existence.
	if !src.IsPathRegularFile() {
		return fmt.Errorf("Error: %s is not a file!\n", src.String())
	}

	// Open the input file.
	fileIn, err := os.Open(src.Absolute())
	if err != nil {
		return err
	}
	defer fileIn.Close()

	// Create the output file.
	fileOut, err := os.Create(dst.Absolute())
	if err != nil {
		return err
	}
	defer fileOut.Close()

	// Perform the copy and set the output file's privileges
	// to same as input file.
	_, err = io.Copy(fileOut, fileIn)
	if err == nil {
		si, err := os.Stat(src.Absolute())
		if err == nil {
			err = os.Chmod(dst.Absolute(), si.Mode())
		}
	}

	return err
}

//----------------------------------------------------------------------------
//                             FileCompare
//----------------------------------------------------------------------------

// FileCompareEqual compares two files returning true
// if they are equal.
func FileCompareEqual(file1, file2 *Path) bool {
	var err 		error

	if !file1.IsPathRegularFile() {
		return false
	}

	if !file2.IsPathRegularFile() {
		return false
	}

	if file1.Size() != file2.Size() {
		return false
	}

	f1, err := os.Open(file1.Absolute())
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()

	f2, err := os.Open(file2.Absolute())
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	b1 := make([]byte, 8192)
	b2 := make([]byte, 8192)
	for {
		c1, err1 := f1.Read(b1)
		c2, err2 := f2.Read(b2)
		if c1 != c2 {
			return false
		}

		if err1 != nil || err2 != nil {
			if err1 == io.EOF && err2 == io.EOF {
				return true
			} else if err1 == io.EOF || err2 == io.EOF {
				return false
			} else {
				log.Fatal(err1, err2)
			}
		}

		if !bytes.Equal(b1, b2) {
			return false
		}
	}

	return false
}

//----------------------------------------------------------------------------
//                            ReadJsonFile
//----------------------------------------------------------------------------

// ReadJsonFile preprocesses out comments and then unmarshals the data
// generically.
func ReadJsonFile(jsonPath string) (interface{}, error) {
	var err error
	var jsonOut interface{}

	// Open the input template file
	input, err := os.Open(jsonPath)
	if err != nil {
		return jsonOut, err
	}
	textBuf := strings.Builder{}
	err = jsonpreprocess.WriteMinifiedTo(&textBuf, input)
	if err != nil {
		return jsonOut, err
	}

	// Read and process the JSON.
	err = json.Unmarshal([]byte(textBuf.String()), &jsonOut)
	if err != nil {
		return jsonOut, err
	}

	return jsonOut, err
}

//----------------------------------------------------------------------------
//                            ReadJsonFileToData
//----------------------------------------------------------------------------

// ReadJsonFileToData preprocesses out comments and then unmarshals the data
// into a data structure previously defined.
func ReadJsonFileToData(jsonPath string, jsonOut interface{}) error {
	var err error

	// Open the input template file
	input, err := os.Open(jsonPath)
	if err != nil {
		return err
	}
	textBuf := strings.Builder{}
	err = jsonpreprocess.WriteMinifiedTo(&textBuf, input)
	if err != nil {
		return err
	}

	// Read and process the template file
	err = json.Unmarshal([]byte(textBuf.String()), jsonOut)
	if err != nil {
		return err
	}

	return err
}


