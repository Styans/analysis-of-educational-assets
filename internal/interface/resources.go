package interface_app

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (i *Interface_app) ResourcesTable() *fyne.Container {

	types, err := i.service.ResourcesService.GetResourcesTypes()
	if err != nil {
		fmt.Println(err)
	}
	sel := widget.NewSelect(
		types,
		nil,
	)
	t := widget.NewTable(

		func() (int, int) {
			return 3, 3
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("default text")
		},
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			// co.(*widget.Label).SetText("asd")
		},
	)
	return container.NewVBox(sel, t)

}
