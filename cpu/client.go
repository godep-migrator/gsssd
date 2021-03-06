package cpu

import (
	"fmt"

	"github.com/cactus/go-statsd-client/statsd"
	"github.com/ossareh/libgosysstat/core"
	"github.com/ossareh/libgosysstat/processor/cpu"
)

type GaugeData struct {
	prefix string
	value  uint64
}

type CpuClient struct {
	client statsd.Statter
}

func prepareCpuValues(values []uint64) (user, sys, idle, io uint64) {
	user = values[0] + values[1]
	sys = values[2]
	idle = values[3]
	io = values[4]
	return
}

func (c *CpuClient) prep(stats []core.Stat) []GaugeData {
	vals := []GaugeData{}
	for _, s := range stats {
		values := s.Values()
		switch s.Type() {
		case cpu.TOTAL:
			user, sys, idle, io := prepareCpuValues(values)
			vals = append(vals, GaugeData{"cpu.total.user", user})
			vals = append(vals, GaugeData{"cpu.total.sys", sys})
			vals = append(vals, GaugeData{"cpu.total.idle", idle})
			vals = append(vals, GaugeData{"cpu.total.io", io})
		case cpu.INTR:
			vals = append(vals, GaugeData{"cpu.interrupts", values[0]})
		case cpu.CTXT:
			vals = append(vals, GaugeData{"cpu.context_switches", values[0]})
		case cpu.PROCS:
			vals = append(vals, GaugeData{"cpu.processes.created", values[0]})
		case cpu.PROCS_RUNNING:
			vals = append(vals, GaugeData{"cpu.processes.running", values[0]})
		case cpu.PROCS_BLOCKED:
			vals = append(vals, GaugeData{"cpu.processes.blocked", values[0]})
		default:
			// CPU
			user, sys, idle, io := prepareCpuValues(values)
			prefix := fmt.Sprintf("cpu.%s", s.Type())
			vals = append(vals, GaugeData{prefix + ".user", user})
			vals = append(vals, GaugeData{prefix + ".sys", sys})
			vals = append(vals, GaugeData{prefix + ".idle", idle})
			vals = append(vals, GaugeData{prefix + ".io", io})
		}
	}
	return vals
}

func (c *CpuClient) Send(stats []core.Stat) {
	values := c.prep(stats)
	for _, data := range values {
		c.client.Gauge(data.prefix, int64(data.value), 1.0)
	}
}

func New(client statsd.Statter) *CpuClient {
	return &CpuClient{client}
}
