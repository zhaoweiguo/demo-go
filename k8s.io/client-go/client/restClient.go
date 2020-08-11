package main

import (
	"context"
	client_go "github.com/zhaoweiguo/demo-go/k8s.io/client-go"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"log"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	restClient := client_go.GetRestClient()

	ctx := context.Background()
	result := &corev1.PodList{}
	err := restClient.Get().Namespace("default").
		Resource("pods").
		VersionedParams(&metav1.ListOptions{Limit: 500}, scheme.ParameterCodec).
		Do(ctx).Into(result)
	if err != nil {
		panic(err)
	}

	for _, d := range result.Items {
		log.Printf("NameSpace:%v \t Name:%v \t Status:%+v\n", d.Namespace, d.Name, d.Status)
	}
}
