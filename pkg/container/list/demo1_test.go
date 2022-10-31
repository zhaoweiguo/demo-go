package main

import (
	"container/list"
	"fmt"
	"log"
	"testing"
)

var (
	l *list.List
)

func init()  {
	log.SetFlags(log.Lshortfile)
}
func newList()  {
	l = list.New() //创建一个新的list
	for i := 0; i < 5; i++ {
		l.PushBack(i)
	}
	myprint(l)
}

func myprint(l *list.List)  {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) //输出l2的值,01234
	}
	fmt.Println()
}

func TestNormal(t *testing.T) {

	newList()
	log.Println()
	myprint(l)
	fmt.Println()

	log.Println(l.Front().Value) //输出首部元素的值,0
	log.Println(l.Back().Value)  //输出尾部元素的值,4
	l.InsertAfter(6, l.Front())  //首部元素之后插入一个值为6的元素
	log.Println()
	myprint(l)
	fmt.Println()

	l.MoveBefore(l.Front().Next(), l.Front()) //首部两个元素位置互换
	log.Println()
	myprint(l)
	fmt.Println()

	l.MoveToFront(l.Back()) //将尾部元素移动到首部
	log.Println()
	myprint(l)
	fmt.Println()
}

func TestCombine(t *testing.T) {

	newList()
	l2 := list.New()
	l2.PushBackList(l) //将l中元素放在l2的末尾
	log.Println()
	fmt.Println()

	l.Init()           //清空l
	log.Print(l.Len()) //0
	myprint(l)
}

func TestRemove(t *testing.T) {
	newList()

	log.Println()
	front := l.Remove(l.Front())
	myprint(l)
	log.Println(front)

	log.Println()
	back := l.Remove(l.Back())
	myprint(l)
	log.Println(back)
}





