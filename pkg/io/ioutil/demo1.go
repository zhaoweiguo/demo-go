package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

func main(){

	data, err := ioutil.ReadFile("./city.txt")
	if err != nil {
		panic(err.Error())
	}
	
	citys := strings.Split(string(data), "--")

	for _, city := range citys {
		if city == "" {
			continue
		}
		fmt.Println(city)
	}

	fmt.Println("=======================")
	dirs, err := ioutil.ReadDir(".")	// 获得指定目录下的文件列表
	if err != nil {
		panic(err.Error())
	}
	for _, dir := range dirs {
		Name := dir.Name()
		Size := dir.Size()
		Mode := dir.Mode()
		ModeTime := dir.ModTime()
		IsDir := dir.IsDir()
		//Sys := dir.Sys()
		fmt.Printf("%s\t%d\t%d\t%s\t\t%t\n",
			Name, Size, Mode, ModeTime.String(),IsDir)
	}


//	fmt.Println(len(citys))


	



}	
/*	


	// step1: add host
	hostname := "*.corge.at"
	host1 := &Host{
		Name:      hostname,
		Locations: []*Location{},
	}
	backend.AddHost(host1)

	// step2: add upstream
	upstreamId := "up1"
	upstream1 := &Upstream{
		Id:        upstreamId,
		Endpoints: []*Endpoint{},
	}
	upstream1, err = backend.AddUpstream(upstream1)
	if err != nil { // 错误类型处理@todo
		log.Errorf("init upstream error:%v", err)
		// @todo send signal 
		return err
	}
	// step3: add endpoint
	endpointId := "e1"
	endpoint1 := &Endpoint{
		Url:        "http://192.168.35.141:5101",
		Id:         endpointId,
		UpstreamId: upstreamId,
	}
	endpoint1, err = backend.AddEndpoint(endpoint1)
	if err != nil {
		log.Errorf("init endpoint error:%v", endpoint1)
		//@todo
		return err
	}
	//  add other endpoint
	endpointId2 := "e2"
	endpoint2 := &Endpoint{
		Url:        "http://192.168.35.141:5102",
		Id:         endpointId2,
		UpstreamId: upstreamId,
	}
	endpoint1, err = backend.AddEndpoint(endpoint2)
	if err != nil {
		log.Errorf("init endpoint error:%v", endpoint2)
		//@todo
		return err
	}

	//  step4: add location
	locationId := "root"
	path := "/"
	location := &Location{
		Hostname:    hostname,
		Path:        path,
		Id:          locationId,
		Upstream:    upstream1,
		Middlewares: []*MiddlewareInstance{},
	}
	location, err = backend.AddLocation(location)
	if err != nil {
		log.Errorf("init location error:%v", location)
	}
	//  add another location
	locationId2 := "home"
	path2 := "/home"
	location2 := &Location{
		Hostname:    hostname,
		Path:        path2,
		Id:          locationId2,
		Upstream:    upstream1,
		Middlewares: []*MiddlewareInstance{},
	}
	location2, err = backend.AddLocation(location2)
	if err != nil {
		log.Errorf("init location error:%v", location2)
	}

	// step5: add middleware
	bytes := `{"Id": "connlimit", "Priority": 0, "Type": "connlimit", "Middleware":{"variable": "client.ip", "connections": 1}}`
	middlewareId1 := "middleware1"
	middleware1, err := MiddlewareFromJson([]byte(bytes), registry.GetSpec)
	if err != nil {
		log.Errorf("[error] in MiddlewareFromJson:%v\n", err)
		return err
	}

	//log.Infof("2222222midleware:%v", middleware1)

	middleware1.Id = middlewareId1
	middleware1, err = backend.AddLocationMiddleware(hostname, locationId, middleware1)


	//checkData(backend)
	return nil
}

func checkData(backend Backend) {
	hosts, _ := backend.GetHosts()
	
	for _, h := range hosts {
		log.Infof("%v", h)
		fmt.Printf("%v\n", h)
	}
}


*/
