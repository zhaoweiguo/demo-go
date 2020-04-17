package main

import (
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func main() {
	data := []byte("")
	var d bson.D
	err := bson.Unmarshal(data, &d)
	if err != nil {
		log.Println(err)
	}

}

/*
M: Map
M is an unordered representation of a BSON document.
This type should be used when the order of the elements does not matter.
This type is handled as a regular map[string]interface{} when encoding and decoding.
Elements will be serialized in an undefined, random order.
If the order of the elements matters, a D should be used instead.
*/
func bsonM() {
	m := bson.M{"foo": "bar", "hello": "world", "pi": 3.14159}
}

/*
A: A is an ordered representation of a BSON array.
*/
func bsonA() {
	a := bson.A{"bar", "world", 3.14159, bson.D{{"qux", 12345}}}

}

/*
D: D is an ordered representation of a BSON document.
*/
func bsonD() {
	d := bson.D{{"foo", "bar"}, {"hello", "world"}, {"pi", 3.14159}}
	log.Println(d)
}

/*
E: E represents a BSON element for a D.
It is usually used inside a D.
*/
func bsonE() {

}
