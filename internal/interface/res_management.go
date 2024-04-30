package interface_app

import (
	"time"

	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func (i *Interface_app) Resources() *widget.Form {
	currentTime := time.Now()

	currentDate := currentTime.Format("2006-01-02")

	return widget.NewForm(
		widget.NewFormItem("Ресурс", widget.NewEntry()),
		widget.NewFormItem("Стоимость", widget.NewEntry()),
		// widget.NewFormItem("Стоимость обучения", widget.NewEntry()),
		widget.NewFormItem("Дата", widget.NewEntryWithData(binding.BindString(&currentDate))),
		// widget.NewFormItem("Дата окончания", widget.NewEntryWithData(binding.BindString(&currentDat))),
	)
}
