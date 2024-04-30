package starter

import (
	interface_app "myApp/internal/interface"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func ShowAndRun(i *interface_app.Interface_app) {
	a := app.New()
	w := a.NewWindow("Анализатор")
	w.Resize(fyne.NewSize(500, 500))
	w.CenterOnScreen()
	t := i.AddTabs()
	t.SetTabLocation(container.TabLocationLeading)
	container := container.New(layout.NewMaxLayout(), t)
	w.SetContent(container)
	w.Show()
	a.Run()
}
