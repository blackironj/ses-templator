package cmd

import (
	"fmt"

	"github.com/blackironj/ses-templator/ses"

	. "github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var (
	deleteCmd = &cobra.Command{
		Use:   "del",
		Short: "delete a SES-email template",
		Run: func(cmd *cobra.Command, args []string) {
			metaDatas, listErr := ses.ListSEStemplates(10)
			if listErr != nil {
				er(listErr)
				return
			}

			findFl := false
			for _, data := range metaDatas {
				if templateName == *data.Name {
					findFl = true
					break
				}
			}

			if !findFl {
				fmt.Println(Red("Fail to find the name of template.\nPlease check the name"))
			}

			delErr := ses.DeleteSEStemplate(&templateName)
			if delErr != nil {
				fmt.Println(Red("Fail to delete"))
				er(delErr)
			} else {
				fmt.Println(Green("Success to delete a template : "), templateName)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringVarP(&templateName, "name", "n", "", "email template name (required)")
	deleteCmd.MarkFlagRequired("name")
}
