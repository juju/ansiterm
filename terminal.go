// Copyright 2016 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package ansiterm

import (
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
)

// colorEnabledWriter returns a writer that can handle the ansi color codes
// and true if the writer passed in is a terminal capable of color. If the
// TERM environment variable is set to "dumb", the terminal is not considered
// color capable.
func colorEnabledWriter(w io.Writer) (io.Writer, bool) {
	f, ok := w.(*os.File)
	if !ok {
		return w, false
	}
	// NO_COLOR is a relatively new standard for preventing color enabled
	// writers rather than using the TERM env.
	// See: https://no-color.org/
	if truthy(os.Getenv("NO_COLOR")) {
		return w, false
	}
	// Check the TERM environment variable specifically
	// to check for "dumb" terminals.
	if os.Getenv("TERM") == "dumb" {
		return w, false
	}
	if !isatty.IsTerminal(f.Fd()) {
		return w, false
	}
	return colorable.NewColorable(f), true
}

func truthy(value string) bool {
	if v := strings.ToLower(value); v == "y" || v == "true" {
		return true
	}
	if num, _ := strconv.Atoi(value); num > 0 {
		return true
	}
	return false
}
