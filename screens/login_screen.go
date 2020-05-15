package screens

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"

	"github.com/blackironj/ses-templator/container"
	"github.com/blackironj/ses-templator/ses"
)

var REGIONS = [...]string{
	"ap-south-1",
	"ap-south-2",
	"us-east-1",
	"us-west-2",
	"eu-west-1",
	"eu-west-2",
	"eu-central-1",
	"sa-east-1",
	"ca-central-1"}

func makeLoginForm(win fyne.Window) fyne.Widget {
	accessKey := widget.NewEntry()
	secretKey := widget.NewEntry()

	selectEntry := widget.NewEntry()

	region := "select a your region"

	selectEntry.SetPlaceHolder(region)
	selectEntry.SetReadOnly(true)

	selector := widget.NewSelect(
		REGIONS[:],
		func(r string) {
			region = r
		})

	signInbtn := widget.NewButton("LOGIN", func() {
		var err error
		ses.EamilServiceSess, err = ses.NewSession(accessKey.Text, secretKey.Text, region)
		if err != nil {
			dialog.ShowInformation("Fail to login", "Please check your keys", win)
		}

		_, err = ses.GetSESTemplateList(1)
		if err != nil {
			dialog.ShowInformation("Fail to login", "Please check your keys", win)
		} else {
			container.TabItemMap["list"] = widget.NewTabItemWithIcon("list", theme.FolderIcon(), MakeTemplateListScreen())
			container.TabItemMap["upload"] = widget.NewTabItemWithIcon("upload", theme.ContentAddIcon(), MakeUploadTemplateScreen(win))

			container.RemoveTabItem("login")
			container.AppendTabItem("upload")
			container.AppendTabItem("list")
			container.TabContainer.SelectTab(container.TabItemMap["list"])

		}
	})

	form := &widget.Form{}
	form.Append("Access Key", accessKey)
	form.Append("Secret Key", secretKey)
	form.Append("Region", selector)
	form.Append("", signInbtn)

	return form
}

func MakeLoginScreen(win fyne.Window) fyne.CanvasObject {
	return makeLoginForm(win)
}
