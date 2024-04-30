package interface_app

import (
	"myApp/pkg/forms"
	"time"

	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func (i *Interface_app) Resources() *widget.Form {
	currentTime := time.Now()

	currentDate := currentTime.Format("2006-01-02")

	resource := widget.NewEntry()
	cost := widget.NewEntry()
	date := widget.NewEntryWithData(binding.BindString(&currentDate))
	category := widget.NewSelect([]string{"Долгосрочные активы", "Расходные ресурсы"}, nil)

	category.PlaceHolder = "Выберете категорию"

	form := widget.NewForm(
		widget.NewFormItem("Ресурс", resource),
		widget.NewFormItem("Стоимость", cost),
		widget.NewFormItem("Категория", category),
		widget.NewFormItem("Дата", date),
	)

	form.SubmitText = "Отправить"
	form.OnSubmit = func() {
		err := forms.CheckForm(
			resource.Text,
			cost.Text,
			category.Selected,
			date.Text,
		)
		if err != nil {
			dialog.ShowError(err, i.Window)
		} else {
			// dialog.ShowInformation("")
		}
	}

	return form
}
