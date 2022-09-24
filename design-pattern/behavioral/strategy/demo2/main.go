/*
这部分代码 Visitor 模式实现方案: ../../visitor/demo1
*/
package main

import (
	"encoding/json"
	"encoding/xml"
	"log"
)

type IStrategy interface {
	do(content interface{}) string
}
type XmlVisitor struct{}

func (xv XmlVisitor) do(content interface{}) string {
	bytes, err := xml.Marshal(content)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

type JsonVisitor struct{}

func (jv JsonVisitor) do(content interface{}) string {
	bytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

type Circle struct {
	Radius int
}

type Rectangle struct {
	Width, Height int
}

type Operator struct {
	strategy IStrategy
}

// 设置策略
func (operator *Operator) setStrategy(strategy IStrategy) {
	operator.strategy = strategy
}

// 调用策略中的方法
func (operator *Operator) marshal(content interface{}) string {
	return operator.strategy.do(content)
}

func init() {
	log.SetFlags(log.Llongfile)
}

func main() {
	operator := Operator{}
	rectangle := Rectangle{10, 20}
	circle := Circle{15}
	// 1. json
	operator.setStrategy(JsonVisitor{})
	log.Println(operator.marshal(rectangle))
	log.Println(operator.marshal(circle))

	// 2. xml
	operator.setStrategy(XmlVisitor{})
	log.Println(operator.marshal(rectangle))
	log.Println(operator.marshal(circle))

}
