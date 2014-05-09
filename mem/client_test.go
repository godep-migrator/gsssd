package mem

import (
	"log"
	"reflect"
	"testing"

	"github.com/cactus/go-statsd-client/statsd"
	"github.com/ossareh/libgosysstat/core"
)

type TestingStat struct {
	t   string
	val uint64
}

func (t *TestingStat) Type() string {
	return t.t
}

func (t *TestingStat) Values() []uint64 {
	return []uint64{t.val}
}

func TestStatPrepWithSwap(t *testing.T) {
	noop, _ := statsd.NewNoop("123.123.123.123", "test.prefix")
	client := &MemClient{noop}
	expected := []GaugeData{
		GaugeData{"mem.main.total", (24684396 * 1024)},
		GaugeData{"mem.main.used", (3867420 * 1024)},
		GaugeData{"mem.main.cached", (651036 * 1024)},
		GaugeData{"mem.swap.total", (1000 * 1024)},
		GaugeData{"mem.swap.used", (500 * 1024)},
	}

	results := client.prep([]core.Stat{
		&TestingStat{"total", 24684396},
		&TestingStat{"used", 3867420},
		&TestingStat{"cached", 651036},
		&TestingStat{"swap_total", 1000},
		&TestingStat{"swap_used", 500},
	})

	if !reflect.DeepEqual(expected, results) {
		for idx, e := range expected {
			if !reflect.DeepEqual(e, results[idx]) {
				log.Fatalf("Expected item %d to be %v, got %v", idx, e, results[idx])
			}
		}
	}
}

func TestStatPrepNoSwap(t *testing.T) {
	noop, _ := statsd.NewNoop("123.123.123.123", "test.prefix")
	client := &MemClient{noop}
	expected := []GaugeData{
		GaugeData{"mem.main.total", (24684396 * 1024)},
		GaugeData{"mem.main.used", (3867420 * 1024)},
		GaugeData{"mem.main.cached", (651036 * 1024)},
	}

	results := client.prep([]core.Stat{
		&TestingStat{"total", 24684396},
		&TestingStat{"used", 3867420},
		&TestingStat{"cached", 651036},
		&TestingStat{"swap_total", 0},
	})

	if !reflect.DeepEqual(expected, results) {
		for idx, e := range expected {
			if !reflect.DeepEqual(e, results[idx]) {
				log.Fatalf("Expected item %d to be %v, got %v", idx, e, results[idx])
			}
		}
	}
}
