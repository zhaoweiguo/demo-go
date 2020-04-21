package client_go

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"testing"
)

func TestGetDeployment(t *testing.T) {
	clientset := getClientset()

	cases := []struct {
		name      string
		namespace string
	}{
		{"addms", "default"},
		{"svc-notice", "pvt"},
	}

	for _, c := range cases {
		ns := c.namespace
		name := c.name
		app, err := clientset.AppsV1().Deployments(ns).Get(name, metav1.GetOptions{})
		if err != nil {
			panic(err.Error())
		}
		log.Println(app.GetName(), *app.Spec.Replicas)
		log.Println(app.Status, app.Name, app.ClusterName)
		log.Println(app)
	}
}
