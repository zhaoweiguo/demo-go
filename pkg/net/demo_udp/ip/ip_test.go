package ip

import (
	"log"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIp(t *testing.T) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	assert.Nil(t, err)
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	log.Println(localAddr.String())
}
