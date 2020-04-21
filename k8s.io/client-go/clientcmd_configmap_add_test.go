package client_go

import (
	"io/ioutil"
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestAddConfigMap(t *testing.T) {
	clientSet := getClientset()

	cases := []struct {
		name      string
		namespace string
	}{
		{"volume-name", "pvt"},
	}

	labels := map[string]string{"kt": "kt"}
	for _, c := range cases {
		cli := clientSet.CoreV1().ConfigMaps(c.namespace)
		configMap := getConfigMap(labels, c.name, c.namespace)
		cli.Create(&configMap)
	}
}

func getConfigMap(labels map[string]string, name, namespace string) v1.ConfigMap {
	objectMeta := metav1.ObjectMeta{
		Name:      name,
		Namespace: namespace,
		Labels:    labels,
	}
	publicKey, err := ioutil.ReadFile("./files/public.key")
	privateKey, err := ioutil.ReadFile("./files/private.key")
	if err != nil {
		panic(err.Error())
	}
	data := map[string]string{
		SSHAuthKey:        string(publicKey),
		SSHAuthPrivateKey: string(privateKey),
	}
	cm := v1.ConfigMap{
		ObjectMeta: objectMeta,
		Data:       data,
	}
	return cm
}
