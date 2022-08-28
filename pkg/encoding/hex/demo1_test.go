package hex

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func init()  {
	log.SetFlags(log.Lshortfile)
}

var (
	str = []byte("Hello world!")
	strHex = "48656c6c6f20776f726c6421"
)

func TestEncodedLen(t *testing.T) {
	// hex.EncodedLen: 长度检测
	// hex.EncodedLen(len(srcCode)) 返回值实际是 len(srcCode) 长度的2倍
	lenStr := len(str)
	lenHex := hex.EncodedLen(len(str))
	assert.Equal(t, lenStr, 12)
	assert.Equal(t, lenHex, 24)

	log.Printf("srccode(%T) = %v\n", str, str)
}

func TestDecodedLen(t *testing.T)  {
	// hex.DecodedLen(len(dstEncode)) 返回值实际是len(dstEncode) 长度一半
	lenStrHex := len(strHex)
	lenStrDecode := hex.DecodedLen(lenStrHex)
	assert.Equal(t, lenStrHex, 24)
	assert.Equal(t, lenStrDecode, 12)
}


func TestEncodeAndDecode(t *testing.T) {
	str := []byte("Hello world!")

	lenEncode := hex.EncodedLen(len(str))

	// hex.Encode: 转为对应16进制的字符串
	bytesHex := make([]byte, lenEncode)
	lenEncode2 := hex.Encode(bytesHex, str)
	assert.Equal(t, lenEncode, lenEncode2)
	assert.Equal(t, strHex, []byte(""))

	// hex.Decode: 把16进制的字符串转回原字符串
	lenDecode := hex.DecodedLen(len(bytesHex))
	bytesBack := make([]byte, lenDecode)
	lenDecode2, err := hex.Decode(bytesHex, bytesBack)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, lenDecode, lenDecode2)
	assert.Equal(t, str, bytesBack)

	log.Printf("hex.Decode(%T) = %v\n", bytesBack, bytesBack)
}

func TestEncodeStringAndDecodeString(t *testing.T) {
	// 将 srcCode 编码为字符串
	encodedStr := hex.EncodeToString(str)
	assert.Equal(t, encodedStr, strHex)

	strBack, _ := hex.DecodeString(encodedStr)
	assert.Equal(t, str, strBack)

}




func TestDump(t *testing.T) {
	content := []byte("Go is an open source programming language.")
	// 转储返回一个包含给定数据的十六进制转储的字符串。十六进制转储的格式与hexdump -C命令行上的输出相匹配
	log.Printf("dump(%T) =  \n%v\n", hex.Dump(content), hex.Dump(content))

	log.Println("-------------------------------dumper部分---------------------------------")
	lines := []string{
		"Go is an open source programming language.",
		"\n",
		"We encourage all Go users to subscribe to golang-announce.",
	}
	//Dumper 返回一个 WriteCloser，它将所有写入数据的十六进制转储写入 w。转储的格式与hexdump -C命令行上的输出相匹配
	stdoutDumper := hex.Dumper(os.Stdout)
	defer stdoutDumper.Close()

	for _, line := range lines {
		stdoutDumper.Write([]byte(line))
	}
}


