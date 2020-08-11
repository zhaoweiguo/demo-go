package client_go

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func GetClientset() *kubernetes.Clientset {
	return getClientset()
}

func getClientset() *kubernetes.Clientset {
	config := getBuildConfig()
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset
}

func GetRestClient() *rest.RESTClient {
	config := getBuildConfig()
	config.APIPath = "api"
	config.GroupVersion = &corev1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs

	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}
	return restClient

}

func GetDiscoveryClient() *discovery.DiscoveryClient {
	config := getBuildConfig()
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		log.Panic(err)
	}
	return discoveryClient

}

func getBuildConfig() *rest.Config {

	kubeconfig := "/Users/zhaoweiguo/.kube/ali.config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	return config
}
