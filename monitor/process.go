package monitor

import (
	"fmt"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	gopsnet "github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
)

type ProcInfo struct {
	PID    int32
	PIDStr string
	Name   string
	CPU    float64
	Mem    float64
	CPUStr string
	MemStr string
	Ports  string
}

func GetProcessList() ([]*process.Process, map[int32][]string, error) {
	processes, err := process.Processes()
	if err != nil {
		return nil, nil, err
	}
	portMap := map[int32][]string{}
	connections, _ := gopsnet.Connections("tcp")
	for _, conn := range connections {
		if conn.Status == "LISTEN" {
			port := strconv.Itoa(int(conn.Laddr.Port))
			portMap[conn.Pid] = append(portMap[conn.Pid], port)
		}
	}
	return processes, portMap, nil
}

func BuildProcInfoList(processes []*process.Process, portMap map[int32][]string) []ProcInfo {
	var procList []ProcInfo
	for _, p := range processes {
		pid := p.Pid
		pidStr := strconv.Itoa(int(pid))
		name := "Unknown"
		if n, err := p.Name(); err == nil {
			name = n
		} else if exe, err := p.Exe(); err == nil {
			name = filepath.Base(exe)
		}
		cpu := 0.0
		cpuStr := ""
		if currentCPU, err := p.CPUPercent(); err == nil {
			cpu = currentCPU
			cpuStr = fmt.Sprintf("%.1f", currentCPU)
		} else {
			cpuStr = "err"
		}
		mem := 0.0
		memStr := ""
		if mp, err := p.MemoryPercent(); err == nil {
			mem = float64(mp)
			memStr = fmt.Sprintf("%.1f", mp)
		} else {
			memStr = "err"
		}
		ports := strings.Join(portMap[pid], ",")
		procList = append(procList, ProcInfo{pid, pidStr, name, cpu, mem, cpuStr, memStr, ports})
	}
	return procList
}

func SortProcInfoList(procList []ProcInfo) {
	sort.Slice(procList, func(i, j int) bool {
		if procList[i].CPU == procList[j].CPU {
			return procList[i].Mem > procList[j].Mem
		}
		return procList[i].CPU > procList[j].CPU
	})
}

func KillProcessByPID(pid int32) error {
	proc, err := process.NewProcess(pid)
	if err != nil {
		return err
	}
	return proc.Kill()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
