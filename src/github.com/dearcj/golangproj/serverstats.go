package main

import (
	"context"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"go.opencensus.io/exporter/prometheus"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"log"
	"sync/atomic"
	//"time"
)

type ServerMetrics struct {
	//CPULoad *stats.MeasureFloat64
	MemUsed            *stats.Int64Measure
	MemUsedPct         *stats.Float64Measure
	CPULoadPct         *stats.Float64Measure
	SocketRead         *stats.Int64Measure
	SocketWrite        *stats.Int64Measure
	PacketSize         *stats.Int64Measure
	SocketReadCounter  int64
	SocketWriteCounter int64
}

type ServerStats struct {
	ctx      context.Context
	Metrics  ServerMetrics
	Exporter *prometheus.Exporter
}

func (s *ServerStats) EverySec() {
	stats.Record(context.Background(), s.Metrics.SocketWrite.M(atomic.LoadInt64(&s.Metrics.SocketWriteCounter)))
	stats.Record(context.Background(), s.Metrics.SocketRead.M(atomic.LoadInt64(&s.Metrics.SocketReadCounter)))

	atomic.StoreInt64(&s.Metrics.SocketReadCounter, 0)
	atomic.StoreInt64(&s.Metrics.SocketWriteCounter, 0)
}

func (s *ServerStats) OnServerLoop() {
	v, _ := mem.VirtualMemory()
	x, _ := cpu.Percent(0, false)

	stats.Record(s.ctx,
		s.Metrics.CPULoadPct.M(v.UsedPercent),
		s.Metrics.MemUsed.M(int64(v.Used)),
		s.Metrics.MemUsedPct.M(x[0]),
	)
}

func InitServerStats() *ServerStats {
	exporter, err := prometheus.NewExporter(prometheus.Options{})

	if err != nil {
		log.Fatal(err)
	}

	view.RegisterExporter(exporter)
	ss := &ServerStats{
		Exporter: exporter,
		ctx:      context.Background(),
	}

	MemUsed := stats.Int64("mem_used", "mem used", "KB")
	MemUsedPct := stats.Float64("mem_used_pct", "mem used percentage", "%")
	CPULoadPct := stats.Float64("cpu_load_pct", "cpu load percentage", "%")
	SocketRead := stats.Int64("num_sockets_read", "sockets read per sec", "num")
	SocketWrite := stats.Int64("num_sockets_write", "sockets write per sec", "num")
	PacketSize := stats.Int64("packet_size", "packet avg size in bytes", "num")

	view.Register(&view.View{
		Name:        "PacketSize",
		Description: "Size of server message in bytes",
		Measure:     PacketSize,
		Aggregation: view.Distribution(0, 100, 200, 300, 500, 1000, 1500, 2000),
	})

	view.Register(&view.View{
		Name:        "Socket Read",
		Description: "Sockets Read",
		Measure:     SocketRead,
		Aggregation: view.Distribution(0, 100, 200, 300, 500, 1000, 1500, 2000),
	})

	/*	ss.MakeView(SocketRead, stats.SumAggregation{},
				stats.Cumulative{})
		//		stats.Interval{Intervals: 10, Duration: 10 * time.Second})
			ss.MakeView(SocketWrite, stats.SumAggregation{},
				stats.Cumulative{})

			ss.MakeView(MemUsed, stats.SumAggregation{},
				stats.Cumulative{})
			ss.MakeView(MemUsedPct, stats.SumAggregation{},
				stats.Cumulative{})
			ss.MakeView(CPULoadPct, stats.SumAggregation{},
				stats.Cumulative{})
			ss.MakeView(PacketSize, stats.SumAggregation{},
				stats.Cumulative{})
	*/

	view.SetReportingPeriod(nodeConfig.Server.StatsReportDelay)

	ss.Metrics = ServerMetrics{
		CPULoadPct:  CPULoadPct,
		MemUsed:     MemUsed,
		MemUsedPct:  MemUsedPct,
		SocketRead:  SocketRead,
		SocketWrite: SocketWrite,
		PacketSize:  PacketSize,
	}

	return ss
}
