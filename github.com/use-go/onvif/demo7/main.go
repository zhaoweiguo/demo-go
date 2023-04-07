package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/use-go/onvif"
	"github.com/use-go/onvif/gosoap"
	"github.com/use-go/onvif/networking"
)

func main() {
	username := "admin"
	password := "admin123"
	acceptedData := `
<GetStreamUri xmlns="http://www.onvif.org/ver10/media/wsdl">   
    <StreamSetup>
        <Stream xmlns="http://www.onvif.org/ver10/schema">RTP-Unicast</Stream>
        <Transport xmlns="http://www.onvif.org/ver10/schema">
          <Protocol>UDP</Protocol>
        </Transport>
      </StreamSetup>   
      <ProfileToken>000</ProfileToken>
    </GetStreamUri>
`
	do(acceptedData, username, password)
}

func do(acceptedData, username, password string) {

	endpoint := "http://192.168.124.58:8899/onvif/media_service"
	soap := gosoap.NewEmptySOAP()
	soap.AddStringBodyContent(acceptedData)
	soap.AddRootNamespaces(onvif.Xlmns)
	soap.AddWSSecurity(username, password)

	servResp, err := networking.SendSoap(new(http.Client), endpoint, soap.String())
	if err != nil {
		panic(err)
	}
	rsp, err := ioutil.ReadAll(servResp.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(rsp))

}
