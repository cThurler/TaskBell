package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func setupUI(w fyne.Window) {
	tasks := loadTasks()

	taskList := widget.NewList(
		func() int { return len(tasks) },
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(tasks[i].Title)
		},
	)

	w.SetContent(container.NewBorder(nil, nil, nil, nil, taskList))
}
