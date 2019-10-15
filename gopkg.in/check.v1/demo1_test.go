package check

import(
	"testing"
	"io"
	. "gopkg.in/check.v1"
)

/*
    func (s *SuiteType) SetUpSuite(c *C) - Run once when the suite starts running.
	func (s *SuiteType) SetUpTest(c *C) - Run before each test or benchmark starts running.
	func (s *SuiteType) TearDownTest(c *C) - Run after each test or benchmark runs.
	func (s *SuiteType) TearDownSuite(c *C) - Run once after all tests or benchmarks have finished running.
*/

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{
	param1 string
}

var _ = Suite(&MySuite{})

func (s *MySuite)SetUpSuite(c *C) {
	s.param1 = "42"
}

func (s *MySuite)TestDemo1(c *C) {
	c.Assert(s.param1, Equals, "42")
    c.Assert(io.ErrClosedPipe, ErrorMatches, "io: .*on closed pipe")
    c.Check(42, Equals, 42)
}

func (s *MySuite)TestDemo2(c *C) {
	c.Assert(s.param1, Equals, 42)

}

