// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// Miscellaneous utility functions

// Note: I originally did this to get WriteStringf. However, the following
// 		will work without the need for StringBuilder:\
//			var str strings.Builder
//			fmt.Fprintf(&str, "Hello, world!\n")

// Note that these functions are made unique by attaching them to a structure
// which must be instanced before they can be used.

// Some functions were taken from https://blog.kowalczyk.info/book/go-cookbook.html
// which was declared public domain at the time that the functions were taken.

package util

import (
	"fmt"
	"strings"
)

//============================================================================
//								String Builder
//============================================================================

// StringBuilder is a composition of strings.Builder so that
// we can add supplemental functions such as formatted strings
// easily.
type StringBuilder struct {
	str		strings.Builder
}

func NewStringBuilder() *StringBuilder {
	sb := StringBuilder{}
	return &sb
}

func (s *StringBuilder) Len( ) int {
	return s.str.Len()
}

func (s *StringBuilder) String( ) string {
	return s.str.String()
}

// WriteString allows us to write a string to the buffer.
func (s *StringBuilder) WriteString(format string) error {
	_, err := s.str.WriteString(format)
	return err
}

// WriteStringf allows us to write a formatted string.
func (s *StringBuilder) WriteStringf(format string, a ...interface{}) error {
	str := fmt.Sprintf(format, a...)
	err := s.WriteString(str)
	return err
}


