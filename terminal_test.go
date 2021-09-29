// Copyright 2021 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package ansiterm

import (
	"io/ioutil"
	"os"

	gc "gopkg.in/check.v1"
)

type colorWriterSuite struct{}

var _ = gc.Suite(&colorWriterSuite{})

func (s *colorWriterSuite) TestNoColor(c *gc.C) {
	file, err := ioutil.TempFile("", "")
	c.Assert(err, gc.IsNil)

	os.Setenv("NO_COLOR", "")
	defer os.Unsetenv("NO_COLOR")

	writer, ok := colorEnabledWriter(file)
	c.Assert(ok, gc.Equals, false)
	c.Assert(writer, gc.Equals, file)

}

func (s *colorWriterSuite) TestNoColorEvenIfFalse(c *gc.C) {
	file, err := ioutil.TempFile("", "")
	c.Assert(err, gc.IsNil)

	os.Setenv("NO_COLOR", "false")
	defer os.Unsetenv("NO_COLOR")

	writer, ok := colorEnabledWriter(file)
	c.Assert(ok, gc.Equals, false)
	c.Assert(writer, gc.Equals, file)
}
