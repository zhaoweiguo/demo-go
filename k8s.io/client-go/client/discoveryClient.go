package main

import (
	client_go "github.com/zhaoweiguo/demo-go/k8s.io/client-go"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)
}

func main() {
	discoveryClient := client_go.GetDiscoveryClient()
	groups, apiResourceLists, err := discoveryClient.ServerGroupsAndResources()
	if err != nil {
		log.Panic(err)
	}
	//for _, group := range groups {
	//	log.Println(group)
	//}
	for _, apiResourceList := range apiResourceLists {
		kind := apiResourceList.Kind
		apiVersion := apiResourceList.APIVersion
		apiResources := apiResourceList.APIResources
		groupVersion := apiResourceList.GroupVersion
		log.Println(kind)
		log.Println(apiVersion)
		log.Println(apiResources)
		log.Println(groupVersion)
		log.Println("=====================")
	}

}
