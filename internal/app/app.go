package starter

import (
	interface_app "myApp/internal/interface"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func ShowAndRun(i *interface_app.Interface_app) {

	i.Window.Resize(fyne.NewSize(500, 500))
	i.Window.CenterOnScreen()

	t := i.AddTabs() 
	t.SetTabLocation(container.TabLocationLeading)
	container := container.New(layout.NewMaxLayout(), t)
	i.Window.SetContent(container)
	i.Window.Show()
	i.App.Run()
}
