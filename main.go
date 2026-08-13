package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// borderDemo()
	// formDemo()
}

func formDemo() {
	a := app.NewWithID("form")
	w := a.NewWindow("Form Demo")

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Name")

	themeSelect := widget.NewSelect([]string{"light", "dark"}, func(s string) {
		fmt.Println("chosen:", s)
	})

	notifyCheck := widget.NewCheck("Enable notifs", func(b bool) {
		fmt.Println("notif", b)
	})

	form := widget.NewForm(
		widget.NewFormItem("Name", nameEntry),
		widget.NewFormItem("Theme", themeSelect),
		widget.NewFormItem("", notifyCheck),
	)

	w.SetContent(form)
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}

func borderDemo() {
	a := app.NewWithID("border")
	w := a.NewWindow("Border Demo")

	top := widget.NewLabel("This is the top")
	bottom := widget.NewButton("Save", func() {})
	center := widget.NewLabel("This is the center")

	content := container.NewBorder(top, bottom, nil, nil, center)
	w.SetContent(content)
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}
