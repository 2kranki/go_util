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
	"time"
)

//============================================================================
//                          	Shared Data
//============================================================================

// SharedData represents a common structure to share data across packages
// easily. It would normally be created in package, main, and passed as
// needed to other packages.
type SharedData struct {
	cmd			string
	dataPath	string
	defns		map[string]interface{}
	funcs		map[string]interface{}
	mainPath	string
	outDir		string
}

func (s *SharedData) Init() {
	s.defns  = map[string]interface{}{}
	s.funcs  = map[string]interface{}{}
	s.outDir = "/tmp"
	s.defns["Debug"] = false
	s.defns["Force"] = false
	s.defns["Noop"] = false
	s.defns["Quiet"] = false
	s.defns["Replace"] = true
	s.defns["Time"] = time.Now().Format("Mon Jan _2, 2006 15:04")
	s.funcs["Time"] = s.Time
}

func (s *SharedData) Cmd() string {
	return s.cmd
}

func (s *SharedData) SetCmd(f string) {
	s.cmd = f
}

// DataPath is the path to the app json file.
func (s *SharedData) DataPath() string {
	return s.dataPath
}

func (s *SharedData) SetDataPath(f string) {
	s.dataPath = f
}

func (s *SharedData) Debug() bool {
	return s.defns["Debug"].(bool)
}

func (s *SharedData) SetDebug(f bool) {
	s.defns["Debug"] = f
}

func (s *SharedData) Defn(nm string) interface{} {
	switch nm {
	case "cmd":
		return s.cmd
	case "dataPath":
		return s.dataPath
	case "mainPath":
		return s.mainPath
	case "outDir":
		return s.outDir
	}
	defn, ok := s.defns[nm]
	if ok {
		return defn
	}
	return nil
}

func (s *SharedData) IsDefined(nm string) bool {
	x := s.Defn(nm)
	if x != nil {
		return true
	}
	return false
}

func (s *SharedData) SetDefn(nm string, d interface{}) {
	var ok		bool
	var sw		bool
	var str		string

	switch nm {
	case "cmd":
		if str, ok = d.(string); ok {
			s.cmd = str
		}
	case "dataPath":
		if str, ok = d.(string); ok {
			s.dataPath = str
		}
	case "Debug":
		if sw, ok = d.(bool); ok {
			s.defns["Debug"] = sw
		}
	case "Force":
		if sw, ok = d.(bool); ok {
			s.defns["Force"] = sw
		}
	case "mainPath":
		if str, ok = d.(string); ok {
			s.mainPath = str
		}
	case "Noop":
		if sw, ok = d.(bool); ok {
			s.defns["Noop"] = sw
		}
	case "outDir":
		if str, ok = d.(string); ok {
			s.outDir = str
		}
	case "Quiet":
		if sw, ok = d.(bool); ok {
			s.defns["Quiet"] = sw
		}
	case "Replace":
		if sw, ok = d.(bool); ok {
			s.defns["Replace"] = sw
		}
	case "Time":
		if str, ok = d.(string); ok {
			s.defns["Time"] = str
		}
	default:
		s.defns[nm] = d
	}
}

func (s *SharedData) Force() bool {
	return s.defns["Force"].(bool)
}

func (s *SharedData) SetForce(f bool) {
	s.defns["Force"] = f
}

func (s *SharedData) Funcs() map[string]interface{} {
	return s.funcs
}

func (s *SharedData) FuncsSlice() []interface{} {
	var f = []interface{}{}

	for _, v := range s.funcs {
		f = append(f, v)
	}

	return f
}

func (s *SharedData) SetFunc(nm string, d interface{}) {
	s.funcs[nm] = d
}

func (s *SharedData) GenDebugging() bool {
	return s.defns["GenDebugging"].(bool)
}

func (s *SharedData) GenLogging() bool {
	return s.defns["GenLogging"].(bool)
}

func (s *SharedData) GenMuxWrapper() bool {
	return s.defns["GenMuxWrapper"].(bool)
}

// MainPath is the path to the main json file.
func (s *SharedData) MainPath() string {
	return s.mainPath
}

func (s *SharedData) SetMainPath(f string) {
	s.mainPath = f
}

// MergeFrom merges the given map into the shared
// data definitions optionally replacing any that
// already exist.
func (s *SharedData) MergeFrom(m map[string]interface{}, rep bool) {
	var ok			bool

	for k, v := range m {
		if rep {
			s.defns[k] = v
		} else {
			if _, ok = s.defns[k]; !ok {
				s.defns[k] = v
			}
		}
	}
}

// MergeTo merges the shared into the given map
// optionally replacing any in the given map.\
func (s *SharedData) MergeTo(m map[string]interface{}, rep bool) {
	var ok			bool

	for k, v := range s.defns {
		if rep {
			m[k] = v
		} else {
			if _, ok = m[k]; !ok {
				m[k] = v
			}
		}
	}
}

func (s *SharedData) Noop() bool {
	return s.defns["Noop"].(bool)
}

func (s *SharedData) SetNoop(f bool) {
	s.defns["Noop"] = f
}

func (s *SharedData) OutDir() string {
	return s.outDir
}

func (s *SharedData) SetOutDir(f string) {
	s.outDir = f
}

func (s *SharedData) Quiet() bool {
	return s.defns["Quiet"].(bool)
}

func (s *SharedData) SetQuiet(f bool) {
	s.defns["Quiet"] = f
}

func (s *SharedData) Replace() bool {
	return s.defns["Replace"].(bool)
}

func (s *SharedData) SetReplace(f bool) {
	s.defns["Replace"] = f
}

// String returns a stringified version of the shared data
func (s *SharedData) String() string {
	str := "{"
	str += fmt.Sprintf("cmd:%q,",s.cmd)
	str += fmt.Sprintf("dataPath:%q,",s.dataPath)
	str += fmt.Sprintf("Debug:%v,",s.defns["Debug"])
	str += fmt.Sprintf("Force:%v,",s.defns["Force"])
	str += fmt.Sprintf("mainPath:%q,",s.mainPath)
	str += fmt.Sprintf("Noop:%v,",s.defns["Noop"])
	str += fmt.Sprintf("outDir:%q,",s.outDir)
	str += fmt.Sprintf("Quiet:%v,",s.defns["Quiet"])
	str += fmt.Sprintf("Time:%q,",s.defns["Time"])
	str += "}"
	return str
}

// MainPath is the path to the main json file.
func (s *SharedData) Time() string {
	return s.defns["Time"].(string)
}

func (s *SharedData) SetTime(f string) {
	s.defns["Time"] = f
}

