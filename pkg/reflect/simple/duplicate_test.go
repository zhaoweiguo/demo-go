package simple

import (
	"github.com/bmizerany/assert"
	"sort"
	"testing"
)

func TestDuplicate(t *testing.T) {

	b := []string{"a", "b", "c", "c", "e", "f", "a", "g", "b", "b", "c"}
	sort.Strings(b)
	newb := Duplicate(b)
	assert.Equal(t, newb, []interface{}{"a", "b", "c", "e", "f", "g"})

	c := []int{1, 1, 2, 4, 6, 7, 8, 4, 3, 2, 5, 6, 6, 8}
	sort.Ints(c)
	assert.Equal(t, Duplicate(c), []interface{}{1, 2, 3, 4, 5, 6, 7, 8})
}
