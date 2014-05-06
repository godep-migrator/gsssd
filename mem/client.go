package mem

import (
	"github.com/cactus/go-statsd-client/statsd"
	"github.com/ossareh/libgosysstat/core"
)

type GaugeData struct {
	prefix string
	value  uint64
}

type MemClient struct {
	client statsd.Statter
}

func (c *MemClient) prep(stats []core.Stat) []GaugeData {
	vals := []GaugeData{}
	for _, s := range stats {
		value := s.Values()[0]
		switch s.Type() {
		case "total":
			vals = append(vals, GaugeData{"mem.main.total", value})
		case "used":
			vals = append(vals, GaugeData{"mem.main.used", value})
		case "cached":
			vals = append(vals, GaugeData{"mem.main.cached", value})
		case "swap_total":
			if value > 0 {
				vals = append(vals, GaugeData{"mem.swap.total", value})
			}
		case "swap_used":
			if value > 0 {
				vals = append(vals, GaugeData{"mem.swap.used", value})
			}
		}
	}
	return vals
}

func (c *MemClient) Send(stats []core.Stat) {
	values := c.prep(stats)
	for _, data := range values {
		c.client.Gauge(data.prefix, int64(data.value), 1.0)
	}
}

func New(client statsd.Statter) *MemClient {
	return &MemClient{client}
}
