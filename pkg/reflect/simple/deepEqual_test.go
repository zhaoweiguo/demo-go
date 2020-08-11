package simple

import (
	"github.com/bmizerany/assert"
	"log"
	"reflect"
	"testing"
)

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile | log.Ltime)
}

func TestDeepEqual(t *testing.T) {

	a1 := A{s: "abc"}
	a2 := A{s: "abc"}
	assert.Equal(t, reflect.DeepEqual(a1, a2), true)

	b1 := []int{1, 2}
	b2 := []int{1, 2}
	assert.Equal(t, reflect.DeepEqual(b1, b2), true)

	c1 := map[string]int{"a": 1, "b": 2}
	c2 := map[string]int{"a": 1, "b": 2}
	assert.Equal(t, reflect.DeepEqual(c1, c2), true)

}
