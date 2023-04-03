package LittleEndian

import (
	"encoding/binary"
	"encoding/hex"
	"github.com/bmizerany/assert"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func TestLittleEndian(t *testing.T) {

	tests := []struct {
		name string
		args int
		want []byte
	}{
		{"demo1", 1, []byte{1, 0}},
		{"demo2_user", 16, []byte{16, 0}},
		{"demo2_user", 128, []byte{128, 0}},
		{"demo2_user", 255, []byte{255, 0}},
		{"demo2_user", 256, []byte{0, 1}},
		{"demo2_user", 65535, []byte{255, 255}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var b = make([]byte, 2)
			binary.LittleEndian.PutUint16(b, uint16(tt.args))
			log.Println(b)
			log.Println(hex.EncodeToString(b))
			assert.Equal(t, b, tt.want)
		})
	}
}
