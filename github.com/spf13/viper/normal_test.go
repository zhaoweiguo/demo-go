package viper

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestKubeWatch(t *testing.T) {
	viper.SetConfigName(".kubewatch.yaml")
	viper.AddConfigPath("/Users/zhaoweiguo/9tool/go/src/github.com/zhaoweiguo/demo-go/github.com/spf13/viper/")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		log.Println(viper.ConfigFileUsed())
	}
}

func TestNormal(t *testing.T) {
	viper.SetConfigName("config")     // name of config file (without extension)
	viper.SetConfigType("yaml")       // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME")      // path to look for the config file in
	viper.AddConfigPath("$HOME/test") // call multiple times to add many search paths
	viper.AddConfigPath(".")          // optionally look for config in the working directory
	err := viper.ReadInConfig()       // Find and read the config file
	if err != nil {                   // Handle errors reading the config file
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

func TestSub(t *testing.T) {
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
