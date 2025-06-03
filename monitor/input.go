package monitor

import (
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// InputHandler returns a tview input capture handler for the main app.
func InputHandler(ctx *AppContext) func(event *tcell.EventKey) *tcell.EventKey {
	inputField := tview.NewInputField().
		SetLabel("PID to kill: ").
		SetFieldWidth(10)
	return func(event *tcell.EventKey) *tcell.EventKey {
		switch {
		case event.Rune() == 'q' || event.Key() == tcell.KeyEscape:
			ctx.App.Stop()
			return nil
		case event.Rune() == 'k':
			row, _ := ctx.ProcTable.GetSelection()
			if row > 0 {
				pidCell := ctx.ProcTable.GetCell(row, 0)
				pidStr := pidCell.Text
				pidStr = strings.TrimLeft(pidStr, "[red]")
				if pid, err := strconv.Atoi(pidStr); err == nil {
					err := KillProcessByPID(int32(pid))
					if err == nil {
						ctx.CPUMemView.SetText("[red]Process " + pidStr + " killed successfully")
					} else {
						ctx.CPUMemView.SetText("[red]Failed to kill process: " + err.Error())
					}
				}
			}
			return nil
		case event.Rune() == 'i':
			inputField.SetText("")
			form := tview.NewForm().
				AddFormItem(inputField).
				AddButton("Kill", func() {
					pidStr := inputField.GetText()
					if pid, err := strconv.Atoi(pidStr); err == nil {
						err := KillProcessByPID(int32(pid))
						if err == nil {
							ctx.CPUMemView.SetText("[red]Process " + pidStr + " killed successfully")
						} else {
							ctx.CPUMemView.SetText("[red]Failed to kill process: " + err.Error())
						}
					}
					ctx.App.SetRoot(ctx.Flex, true)
				}).
				AddButton("Cancel", func() {
					ctx.App.SetRoot(ctx.Flex, true)
				})
			form.SetBorder(true).SetTitle("Kill Process by PID").SetTitleAlign(tview.AlignLeft)
			ctx.App.SetRoot(form, true).SetFocus(form)
			return nil
		}
		return event
	}
}
