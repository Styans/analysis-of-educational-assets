package interface_app

import (
	"errors"
	"myApp/internal/models"
	"myApp/pkg/forms"
	"time"

	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func (i *Interface_app) Resources() *widget.Form {

	resource := widget.NewEntry()
	cost := widget.NewEntry()
	category := widget.NewSelect([]string{"Долгосрочные активы", "Расходные ресурсы"}, nil)

	category.PlaceHolder = "Выберете категорию"

	form := widget.NewForm(
		widget.NewFormItem("Ресурс", resource),
		widget.NewFormItem("Стоимость", cost),
		widget.NewFormItem("Категория", category),
	)

	form.SubmitText = "Отправить"
	form.OnSubmit = func() {
		err := forms.CheckForm(
			resource.Text,
			cost.Text,
			category.Selected,
		)

		if err != nil {
			dialog.ShowError(err, i.Window)
		} else {
			n, err := forms.IsInt(cost.Text)
			if err != nil {
				dialog.ShowError(errors.New("неверное значение в поле (цены)"), i.Window)
			} else {
				err := i.service.ResourcesService.CreateResources(
					&models.ResourcesDTO{
						Resource: resource.Text,
						Cost:     n,
						Category: category.Selected,
						Date:     time.Now(),
					})
				if err != nil {
					dialog.ShowError(err, i.Window)
				} else {
					dialog.ShowInformation("Done", "ресурс добавлен в список", i.Window)

				}
			}
		}
	}

	return form
}
