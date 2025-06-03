package monitor

import (
	"github.com/rivo/tview"
)

type AppContext struct {
	App        *tview.Application
	CPUMemView *tview.TextView
	ProcTable  *tview.Table
	Flex       *tview.Flex
}
