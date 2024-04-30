package interface_app

import (
	"fyne.io/fyne/v2/container"
)

func (i *Interface_app) AddTabs() *container.AppTabs {
	return container.NewAppTabs(
		container.NewTabItem("Ресурсы", i.ResourcesTable()),
		container.NewTabItem("Задолженности студентов", i.Students()),
		container.NewTabItem("Добавить Студента", i.StudManager()),
		container.NewTabItem("Управление ресурсами", i.Resources()),
	)
}
