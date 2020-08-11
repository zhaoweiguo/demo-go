package normal

import "log"

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)
}

// Bird is a nice type.
type Bird struct {
	Name           string
	LifeExpectance int
}

// Fly is fly.
func (b Bird) Fly() string {
	log.Println("I am flying...")
	return "returnFly"
}
func (b Bird) FlyA(a string) {
	log.Println("I am flying...", a)
}

// Fly is fly.
func (b *Bird) Fly2() {
	log.Println("I am flying2...")
}

type User struct {
	Name string `json:"name" bson:"b_name"`
	Age  int    `json:"age" bson:"b_age"`
}
