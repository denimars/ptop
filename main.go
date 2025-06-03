package main

import (
	"ptop/monitor"
)

func main() {
	ctx := monitor.SetupUI()
	monitor.StartCPUMemUpdateLoop(ctx)
	monitor.StartProcessUpdateLoop(ctx)
	ctx.App.SetInputCapture(monitor.InputHandler(ctx))

	if err := ctx.App.SetRoot(ctx.Flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
