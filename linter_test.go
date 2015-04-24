package main

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

type TestSuite struct{}

var _ = Suite(&TestSuite{})

func (s *TestSuite) TestSuccessfulFetch(c *C) {
	res := fetch("http://www.google.com")
	c.Assert(res.StatusCode, Equals, 200)
}
