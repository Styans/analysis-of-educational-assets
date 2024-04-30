package interface_app

import (
	"time"

	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func (i *Interface_app) StudManager() *widget.Form {
	currentTime := time.Now()

	Start := time.Date(
		currentTime.Year(),
		time.September,
		1, 0, 0, 0, 0,
		currentTime.Location(),
	)

	currentDate := Start.Format("2006-01-02")

	currentDat := Start.AddDate(1, 0, 0).Format("2006-01-02")

	return widget.NewForm(
		widget.NewFormItem("Ф.И.О.", widget.NewEntry()),
		widget.NewFormItem("Группа", widget.NewEntry()),
		widget.NewFormItem("Стоимость обучения", widget.NewEntry()),
		widget.NewFormItem("Дата поступления", widget.NewEntryWithData(binding.BindString(&currentDate))),
		widget.NewFormItem("Дата окончания", widget.NewEntryWithData(binding.BindString(&currentDat))),
	)
}
