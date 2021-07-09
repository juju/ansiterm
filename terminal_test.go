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

	os.Setenv("NO_COLOR", "true")

	writer, ok := colorEnabledWriter(file)
	c.Assert(ok, gc.Equals, false)
	c.Assert(writer, gc.Equals, file)
}

func (s *colorWriterSuite) TestTruthy(c *gc.C) {
	tests := []struct {
		Value    string
		Expected bool
	}{{
		Value:    "y",
		Expected: true,
	}, {
		Value:    "Y",
		Expected: true,
	}, {
		Value:    "TRUE",
		Expected: true,
	}, {
		Value:    "true",
		Expected: true,
	}, {
		Value:    "1",
		Expected: true,
	}, {
		Value:    "2",
		Expected: true,
	}, {
		Value:    "10000",
		Expected: true,
	}, {
		Value:    "f",
		Expected: false,
	}, {
		Value:    "",
		Expected: false,
	}, {
		Value:    "FALSE",
		Expected: false,
	}, {
		Value:    "0",
		Expected: false,
	}, {
		Value:    "-1",
		Expected: false,
	}}
	for k, test := range tests {
		c.Logf("test: %d, value: %v", k, test.Value)

		t := truthy(test.Value)
		c.Assert(t, gc.Equals, test.Expected)
	}
}
