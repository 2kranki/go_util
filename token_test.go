// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// Test Token

package util

import (
	"testing"
)

func TestLocationCreate(t *testing.T) {
	var loc    	*Location

	t.Logf("TestLocationCreate():\n")

	loc = NewLocation()
	if nil == loc {
		t.Error("Error: Could not allocate Location!")
	}

	t.Logf("...TestLocationCreate completed.\n")

}

func TestLocationJson(t *testing.T) {
	var err		error
	var loc    	*Location
	var loc2    *Location
	var bytes	[]byte

	t.Logf("TestLocationJson():\n")

	loc = NewLocation()
	if nil == loc {
		t.Error("Error: Could not allocate Location!")
	}
	loc.Path = "abc"
	loc.LineNo = 1
	loc.ColNo = 12
	bytes, err = loc.JsonMarshal()
	if err != nil {
		t.Errorf("Error: %s\n", err.Error())
	}
	t.Logf("json: %s\n", bytes)

	loc2 = NewLocation()
	if nil == loc2 {
		t.Error("Error: Could not allocate Location!")
	}
	err = loc2.JsonUnmarshal(bytes)
	if err != nil {
		t.Errorf("Error: %s\n", err.Error())
	}
	if *loc != *loc2 {
		t.Errorf("Error: loc2 != loc, loc2:%s\n", loc2)
	}

	t.Logf("...TestLocationJson completed.\n")

}

func TestTokenCreate(t *testing.T) {
    var tk      *Token

    t.Logf("TestTokenCreate():\n")

    tk = NewToken()
    if nil == tk {
		t.Error("Error: Could not allocate Token!")
	}

	t.Logf("...TestTokenCreate completed.\n")

}

func TestTokenJson(t *testing.T) {
	var err		error
	var tk1    	*Token
	var tk2    	*Token
	var bytes	[]byte

	t.Logf("TestTokenJson():\n")

	tk1 = NewToken()
	if nil == tk1 {
		t.Error("Error: Could not allocate Token!")
	}
	tk1.Loc.Path = "abc"
	tk1.Loc.LineNo = 1
	tk1.Loc.ColNo = 12
	tk1.Data = "def"
	bytes, err = tk1.JsonMarshal()
	if err != nil {
		t.Errorf("Error: %s\n", err.Error())
	}
	t.Logf("json: %s\n", bytes)

	tk2 = NewToken()
	if nil == tk2 {
		t.Error("Error: Could not allocate Token!")
	}
	err = tk2.JsonUnmarshal(bytes)
	if err != nil {
		t.Errorf("Error: %s\n", err.Error())
	}
	if *tk1 != *tk2 {
		t.Errorf("Error: tk2 != tk1, tk2:%s\n", tk2)
	}

	t.Logf("...TestTokenJson completed.\n")

}

