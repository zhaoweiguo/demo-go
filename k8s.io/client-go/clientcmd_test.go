package client_go

import (
	"fmt"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func TestListAllPods(t *testing.T) {
	clientset := getClientset()
	// 取出所有的pod列表
	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	log.Printf("There are %d pods in the cluster\n", len(pods.Items))
}

func TestListPods(t *testing.T) {
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
		pods, err := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		log.Printf("There are %d pods in ns:%s\n", len(pods.Items), namespace)
	}
}

func TestGetPod(t *testing.T) {
	cases := []struct {
		namespace string
		pod       string
	}{
		{"default", "vision-ai-8578f6f87c-ngn7p"},
		{"default", "not-exist-pod"},
		{"default", "goanalysis-1587254400-vqqbq"},
		{"default", "goanalysis-1586563200-98s4h"},
	}
	clientset := getClientset()
	for _, c := range cases {
		namespace := c.namespace
		pod := c.pod
		_, err := clientset.CoreV1().Pods(namespace).Get(pod, metav1.GetOptions{})

		if errors.IsNotFound(err) {
			fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
			fmt.Printf("Error getting pod %s in namespace %s: %v\n",
				pod, namespace, statusError.ErrStatus.Message)
		} else if err != nil {
			panic(err.Error())
		} else {
			fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
		}
	}

}

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
