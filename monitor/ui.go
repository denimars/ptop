package monitor

import (
	"github.com/rivo/tview"
)

// SetupUI initializes the main UI components and returns a context struct.
func SetupUI() *AppContext {
	app := tview.NewApplication()
	cpuMemView := tview.NewTextView().
		SetDynamicColors(true).
		SetChangedFunc(func() { app.Draw() })
	procTable := tview.NewTable().
		SetBorders(false).
		SetFixed(1, 0)
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(cpuMemView, 2, 1, false).
		AddItem(procTable, 0, 6, true)
	return &AppContext{
		App:        app,
		CPUMemView: cpuMemView,
		ProcTable:  procTable,
		Flex:       flex,
	}
}

// UpdateProcTable updates the process table UI with the provided process list.
func UpdateProcTable(procTable *tview.Table, procList []ProcInfo) {
	procTable.Clear()
	procTable.SetSelectable(true, false)
	headers := []string{"PID", "Name", "CPU%", "Mem%", "Port"}
	for i, h := range headers {
		procTable.SetCell(0, i, tview.NewTableCell("[::b]"+h))
	}
	maxRows := min(50, len(procList))
	for i := 0; i < maxRows; i++ {
		p := procList[i]
		procTable.SetCell(i+1, 0, tview.NewTableCell("[red]"+p.PIDStr))
		procTable.SetCell(i+1, 1, tview.NewTableCell(p.Name))
		procTable.SetCell(i+1, 2, tview.NewTableCell("[yellow]"+p.CPUStr))
		procTable.SetCell(i+1, 3, tview.NewTableCell("[yellow]"+p.MemStr))
		procTable.SetCell(i+1, 4, tview.NewTableCell(p.Ports))
	}
}
