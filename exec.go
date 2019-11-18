// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// Miscellaneous utility functions

// Note that these functions are made unique by attaching them to a structure
// which must be instanced before they can be used.

// Some functions were taken from https://blog.kowalczyk.info/book/go-cookbook.html
// which was declared public domain at the time that the functions were taken.

package util

import (
	"os/exec"
	"strings"
	"unicode"
)

//============================================================================
//                             Command Execution
//============================================================================

// os.Exec contains further details
type ExecCmd struct {
	cmd       	*exec.Cmd
}

func (c *ExecCmd) Cmd( ) *exec.Cmd {
	return c.cmd
}

func (c *ExecCmd) CommandString( ) string {
	n := len(c.cmd.Args)
	a := make([]string, n, n)
	for i := 0; i < n; i++ {
		a[i] = c.QuoteArgIfNeeded(i)
	}
	return strings.Join(a, " ")
}

func (c *ExecCmd) SetCommandString(cmd string) {
	args, err := ParseCommandLine(cmd)
	if err == nil {
		c.cmd.Args = args
	}
}

func (c *ExecCmd) QuoteArgIfNeeded(n int) string {
	var s		string

	s = c.cmd.Args[n]
	if strings.Contains(s, " ") || strings.Contains(s, "\"") {
		s = strings.Replace(s, `"`, `\"`, -1)
		return `"` + s + `"`
	}
	return s
}

// Run runs the previously set up command.
func (c *ExecCmd) Run( ) error {
	var err		error

	err = c.cmd.Run()

	return err
}

// RunWithOutput runs the previously set up command, gets the combined output
// of sysout and syserr, trims whitespace from it and returns if error free.
// If any error occurs, it is simply returned.
func (c *ExecCmd) RunWithOutput( ) (string, error) {
	var err		error

	outBytes, err := c.cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	s := string(outBytes)
	s = strings.TrimSpace(s)

	return s, nil
}

//----------------------------------------------------------------------------
//							Class Functions
//----------------------------------------------------------------------------

func NewExec() *ExecCmd {
	ce := ExecCmd{}
	return &ce
}

func NewExecArgs(name string, args... string) *ExecCmd {
	ce := ExecCmd{}
	if len(name) > 0 {
		ce.cmd = exec.Command(name, args...)
	}
	return &ce
}

/*   ParseCommand parses the program command line breaking it
     up into substrings forming an argv/argc structure. Arguments are
     separated by whitespace. An AStrC Array is used for argv/argc.
     We don't have to completely parse the command line, just break
     it up being coznizant of the quoted strings which assumes that
     they should either be by themselves or at the end of an argument.
     Note: we currently do not handle embedded quotes.
 */
func ParseCommandLine(cmd string) ([]string, error) {
	var args	[]string
	var i		int

	cmdRunes := []rune(cmd)
	iMax := len(cmdRunes)

	for i=0; i<iMax; {
		// Skip leading spaces.
		for (i < iMax) && unicode.IsSpace(cmdRunes[i]) {
			i++
		}
		// Scan off an argument.
		if (i < iMax) {
			start := i
			for (i < iMax) && !unicode.IsSpace(cmdRunes[i]) {
				if (cmdRunes[i] == '"') || (cmdRunes[i] == '\'') {
					quote := cmdRunes[i]; i++
					for (i < iMax) && !(cmdRunes[i] == quote) {
						i++
					}
				}
				if !(i < iMax) || unicode.IsSpace(cmdRunes[i]) {
					break
				}
				i++
			}
			if (start < iMax) && (start <= iMax) {
				args= append(args, string(cmdRunes[start:i]))
			}
			i++
		}
	}

	return args, nil
}



