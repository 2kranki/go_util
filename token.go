// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// Miscellaneous utility functions



package util

import (
    "encoding/json"
	"fmt"
)


//============================================================================
//                  		    Token Types
//============================================================================

// Valid Types for Token:TypeNo
const (
    TypeFloat = iota
    TypeIdentifier
    TypeInteger
    TypeUserStart                   // First Type for Inheritance
)


//============================================================================
//                  		Position Object
//============================================================================

type Location struct {
    Path        string              `json:"path,omitempty"`
    Pos         int                 `json:"pos,omitempty"`
    LineNo      int                 `json:"line_no,omitempty"`
    ColNo       int                 `json:"col_no,omitempty"`
}

//----------------------------------------------------------------------------
//                  		JSON Marshal
//----------------------------------------------------------------------------

func (l *Location) JsonMarshal() ([]byte, error) {
    var err         error
    var text        []byte

    if text, err = json.Marshal(l); err != nil {
        return nil, fmt.Errorf("Error: marshalling json: %s : %v", err, l)
    }

    return text, err
}

//----------------------------------------------------------------------------
//                             JSON Unmarshal
//----------------------------------------------------------------------------

func (l *Location) JsonUnmarshal(text []byte) error {
    var err         error

    if err = json.Unmarshal(text, l); err != nil {
        return fmt.Errorf("Error: unmarshalling json: %s : %s", err, text)
    }

    return err
}

//----------------------------------------------------------------------------
//                             String
//----------------------------------------------------------------------------

func (l *Location) String() string {
    return fmt.Sprintf("%s: %d %d %d", l.Path, l.Pos, l.LineNo, l.ColNo)
}

//----------------------------------------------------------------------------
//                                  New
//----------------------------------------------------------------------------

func NewLocation() *Location {
    return &Location{}
}


//============================================================================
//                  		Token Object
//============================================================================

type Token struct {
    TypeNo      int                 `json:"type_no,omitempty"`
    ClassNo     int                 `json:"class_no,omitempty"`
    Loc         Location            `json:"loc,omitempty"`
    Data        string              `json:"data,omitempty"`
}


//----------------------------------------------------------------------------
//                  		JSON Marshal
//----------------------------------------------------------------------------

func (t *Token) JsonMarshal() ([]byte, error) {
    var err         error
    var text        []byte

    if text, err = json.Marshal(t); err != nil {
        return nil, fmt.Errorf("Error: marshalling json: %s : %v", err, t)
    }

    return text, err
}

//----------------------------------------------------------------------------
//                             JSON Unmarshal
//----------------------------------------------------------------------------

func (t *Token) JsonUnmarshal(text []byte) error {
    var err         error

    if err = json.Unmarshal(text, t); err != nil {
        return fmt.Errorf("Error: unmarshalling json: %s : %s", err, text)
    }

    return err
}

//----------------------------------------------------------------------------
//                             String
//----------------------------------------------------------------------------

func (t *Token) String() string {
    return fmt.Sprintf("%d: %d: %s: %d %d %d - %s", t.TypeNo, t.ClassNo,
                                t.Loc.Path, t.Loc.Pos, t.Loc.LineNo,
                                t.Loc.ColNo, t.Data)
}

//----------------------------------------------------------------------------
//                                  New
//----------------------------------------------------------------------------

func NewToken() *Token {
    return &Token{}
}




