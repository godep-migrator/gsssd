package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/cactus/go-statsd-client/statsd"
	"github.com/ossareh/gsssd/cpu"
	"github.com/ossareh/gsssd/mem"
	"github.com/ossareh/libgosysstat/core"
	cpuProcessor "github.com/ossareh/libgosysstat/processor/cpu"
	memProcessor "github.com/ossareh/libgosysstat/processor/mem"
)

const TICK_INTERVAL = 5

type StatSender interface {
	Send([]core.Stat)
}

var (
	statsdAddress = flag.String("address", "", "The address of the statsd instance to send data to")
	statsdPort    = flag.Int("port", 8125, "the port statsd is listening on")
	statsdPrefix  = flag.String("prefix", "", "host specific prefix")
)

func main() {
	flag.Parse()
	if *statsdAddress == "" || *statsdPrefix == "" {
		log.Fatal(errors.New("Address and Prefix arguments are required"))
	}
	address := fmt.Sprintf("%s:%d", *statsdAddress, *statsdPort)

	client, err := statsd.New(address, *statsdPrefix)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	/**
	 * CPU Setup
	 */
	cpuClient := cpu.New(client)
	procStat, err := os.Open("/proc/stat")
	if err != nil {
		log.Fatal(err)
	}
	defer procStat.Close()
	cpuStatProcessor := cpuProcessor.New(procStat)
	cpuStatResults := make(chan []core.Stat)
	go core.StatProcessor(cpuStatProcessor, TICK_INTERVAL, cpuStatResults)

	/**
	 * Memory Setup
	 */
	memClient := mem.New(client)
	memStat, err := os.Open("/proc/meminfo")
	if err != nil {
		log.Fatal(err)
	}
	defer memStat.Close()
	memStatProcessor := memProcessor.New(memStat)
	memStatResults := make(chan []core.Stat)
	go core.StatProcessor(memStatProcessor, TICK_INTERVAL, memStatResults)

	/**
	 * Main loop
	 */
	for {
		select {
		case stats := <-cpuStatResults:
			cpuClient.Send(stats)
		case stats := <-memStatResults:
			memClient.Send(stats)
		}
	}
}
