package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"
)

type SystemStats struct {
	CPUUsage          float64
	MemoryUsage       uint64
	MemoryTotal       uint64
	MemoryPercent     float64
	NetworkInterfaces []net.InterfaceStat
	Platform          string
	PlatformVersion   string
	Uptime            uint64
}

func getSystemStats() (SystemStats, error) {
	cpuPercents, err := cpu.Percent(time.Second, false)
	if err != nil {
		return SystemStats{}, err
	}

	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return SystemStats{}, err
	}

	hostInfo, err := host.Info()
	if err != nil {
		return SystemStats{}, err
	}

	stats := SystemStats{
		CPUUsage:        cpuPercents[0],
		MemoryUsage:     vmStat.Used,
		MemoryTotal:     vmStat.Total,
		MemoryPercent:   vmStat.UsedPercent,
		Platform:        hostInfo.Platform,
		PlatformVersion: hostInfo.PlatformVersion,
		Uptime:          hostInfo.Uptime / 3600,
	}

	return stats, nil
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	stats, err := getSystemStats()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting system stats: %v", err), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing template: %v", err), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, stats)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing template: %v", err), http.StatusInternalServerError)
		return
	}
}

func apiStatsHandler(w http.ResponseWriter, r *http.Request) {
	stats, err := getSystemStats()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting system stats: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

func main() {
	http.HandleFunc("/stats", statsHandler)
	http.HandleFunc("/api/stats", apiStatsHandler)
	log.Println("Starting server on :7777")
	log.Fatal(http.ListenAndServe(":7777", nil))
}
