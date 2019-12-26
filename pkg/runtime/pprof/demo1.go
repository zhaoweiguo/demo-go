package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
		doit()
	}
}

func doit() {
	var trips = [][]string{
		[]string{"重庆", "新疆"},
		[]string{"广州", "南京"},
		[]string{"上海", "广州"},
		[]string{"西藏", "重庆"},
		[]string{"北京", "上海"},
		[]string{"南京", "西藏"},
	}
	var result []string
	i := 1
	for _, trip := range trips {
		//log.Println("-", trip)
		if len(result) == 0 {
			result = trip
			continue
		}
		//log.Println(newTrips)
		for _, trip := range trips {
			//log.Println("$", trip)
			if result[0] == trip[1] {
				result = append(trip, result[1:]...)
				//log.Println("=====", result)
				break
			} else if result[len(result)-1] == trip[0] {
				result = append(result, trip[1])
				//log.Println("====>", result)
				break
			} else {
				log.Println("=====")
			}
		}
		i++
	}
}
