// Copyright 2016 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package ansiterm

import (
	"io"

	"github.com/juju/ansiterm/tabwriter"
)

// NewTabWriter returns a writer that is able to set colors and styels.
// The ansi escape codes are stripped for width calculations.
func NewTabWriter(output io.Writer, minwidth, tabwidth, padding int, padchar byte, flags uint) *Writer {
	writer, colorCapable := colorEnabledWriter(output)
	tab := tabwriter.NewWriter(writer, minwidth, tabwidth, padding, padchar, flags)

	return &Writer{
		Writer:  tab,
		noColor: !colorCapable,
	}
}
