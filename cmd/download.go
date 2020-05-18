package cmd

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/blackironj/ses-templator/ses"

	. "github.com/logrusorgru/aurora"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use:   "down",
	Short: "download the SES-email template",
	Run: func(cmd *cobra.Command, args []string) {
		var downloadPath string

		if path == "" {
			home, err := homedir.Dir()
			if err != nil {
				er(err)
				return
			}
			downloadPath = home
		} else {
			downloadPath = path
		}

		templateOutput, getErr := ses.GetSEStemplate(&templateName)
		if getErr != nil {
			fmt.Println(Red("Fail to download."))
			er(getErr)
			return
		}

		if path == "" {
			downloadPath = filepath.Join(downloadPath, *templateOutput.Template.TemplateName+".html")
		}
		fmt.Println(Green("Success to download template"))

		fmt.Println("Template name : ", *templateOutput.Template.TemplateName)
		fmt.Println("Subject : ", *templateOutput.Template.SubjectPart)

		writeErr := ioutil.WriteFile(downloadPath, []byte(*templateOutput.Template.HtmlPart), 0644)
		if writeErr != nil {
			fmt.Println(Red("Fail to save html"))
			er(getErr)
			return
		}

		fmt.Println(Green("Success to save html : "), downloadPath)
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	downloadCmd.PersistentFlags().StringVarP(&path, "path", "p", "", "download path (default is $HOME/`template name`.html)")

	downloadCmd.Flags().StringVarP(&templateName, "name", "n", "", "email template name (required)")
	downloadCmd.MarkFlagRequired("name")
}
