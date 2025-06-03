package monitor

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

func StartCPUMemUpdateLoop(ctx *AppContext) {
	go func() {
		for {
			cpuPercent, _ := cpu.Percent(0, false)
			vmStat, _ := mem.VirtualMemory()
			cpuMem := fmt.Sprintf("[yellow]CPU:[white] %.2f%%   [yellow]Memory:[white] %.2f%%",
				cpuPercent[0], vmStat.UsedPercent)
			ctx.App.QueueUpdateDraw(func() {
				ctx.CPUMemView.SetText(cpuMem)
			})
			time.Sleep(2 * time.Second)
		}
	}()
}

// StartProcessUpdateLoop periodically refreshes the process table in the UI.
func StartProcessUpdateLoop(ctx *AppContext) {
	go func() {
		for {
			processes, portMap, err := GetProcessList()
			if err != nil {
				continue
			}
			procList := BuildProcInfoList(processes, portMap)
			SortProcInfoList(procList)
			ctx.App.QueueUpdateDraw(func() {
				UpdateProcTable(ctx.ProcTable, procList)
			})
			time.Sleep(2 * time.Second)
		}
	}()
}
