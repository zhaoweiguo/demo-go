package main

import (
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	f, err := os.Create(".cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
}

func cpuProfile() {
	f, err := os.OpenFile("cpu.prof", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	log.Println("CPU Profile started")
	//pprof.StartCPUProfile(f)
	//defer pprof.StopCPUProfile()
	pprof.WriteHeapProfile(f)

	time.Sleep(10 * time.Second)
	log.Println("CPU Profile stopped")
}

func doSomeThingOne(times int) {
	for i := 0; i < times; i++ {
		for j := 0; j < times; j++ {
			log.Print(i, j)
		}
		log.Println(i)
	}
}

func HoareSort(list []int, low int, high int) {
	if low >= high {
		return
	}

	//[[----递归模板区

	pivot := list[low]
	right := high
	left := low
	for {
		for list[right] >= pivot && right > low {
			right--
		}
		for list[left] <= pivot && left < high {
			left++
		}

		if left < right {
			list[left], list[right] = list[right], list[left]
		} else {
			break
		}
	}
	list[low], list[right] = list[right], list[low]
	//--------------]]

	HoareSort(list, low, right)

	HoareSort(list, right+1, high)
}
