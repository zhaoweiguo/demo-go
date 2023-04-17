package viper

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func TestKubeWatch(t *testing.T) {
	viper.SetConfigName(".kubewatch.yaml")
	viper.AddConfigPath("./github.com/zhaoweiguo/demo-go/github.com/spf13/viper/")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		log.Println(viper.ConfigFileUsed())
	}
}

// 设定配置文件名、后缀；设定查询目录并载入
func TestNormal(t *testing.T) {
	viper.SetConfigName("config")     // name of config file (without extension)
	viper.SetConfigType("yaml")       // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME")      // path to look for the config file in
	viper.AddConfigPath("$HOME/test") // call multiple times to add many search paths
	viper.AddConfigPath(".")          // optionally look for config in the working directory
	err := viper.ReadInConfig()       // Find and read the config file
	assert.NoError(t, err)
}

// 载入并读取相关数据
func TestRead(t *testing.T) {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	log.Println(viper.ConfigFileUsed())
	log.Println(viper.GetString("key"))
	log.Println(viper.GetStringMapString("keymapstring"))
	log.Println(viper.GetStringMap("keymap"))
	log.Println(viper.GetBool("bool"))
	log.Println(viper.GetInt("int"))
	log.Println(viper.GetFloat64("float64"))
	log.Println(viper.GetTime("time"))
	log.Println(viper.GetDuration("duration"))
}

// 更新配置文件中的指令值
func TestWrite(t *testing.T) {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	log.Println(viper.GetString("key"))
	viper.Set("key", "newvalue"+time.Now().String())
	log.Println(viper.GetString("key"))
	viper.WriteConfig()
}

func TestSubRead(t *testing.T) {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	sub := viper.Sub("keymapstring")

	log.Println(sub.GetString("keystr1"))
	log.Println(viper.GetString("keymapstring.keystr1"))

	assert.Equal(t, viper.Get("keymapstring.keystr1"), sub.Get("keystr1"))

}

func TestSubWrite(t *testing.T) {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	sub := viper.Sub("keymapstring")

	log.Println(sub.GetString("keystr2"))
	// 注意：sub的修改不会影响主的内容
	sub.Set("keystr2", "newvalue"+time.Now().String())
	log.Println(sub.GetString("keystr2"))
	log.Println(viper.GetString("keymapstring.keystr2"))
	assert.NotEqual(t, sub.GetString("keystr2"), viper.GetString("keymapstring.keystr2"))

	// 不生效，无作用，但也不报错
	err = viper.WriteConfig()
	assert.NoError(t, err)

	// 要想修改，要使用如下方法
	viper.Set("keymapstring.keystr3", "keystr3"+time.Now().String())
	err = viper.WriteConfig()
	assert.NoError(t, err)

	// 生成新的文件
	err = sub.WriteConfigAs("a.yaml")
	assert.NoError(t, err)

}
