package main

import (
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
)

/*
在sync.WaitGroup功能的基础上，增加了错误传递，以及在发生不可恢复的错误时取消整个goroutine集合，或者等待超时
其中包含的函数如下：

func WithContext(ctx context.Context) (*Group, context.Context)
func (g *Group) Go(f func() error)
func (g *Group) Wait() error
*/

func output(num int) (int, error) {
	if num < 0 {
		return 0, errors.New("math: square root error!")
	}
	return num, nil
}

func main() {
	group := new(errgroup.Group)
	nums := []int{
		-1,
		0,
		1,
	}

	for _, num := range nums {
		num := num
		group.Go(func() error {
			res, err := output(num)
			fmt.Println(res)
			return err
		})
	}

	if err := group.Wait(); err != nil {
		fmt.Println("Get errors: ", err)
	} else {
		fmt.Println("Get all num successfully!")
	}

}
