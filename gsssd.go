package main

import (
	"log"
	"os"

	"github.com/ossareh/gsssd/cpu"
	"github.com/ossareh/libgosysstat/core"
	cpuProcessor "github.com/ossareh/libgosysstat/processor/cpu"
)

const TICK_INTERVAL = 1

type StatSender interface {
	Send([]core.Stat)
	Close()
}

func main() {
	cpuClient, err := cpu.New("123.123.123.123", "prefix")
	if err != nil {
		log.Fatal(err)
	}
	defer cpuClient.Close()

	procStat, err := os.Open("/proc/stat")
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer procStat.Close()
	cpuStatProcessor := cpuProcessor.NewProcessor(procStat)
	cpuStatResults := make(chan []core.Stat)

	go core.StatProcessor(cpuStatProcessor, TICK_INTERVAL, cpuStatResults)
	for {
		select {
		case stats := <-cpuStatResults:
			cpuClient.Send(stats)
		}
	}
}
