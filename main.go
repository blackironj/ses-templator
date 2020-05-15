package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"

	"github.com/blackironj/ses-templator/container"
	"github.com/blackironj/ses-templator/screens"
)

const preferenceCurrentTab = "currentTab"

func init() {
	container.InitTabItemMap()
}

func main() {
	a := app.NewWithID("io.ses.templator")
	a.SetIcon(theme.FyneLogo())

	w := a.NewWindow("Ses Eamil Templator")

	w.SetMaster()

	container.TabItemMap["login"] = widget.NewTabItemWithIcon("login", theme.ConfirmIcon(), screens.MakeLoginScreen(w))

	container.TabContainer = widget.NewTabContainer(container.TabItemMap["login"])
	container.TabContainer.SetTabLocation(widget.TabLocationLeading)
	container.TabContainer.SelectTabIndex(a.Preferences().Int(preferenceCurrentTab))

	w.SetContent(container.TabContainer)

	w.Resize(fyne.Size{Height: 480, Width: 640})

	w.ShowAndRun()
	a.Preferences().SetInt(preferenceCurrentTab, container.TabContainer.CurrentTabIndex())
}
