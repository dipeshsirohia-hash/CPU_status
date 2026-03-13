package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
)

type ProcInfo struct {
	Name string 
	CPU  float64
	Mem  float64
}

func main() {
	for {
		fmt.Println("===================================")
		fmt.Println("System Usage @", time.Now().Format("15:04:05"))
		fmt.Println("===================================")

		// CPU
		cpuPercent, _ := cpu.Percent(0, false)
		fmt.Printf("CPU Usage: %.2f%%\n", cpuPercent[0])

		// Memory

		vmem, _ := mem.VirtualMemory()
		fmt.Printf("Memory Usage: %.2f%% (Used: %.2f GB / Total: %.2f GB)\n",
			vmem.UsedPercent,
			float64(vmem.Used)/1024/1024/1024,
			float64(vmem.Total)/1024/1024/1024)

		fmt.Println("\nTop Processes (by Memory):")

		processes, _ := process.Processes()
		var list []ProcInfo

		for _, p := range processes {
			name, err := p.Name()
			if err != nil {
				continue
			}

			// memPercent, _ := p.MemoryPercent()
			memInfo, err := p.MemoryInfo()
			if err != nil {
				continue
			}

			memMB := float64(memInfo.RSS) / 1024 / 1024
			cpuPercent, _ := p.CPUPercent()

			list = append(list, ProcInfo{
				Name: name,
				CPU:  cpuPercent,
				Mem:  memMB,
			})
		}

		// Sort by memory usage
		sort.Slice(list, func(i, j int) bool {
			return list[i].Mem > list[j].Mem
		})

		// Print top 10
		for i := 0; i < 20 && i < len(list); i++ {
			fmt.Printf("%-25s CPU: %6.2f%%  MEM: %6.2f Mb\n",
				list[i].Name,
				list[i].CPU,
				list[i].Mem,
			)
		}

		time.Sleep(10 * time.Second)
	}
}
