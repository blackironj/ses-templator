package screens

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/blackironj/ses-templator/ses"
)

var TABLE_HEADERS = [...]string{"Created Time", "Template Name"}

func rowsToColumns(headers []string, rows [][]string) [][]string {
	columns := make([][]string, len(headers))
	for _, row := range rows {
		for colK := range row {
			columns[colK] = append(columns[colK], row[colK])
		}
	}
	return columns
}

func makeTable() fyne.CanvasObject {
	templateList, err := ses.GetSESTemplateList(10)
	if err != nil {
		return fyne.NewContainer()
	}

	rows := make([][]string, 0)
	for _, info := range templateList {
		rows = append(rows, []string{info.CreatedTimestamp.String(), *info.Name})
	}

	columns := rowsToColumns(TABLE_HEADERS[:], rows)

	objects := make([]fyne.CanvasObject, len(columns))
	for k, col := range columns {
		box := widget.NewVBox(widget.NewLabelWithStyle(TABLE_HEADERS[k], fyne.TextAlignLeading, fyne.TextStyle{Bold: true}))
		for _, val := range col {
			box.Append(widget.NewLabel(val))
		}
		objects[k] = box
	}

	return fyne.NewContainerWithLayout(layout.NewGridLayout(len(columns)), objects...)
}

func MakeTemplateListScreen() fyne.CanvasObject {
	return makeTable()
}
