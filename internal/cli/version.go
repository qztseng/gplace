package cli

import (
	"fmt"

	"github.com/alecthomas/kong"
)

const Version = "0.1.0"

type VersionFlag string

func (v VersionFlag) Decode(_ *kong.DecodeContext) error { return nil }
func (v VersionFlag) IsBool() bool                       { return true }

func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Fprintln(app.Stdout, vars["version"])
	app.Exit(0)
	return nil
}
