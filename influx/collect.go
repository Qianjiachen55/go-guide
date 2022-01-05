package main

import (
	"fmt"
	client "github.com/influxdata/influxdb1-client/v2"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"log"
	"time"
)

const (
	CpuInfoType  = "cpu"
	MemInfoType  = "mem"
	DistInfoType = "disk"
	NetInfoType  = "net"
)

type SysInfo struct {
	InfoType string
	IP       string
	Data     interface{}
}

type CpuInfo struct {
	CpuPercent float64 `json:"cpuPercent"`
}

type MemInfo struct {
	Total       uint64  `json:"total"`
	Available   uint64  `json:"available"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
	Buffers     uint64  `json:"buffers"`
	Cached      uint64  `json:"cached"`
}


var cli client.Client

func initConnInflux() (err error) {
	cli, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://127.0.0.1:8086",
		Username: "q",
		Password: "",
	})
	return
}

func writesCpuPoints(data *CpuInfo) {


	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "test",
		Precision: "s", //精度，默认ns
	})
	if err != nil {
		log.Fatal(err)
	}

	tags := map[string]string{"cpu": "cpu 0"}
	fields := map[string]interface{}{
		"cpuPercent": data.CpuPercent,
	}

	pt, err := client.NewPoint("cpu_percent", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)
	err = cli.Write(bp)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("insert success")
}

func writesMemPoints(data *MemInfo) {

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "test",
		Precision: "s", //精度，默认ns
	})
	if err != nil {
		log.Fatal(err)
	}

	tags := map[string]string{"mem": "mem"}
	fields := map[string]interface{}{
		"total":       int(data.Total),
		"available":   int(data.Available),
		"used":        int(data.Used),
		"usedPercent": int(data.UsedPercent),
		"buffers":     int(data.Buffers),
		"cached":      int(data.Cached),
	}

	pt, err := client.NewPoint("mem", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)
	err = cli.Write(bp)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("insert success")
}

func getCpuInfo() {
	var cpuInfo = new(CpuInfo)
	percent, _ := cpu.Percent(time.Second, false)
	fmt.Printf("cpu percent:%v\n", percent)
	//
	cpuInfo.CpuPercent = percent[0]
	writesCpuPoints(cpuInfo)
}

func getMemInfo()  {
	var memInfo = new(MemInfo)
	info, err := mem.VirtualMemory()
	if err != nil{
		fmt.Println("getMemInfo err:",err)
	}
	memInfo.Cached = info.Cached
	memInfo.Buffers = info.Buffers
	memInfo.UsedPercent = info.UsedPercent
	memInfo.Used = info.Used
	memInfo.Total = info.Total
	memInfo.Available = info.Available

	writesMemPoints(memInfo)
}

func run(interval time.Duration)  {
	ticker := time.Tick(interval)
	for _= range ticker{
		getCpuInfo()
		getMemInfo()
	}

}


func main() {
	err := initConnInflux()
	if err != nil {
		fmt.Println("init err!")
	}
	run(time.Second/2)
}
