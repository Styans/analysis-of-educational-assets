package interface_app

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (i *Interface_app) ResourcesTable() *fyne.Container {

	var r [][]string
	var err error
	category := widget.NewSelect(
		[]string{
			"Долгосрочные активы",
			"Расходные ресурсы",
		},
		func(s string) {
			switch s {
			case "Долгосрочные активы":
				r, err = i.service.ResourcesService.GetResourcesByType(s)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("great", r)
				}
			case "Расходные ресурсы":
				r, err = i.service.ResourcesService.GetResourcesByType(s)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("great", r)
				}
			}
		},
	)
	category.PlaceHolder = "Выберете категорию"

	t := widget.NewTable(

		func() (int, int) {
			return len(r), 3
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("default text")
		},
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(r[tci.Row][tci.Col])
		},
	)
	scrollContainer := container.NewVScroll(t)
	scrollContainer.SetMinSize(fyne.NewSize(200, 300))
	return container.NewVBox(category, scrollContainer)

}
