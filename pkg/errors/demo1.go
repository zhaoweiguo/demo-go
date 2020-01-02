package main

import (
	"errors"
	"fmt"
)

func main() {
	var ErrTimeOut = errors.New("执行者执行超时")
	var ErrInterrupt = errors.New("执行者被中断")
	fmt.Println(ErrInterrupt, ErrTimeOut)
}
