package ui

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Ui struct {
	plot *widgets.Plot
}

func (appWindow *Ui) Init() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	appWindow.plot = widgets.NewPlot()
	appWindow.plot.Title = "ping wp.pl"

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}

func (appWindow *Ui) Draw(data []int64) {

	if len(data) > 100 {
		data = data[len(data)-100 : len(data)-1]
	}
	floatData := []float64{0}
	for _, element := range data {
		floatData = append(floatData, float64(element))
	}

	appWindow.plot.Data = [][]float64{floatData}
	appWindow.plot.SetRect(0, 0, 100, 15)
	appWindow.plot.AxesColor = ui.ColorWhite
	appWindow.plot.LineColors[0] = ui.ColorGreen

	ui.Render(appWindow.plot)
}
