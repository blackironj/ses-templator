package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/blackironj/ses-templator/ses"

	sessdk "github.com/aws/aws-sdk-go/service/ses"
	. "github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

type UploadTemplateForm struct {
	TemplateName string `json:"template_name"`
	Subject      string `json:"subject"`
	HtmlPath     string `json:"html_path"`
}

var uploadCmd = &cobra.Command{
	Use:   "up",
	Short: "upload the SES-email template",
	Run: func(cmd *cobra.Command, args []string) {
		uploadFormPath, _ := filepath.Abs(path)

		jsonFile, readFileErr := ioutil.ReadFile(uploadFormPath)
		if readFileErr != nil {
			fmt.Println(Red("Fail to read a upload template form.\nPlease check the file path"))
			return
		}

		var uploadForm UploadTemplateForm
		unmarshalErr := json.Unmarshal(jsonFile, &uploadForm)
		if unmarshalErr != nil {
			fmt.Println(Red("Fail to read a upload template form.\nPlease check the fields in file"))
			return
		}

		htmlPath, _ := filepath.Abs(uploadForm.HtmlPath)
		htmlFile, readFileErr := ioutil.ReadFile(htmlPath)
		if readFileErr != nil {
			fmt.Println(Red("Fail to read a html.\nPlease check the file path"))
			return
		}

		htmlStr := string(htmlFile)

		inputTemplate := &sessdk.Template{
			HtmlPart:     &htmlStr,
			TemplateName: &uploadForm.TemplateName,
			SubjectPart:  &uploadForm.Subject,
		}

		uploadErr := ses.UploadSEStemplate(inputTemplate)
		if uploadErr != nil {
			fmt.Println(Red("Fail to upload"))
			er(uploadErr)
		} else {
			fmt.Println(Green("Success to delete a template : "), uploadForm.TemplateName)
		}
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)

	uploadCmd.Flags().StringVarP(&path, "path", "p", "", "upload template fiile path (required)")
	uploadCmd.MarkFlagRequired("path")
}
