package cmd

import (
	"os"

	"github.com/blackironj/ses-templator/ses"
	"github.com/jedib0t/go-pretty/table"

	"github.com/spf13/cobra"
)

var (
	num int

	listCmd = &cobra.Command{
		Use:   "list",
		Short: "get list of SES-email templates",
		Run: func(cmd *cobra.Command, args []string) {
			metaDatas, err := ses.ListSEStemplates(num)
			if err != nil {
				er(err)
			}

			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)
			t.AppendHeader(table.Row{"created time", "Template name"})

			for _, data := range metaDatas {
				t.AppendRows([]table.Row{
					{*data.CreatedTimestamp, *data.Name},
				})
			}
			t.SetStyle(table.StyleColoredBright)
			t.Render()
		},
	}
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().IntVar(&num, "num", 10, "max amount of template")
}
