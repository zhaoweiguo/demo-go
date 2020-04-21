package client_go

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"testing"
)

func TestListService(t *testing.T) {
	clientset := getClientset()
	cases := []struct {
		namespace string
	}{
		{"default"},
		{"pvt"},
		{"api"},
	}

	for _, c := range cases {
		namespace := c.namespace
		// 取出所有的pod列表
		services, err := clientset.CoreV1().Services(namespace).List(metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("len(services): %d\n", len(services.Items))
		hosts := map[string]string{}
		for _, service := range services.Items {
			hosts[service.ObjectMeta.Name] = service.Spec.ClusterIP
		}
		log.Println(hosts)
	}
}
