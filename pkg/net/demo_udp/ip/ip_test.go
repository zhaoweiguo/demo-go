package ip

import (
	"github.com/stretchr/testify/assert"
	"log"
	"net"
	"testing"
)

func TestIp(t *testing.T) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	assert.Nil(t, err)
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	log.Println(localAddr.String())
}
