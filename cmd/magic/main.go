package main

import (
	"runtime"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func currentTime() []fyne.CanvasObject {
	dateLabels := make([]fyne.CanvasObject, 5)
	for i := range dateLabels {
		dateLabels[i] = &widget.Label{Text: "", Alignment: fyne.TextAlignCenter}
	}
	ticker := time.NewTicker(1 * time.Second)
	go func(ticker *time.Ticker) {
		defer ticker.Stop()
		startPosition := 0
		changePosition := false
		for t := range ticker.C {
			println("second")
			for i := range dateLabels {
				if i == startPosition {
					dateLabels[i].(*widget.Label).SetText(t.Format("Mon, 02 Jan 2006"))
				} else if i == startPosition + 1 {
					dateLabels[i].(*widget.Label).SetText(t.Format("15:04:05"))
				} else {
					dateLabels[i].(*widget.Label).SetText("")
				}
				widget.Refresh(dateLabels[i].(*widget.Label))
			}
			//dateLabel.SetText(t.Format("Mon, 02 Jan 2006"))
			//timeLabel.SetText(t.Format("15:04:05"))
			//widget.Refresh(dateLabel)
			//widget.Refresh(timeLabel)
			if changePosition {
				startPosition = (startPosition + 1) % 4
			}
			changePosition = !changePosition
		}
	}(ticker)
	return dateLabels
}

func main() {
	a := app.New()

	dateAndTime := currentTime()
	w := a.NewWindow("Hello")
	w.SetContent(&widget.Box{Children:
	dateAndTime,
	})
	w.Resize(fyne.Size{
		Width:  600,
		Height: 1024,
	})
	if runtime.GOOS == "linux" {
		w.SetFullScreen(true)
	}
	w.ShowAndRun()
}
