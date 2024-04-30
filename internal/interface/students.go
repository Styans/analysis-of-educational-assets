package interface_app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func (i *Interface_app) Students() *widget.Table {
	return widget.NewTable(

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
}
