package main

type Bird struct {
	Name string
}

func (b *Bird) Fly() {
}

type IFly interface {
	Fly()
}

func main() {
	var fly IFly = new(Bird) // 在这儿定义Bird类具有IFly接口
	fly.Fly()
}
