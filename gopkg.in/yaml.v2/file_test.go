package yaml_v2

import (
	"github.com/vinzenz/yaml"
	"io/ioutil"
	"testing"
)

func TestFileLoad(t *testing.T) {
	data, err := ioutil.ReadFile("./files/config.yaml")
	if err != nil {
		t.Error(err)
	}
	t.Log(string(data))

	var obj map[string]interface{}
	err = yaml.Unmarshal(data, &obj)
	if err != nil {
		t.Error(err)
	}
	t.Log(obj)
}
