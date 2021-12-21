package Stack

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

// 获取 goroutine id
func TestGoID(t *testing.T) {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	// 得到id字符串
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	fmt.Println(id)
}
