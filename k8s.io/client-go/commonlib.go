package client_go

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func GetClientset() *kubernetes.Clientset {
	return getClientset()
}

func getClientset() *kubernetes.Clientset {
	kubeconfig := "/Users/zhaoweiguo/.kube/ali.config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset
}
