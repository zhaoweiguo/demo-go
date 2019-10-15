package check

import(
	. "gopkg.in/check.v1"
)

type MyBenchSuite1 struct{
}

var _ = Suite(&MyBenchSuite1{})

func (s *MyBenchSuite1)TestDemo1(c *C) {
	for i:=0; i<c.N; i++{
		
	}
}
