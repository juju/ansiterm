package tabwriter

import (
	"bytes"
	"strings"
	"testing"

	gc "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	gc.TestingT(t)
}

type tabwriterSuite struct{}

var _ = gc.Suite(&tabwriterSuite{})

func (s *tabwriterSuite) TestRightAlignOverflow(c *gc.C) {
	var buf bytes.Buffer
	tw := NewWriter(&buf, 0, 1, 2, ' ', 0)
	tw.SetColumnAlignRight(2)
	tw.Write([]byte("not\tenough\ttabs"))
	tw.Flush()
	c.Assert(buf.String(), gc.Equals, "not  enough  tabs")
}

func (s *tabwriterSuite) TestRuneWidth(c *gc.C) {
	var text strings.Builder
	tw := NewWriter(&text, 0, 1, 2, ' ', 0)
	tw.Write([]byte("界\tab\n"))
	tw.Write([]byte("cd\tef\n"))
	tw.Flush()
	c.Assert(text.String(), gc.Equals, "界  ab\ncd  ef\n")

	text.Reset()
	tw.Write([]byte("界\ta\n"))
	tw.Write([]byte("b\tc\n"))
	tw.Flush()
	c.Assert(text.String(), gc.Equals, "界  a\nb   c\n")
}
