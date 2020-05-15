package container

import (
	"fyne.io/fyne/widget"
)

var TabContainer *widget.TabContainer
var TabItemMap map[string]*widget.TabItem

func InitTabItemMap() {
	TabItemMap = make(map[string]*widget.TabItem)
}

func RemoveTabItem(tabName string) {
	TabContainer.Remove(TabItemMap[tabName])
}

func AppendTabItem(tabName string) {
	TabContainer.Append(TabItemMap[tabName])
}
