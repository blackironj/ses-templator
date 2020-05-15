package screens

import (
	"errors"
	"io/ioutil"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"

	"github.com/blackironj/ses-templator/ses"
)

func readFile(f fyne.FileReadCloser) (string, error) {
	if f == nil {
		return "", nil
	}

	ext := f.URI()[len(f.URI())-5:]
	if ext != ".html" {
		return "", errors.New("check a file format")
	}
	return loadText(f)
}

func loadText(f fyne.FileReadCloser) (string, error) {
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	if data == nil {
		return "", nil
	}
	return string(data), nil
}

func makeUploadTemplateForm(win fyne.Window) fyne.Widget {
	templateName := widget.NewEntry()
	subject := widget.NewEntry()

	var htmlBody string

	uploadHTML := widget.NewButton("Upload HTML", func() {
		dialog.ShowFileOpen(func(reader fyne.FileReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			str, err := readFile(reader)
			if err != nil {
				dialog.ShowError(err, win)
			}
			htmlBody = str

		}, win)
	})

	form := &widget.Form{
		OnSubmit: func() {
			err := ses.CreateSESTemplate(&templateName.Text, &subject.Text, &htmlBody)
			if err != nil {
				dialog.ShowInformation("Fail to upload", "Please try again", win)
			} else {
				dialog.ShowInformation("Success to upload", "", win)
			}
		},
	}

	form.Append("Template Name", templateName)
	form.Append("Subject", subject)
	form.Append("HTML body", uploadHTML)

	return form
}

func MakeUploadTemplateScreen(win fyne.Window) fyne.CanvasObject {
	return makeUploadTemplateForm(win)
}
