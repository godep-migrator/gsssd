package cpu

import (
	"log"
	"reflect"
	"testing"

	"github.com/cactus/go-statsd-client/statsd"
	"github.com/ossareh/libgosysstat/core"
	"github.com/ossareh/libgosysstat/processor/cpu"
)

type TestingStat struct {
	t    string
	vals []uint64
}

func (t *TestingStat) Type() string {
	return t.t
}

func (t *TestingStat) Values() []uint64 {
	return t.vals
}

func TestStatPrep(t *testing.T) {
	noop, _ := statsd.NewNoop("123.123.123.123", "test.prefix")
	client := &CpuClient{noop}
	expected := []GaugeData{
		GaugeData{"cpu.total.user", 3},
		GaugeData{"cpu.total.sys", 3},
		GaugeData{"cpu.total.idle", 4},
		GaugeData{"cpu.total.io", 5},
		GaugeData{"cpu.cpu0.user", 33},
		GaugeData{"cpu.cpu0.sys", 33},
		GaugeData{"cpu.cpu0.idle", 44},
		GaugeData{"cpu.cpu0.io", 55},
		GaugeData{"cpu.cpu1.user", 35},
		GaugeData{"cpu.cpu1.sys", 34},
		GaugeData{"cpu.cpu1.idle", 45},
		GaugeData{"cpu.cpu1.io", 56},
		GaugeData{"cpu.cpu2.user", 22},
		GaugeData{"cpu.cpu2.sys", 14},
		GaugeData{"cpu.cpu2.idle", 16},
		GaugeData{"cpu.cpu2.io", 18},
		GaugeData{"cpu.cpu3.user", 0},
		GaugeData{"cpu.cpu3.sys", 1},
		GaugeData{"cpu.cpu3.idle", 100},
		GaugeData{"cpu.cpu3.io", 10},
		GaugeData{"cpu.cpu4.user", 4},
		GaugeData{"cpu.cpu4.sys", 5},
		GaugeData{"cpu.cpu4.idle", 7},
		GaugeData{"cpu.cpu4.io", 9},
		GaugeData{"cpu.cpu5.user", 6},
		GaugeData{"cpu.cpu5.sys", 6},
		GaugeData{"cpu.cpu5.idle", 8},
		GaugeData{"cpu.cpu5.io", 10},
		GaugeData{"cpu.cpu6.user", 5},
		GaugeData{"cpu.cpu6.sys", 5},
		GaugeData{"cpu.cpu6.idle", 7},
		GaugeData{"cpu.cpu6.io", 11},
		GaugeData{"cpu.cpu7.user", 2},
		GaugeData{"cpu.cpu7.sys", 2},
		GaugeData{"cpu.cpu7.idle", 3},
		GaugeData{"cpu.cpu7.io", 5},
		GaugeData{"cpu.interrupts", 10102},
		GaugeData{"cpu.context_switches", 1010101},
		GaugeData{"cpu.processes.created", 23},
		GaugeData{"cpu.processes.running", 123},
		GaugeData{"cpu.processes.blocked", 3},
	}

	results := client.prep([]core.Stat{
		&TestingStat{cpu.TOTAL, []uint64{1, 2, 3, 4, 5}},
		&TestingStat{"cpu0", []uint64{11, 22, 33, 44, 55}},
		&TestingStat{"cpu1", []uint64{12, 23, 34, 45, 56}},
		&TestingStat{"cpu2", []uint64{10, 12, 14, 16, 18}},
		&TestingStat{"cpu3", []uint64{0, 0, 1, 100, 10}},
		&TestingStat{"cpu4", []uint64{1, 3, 5, 7, 9}},
		&TestingStat{"cpu5", []uint64{2, 4, 6, 8, 10}},
		&TestingStat{"cpu6", []uint64{2, 3, 5, 7, 11}},
		&TestingStat{"cpu7", []uint64{1, 1, 2, 3, 5}},
		&TestingStat{cpu.INTR, []uint64{10102}},
		&TestingStat{cpu.CTXT, []uint64{1010101}},
		&TestingStat{cpu.PROCS, []uint64{23}},
		&TestingStat{cpu.PROCS_RUNNING, []uint64{123}},
		&TestingStat{cpu.PROCS_BLOCKED, []uint64{3}},
	})

	if !reflect.DeepEqual(expected, results) {
		for idx, e := range expected {
			if !reflect.DeepEqual(e, results[idx]) {
				log.Fatalf("Expected item %d to be %s, got %s", idx, e, results[idx])
			}
		}
	}
}

func TestPrepareCpuValues(t *testing.T) {
	expected_user := uint64(51)
	expected_sys := uint64(1)
	expected_idle := uint64(49)
	expected_io := uint64(3)
	known := []uint64{1, 50, 1, 49, 3}
	user, sys, idle, io := prepareCpuValues(known)

	if user != expected_user {
		t.Fatalf("Expected %s to be %d, got %d", "user", expected_user, user)
	}
	if sys != expected_sys {
		t.Fatalf("Expected %s to be %d, got %d", "sys", expected_sys, sys)
	}
	if idle != expected_idle {
		t.Fatalf("Expected %s to be %d, got %d", "idle", expected_idle, idle)
	}
	if io != expected_io {
		t.Fatalf("Expected %s to be %d, got %d", "io", expected_io, io)
	}
}
