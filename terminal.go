// Copyright 2016 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package ansiterm

import (
	"io"
	"os"

	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
)

// colorEnabledWriter returns a writer that can handle
// the ansi color codes and true if the writer passed in
// is a terminal capable of color.
func colorEnabledWriter(w io.Writer) (io.Writer, bool) {
	f, ok := w.(*os.File)
	if !ok {
		return w, false
	}
	if !isatty.IsTerminal(f.Fd()) {
		return w, false
	}
	return colorable.NewColorable(f), true
}
