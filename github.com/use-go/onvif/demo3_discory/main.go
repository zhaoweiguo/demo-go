package main

import (
	"encoding/json"
	"log"
	"path"
	"regexp"
	"strings"

	"github.com/beevik/etree"
	discover "github.com/use-go/onvif/ws-discovery"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}

func main() {
	runDiscovery("en0")
}

// Host host
type Host struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

func runDiscovery(interfaceName string) {
	var hosts []*Host
	devices, err := discover.SendProbe(interfaceName, nil, []string{"dn:NetworkVideoTransmitter"}, map[string]string{"dn": "http://www.onvif.org/ver10/network/wsdl"})
	if err != nil {
		log.Printf("error %s", err)
		return
	}
	for _, j := range devices {
		log.Println(j)
		doc := etree.NewDocument()
		if err := doc.ReadFromString(j); err != nil {
			log.Printf("error %s", err)
		} else {

			doc.Root().Text()
			endpoints := doc.Root().FindElements("./Body/ProbeMatches/ProbeMatch/XAddrs")
			scopes := doc.Root().FindElements("./Body/ProbeMatches/ProbeMatch/Scopes")

			flag := false

			host := &Host{}

			for _, xaddr := range endpoints {
				xaddr := strings.Split(strings.Split(xaddr.Text(), " ")[0], "/")[2]
				host.URL = xaddr
			}
			if flag {
				break
			}
			for _, scope := range scopes {
				re := regexp.MustCompile(`onvif:\/\/www\.onvif\.org\/name\/[A-Za-z0-9-]+`)
				match := re.FindStringSubmatch(scope.Text())
				host.Name = path.Base(match[0])
			}

			hosts = append(hosts, host)

		}

	}

	bys, _ := json.Marshal(hosts)
	log.Printf("done %s", bys)
}
