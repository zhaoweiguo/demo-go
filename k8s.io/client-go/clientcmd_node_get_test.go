package client_go

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"testing"
)

func TestGetNode(t *testing.T) {
	clientSet := getClientset()

	cases := []struct {
		name string
	}{
		{"test1"},
	}

	for _, c := range cases {
		log.Println(c)
		nodeList, err := clientSet.CoreV1().Nodes().List(metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}
		for _, node := range nodeList.Items {
			log.Println(node.Name, node.GetName(), node.GetNamespace(), node.Size(), node.GetClusterName(), node.GetGenerateName())
			log.Println(node.Status.Addresses, node.Status.Size())
			log.Println(len(node.Status.Images), node.Status.Images[0])
			log.Println(node.Status.Size())
		}

	}
}
