package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func DebugPrintStats() {
	printMemoryStats()
	startTime := time.Now()
	fmt.Printf("running the GC ---> ")
	runtime.GC()
	fmt.Printf("took %v\n", time.Since(startTime))
	printMemoryStats()
	fmt.Println()
}

func printMemoryStats() {
	memstats := runtime.MemStats{}
	runtime.ReadMemStats(&memstats)
	fmt.Printf("HeapInuse: %s\n", FormatSize(memstats.HeapInuse))
	fmt.Printf("TotalAlloc: %s\n", FormatSize(memstats.TotalAlloc))
	fmt.Printf("HeapIdle: %s\n", FormatSize(memstats.HeapIdle))
	fmt.Printf("HeapReleased: %s\n", FormatSize(memstats.HeapReleased))
	fmt.Printf("HeapAlloc: %s\n", FormatSize(memstats.HeapAlloc))
	fmt.Printf("Sys: %s\n", FormatSize(memstats.Sys))
	fmt.Printf("Mallocs: %d, Frees: %d, NumGC: %d\n", memstats.Mallocs, memstats.Frees, memstats.NumGC)

	ps, err := getProcMemStat()
	if err != nil {
		fmt.Printf("unable to read process stats: %s\n", err.Error())
	} else {
		fmt.Printf("Process memory stats\n")
		fmt.Printf(
			"PageSize: %d, Virt: %s, RSS: %s, Shared: %s\n",
			ps.PageSize, FormatSize(ps.VirtSizeBytes),
			FormatSize(ps.RSSBytes), FormatSize(ps.SharedBytes),
		)
	}
}

type procMemStat struct {
	PageSize      uint64
	VirtSizeBytes uint64
	RSSBytes      uint64
	SharedBytes   uint64
}

func getProcMemStat() (procMemStat, error) {
	buf, err := ioutil.ReadFile("/proc/self/statm")
	if err != nil {
		return procMemStat{}, err
	}

	fields := strings.Split(string(buf), " ")
	if len(fields) < 3 {
		return procMemStat{}, errors.New("can't parse statm")
	}

	total, _ := strconv.ParseUint(fields[0], 10, 64)
	rss, _ := strconv.ParseUint(fields[1], 10, 64)
	shared, _ := strconv.ParseUint(fields[2], 10, 64)
	pageSize := uint64(os.Getpagesize())

	return procMemStat{
		PageSize:      pageSize,
		VirtSizeBytes: total * pageSize,
		RSSBytes:      rss * pageSize,
		SharedBytes:   shared * pageSize,
	}, nil
}

func FormatSize(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB",
		float64(b)/float64(div), "KMGTPE"[exp])
}
