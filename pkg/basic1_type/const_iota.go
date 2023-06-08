package main

const (
	_           = iota
	IPV6_V6ONLY // 1
	SOMAXCONN   // 2
	SO_ERROR    // 3
)
const (
	_ = iota // 0
	Pin1
	Pin2
	Pin3
	_
	Pin5 // 5
)
const (
	a = iota + 1 // 1, iota = 0
	b            // 2, iota = 1
	c            // 3, iota = 2
)
const (
	i = iota << 1 // 0, iota = 0
	j             // 2, iota = 1
	k             // 4, iota = 2
)
