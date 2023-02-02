package filepath

import (
	"github.com/bmizerany/assert"
	"path/filepath"
	"testing"
)

func TestExt(t *testing.T) {
	cases := []struct {
		filename string
		ext      string
	}{
		{"NO1.txt", ".txt"},
		{"NO2.yaml", ".yaml"},
		{"NO3.json", ".json"},
		{"NO4.exe", ".exe"},
	}
	for _, c := range cases {
		ext := filepath.Ext(c.filename)
		assert.Equal(t, ext, c.ext)
	}

}
