package int

import (
	"github.com/bmizerany/assert"
	"math"
	"testing"
)

// 向上取整 math.Ceil()
// 向下取整 math.Floor()
func TestRound(t *testing.T) {
	assert.Equal(t, 10/3, 3)
	assert.Equal(t, int(math.Ceil(float64(10)/3)), 4)
	assert.Equal(t, int(math.Floor(float64(10)/3)), 3)
}
