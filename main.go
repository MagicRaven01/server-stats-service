package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

type SystemStats struct {
	CPUUsage        float64
	MemoryUsage     uint64
	MemoryTotal     uint64
	MemoryPercent   float64
	Platform        string
	PlatformVersion string
	Uptime          uint64
}

var tmpl = template.Must(template.ParseFiles("template.html"))

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
		log.Printf("Error getting stats: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, stats); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func apiStatsHandler(w http.ResponseWriter, r *http.Request) {
	stats, err := getSystemStats()
	if err != nil {
		log.Printf("Error getting stats: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(stats); err != nil {
		log.Printf("Error encoding JSON: %v", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/stats", statsHandler)
	mux.HandleFunc("/api/stats", apiStatsHandler)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Starting server on :7777")
	log.Fatal(http.ListenAndServe(":7777", mux))
}
