package monitor

import (
	"github.com/rivo/tview"
)

// AppContext holds the main UI components and shared state
// for the procptop/monitor application.
type AppContext struct {
	App        *tview.Application
	CPUMemView *tview.TextView
	ProcTable  *tview.Table
	Flex       *tview.Flex
}
