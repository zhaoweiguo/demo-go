package xml

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"testing"
)

type O struct {
	A string
	B string
}

func TestDemo(t *testing.T) {
	tests := []struct {
		name string
		obj  interface{}
		want string
	}{
		{"demo1", O{A: "a", B: "b"}, "<O><A>a</A><B>b</B></O>"},
		{"demo2_user", map[string]string{"a": "1", "b": "2"}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o1, err := xml.Marshal(tt.obj)
			assert.Nil(t, err)
			assert.Equal(t, string(o1), tt.want)
		})

	}

}
