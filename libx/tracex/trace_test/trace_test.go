package trace_test

import (
	trace "github.com/zhaoweiguo/demo-go/libx/tracex"
)

func a() {
	defer trace.Trace()()
	b()
}
func b() {
	defer trace.Trace()()
	c()
}
func c() {
	defer trace.Trace()()
	d()
}
func d() {
	defer trace.Trace()()
}
func ExampleTrace() {
	a()
	// Output:
	// g[00001]:    ->github.com/zhaoweiguo/demo-go/libx/tracex/trace_test_test.a
	// g[00001]:        ->github.com/zhaoweiguo/demo-go/libx/tracex/trace_test_test.b
	// g[00001]:            ->github.com/zhaoweiguo/demo-go/libx/tracex/trace_test_test.c
	// g[00001]:                ->github.com/zhaoweiguo/demo-go/libx/tracex/trace_test_test.d
	// g[00001]:                <-github.com/zhaoweiguo/demo-go/libx/tracex/trace_test_test.d
	// g[00001]:            <-github.com/zhaoweiguo/demo-go/libx/tracex/trace_test_test.c
	// g[00001]:        <-github.com/zhaoweiguo/demo-go/libx/tracex/trace_test_test.b
	// g[00001]:    <-github.com/zhaoweiguo/demo-go/libx/tracex/trace_test_test.a
}
