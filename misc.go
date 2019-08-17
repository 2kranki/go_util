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
)

//============================================================================
//                             Miscellaneous
//============================================================================

//----------------------------------------------------------------------------
//                             ErrorString
//----------------------------------------------------------------------------

func ErrorString(err error) string {
	if err == nil {
		return "ok"
	} else {
		return err.Error()
	}
}

//----------------------------------------------------------------------------
//                             		FormatArgs
//----------------------------------------------------------------------------

func FormatArgs(args ...interface{}) string {
	if len(args) == 0 {
		return ""
	}
	format := args[0].(string)
	if len(args) == 1 {
		return format
	}
	return fmt.Sprintf(format, args[1:]...)
}

//----------------------------------------------------------------------------
//                             PanicIf
//----------------------------------------------------------------------------

func PanicIf(cond bool, args ...interface{}) {
	if !cond {
		return
	}
	s := FormatArgs(args...)
	if s == "" {
		s = "fatalIf: cond is false"
	}
	panic(s)
}

//----------------------------------------------------------------------------
//                             PanicIfErr
//----------------------------------------------------------------------------

func PanicIfErr(err error, args ...interface{}) {
	if err == nil {
		return
	}
	s := FormatArgs(args...)
	if s == "" {
		s = err.Error()
	}
	panic(s)
}


