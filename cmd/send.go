package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/blackironj/ses-templator/ses"

	"github.com/aws/aws-sdk-go/aws"
	sessdk "github.com/aws/aws-sdk-go/service/ses"
	. "github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

const CharSet = "UTF-8"

type SendConfig struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Subject   string `json:"subject"`
	HtmlPath  string `json:"html_path"`
}

var (
	sendCmd = &cobra.Command{
		Use:   "send",
		Short: "send a email",
		Run: func(cmd *cobra.Command, args []string) {
			sendConfigPath, _ := filepath.Abs(path)

			jsonFile, readFileErr := ioutil.ReadFile(sendConfigPath)
			if readFileErr != nil {
				fmt.Println(Red("Fail to read a send config.\nPlease check the file path"))
				return
			}

			var sendInfo SendConfig
			unmarshalErr := json.Unmarshal(jsonFile, &sendInfo)
			if unmarshalErr != nil {
				fmt.Println(Red("Fail to read a send config.\nPlease check the fields in file"))
				return
			}

			emailInput, makeErr := makeEmailForm(&sendInfo)
			if makeErr != nil {
				fmt.Println(Red(makeErr))
			}

			sendErr := ses.SendEmailWithUnregisteredTemplate(emailInput)
			if sendErr != nil {
				fmt.Println(Red("Fail to send a email"))
				er(sendErr)
			} else {
				fmt.Println(Green("Success to send a email to "), sendInfo.Recipient)
			}
		},
	}
)

func makeEmailForm(sendInfo *SendConfig) (*sessdk.SendEmailInput, error) {
	htmlPath, _ := filepath.Abs(sendInfo.HtmlPath)
	htmlFile, readFileErr := ioutil.ReadFile(htmlPath)
	if readFileErr != nil {
		return nil, errors.New("Fail to read a html.\nPlease check the file path")
	}

	htmlStr := string(htmlFile)

	input := &sessdk.SendEmailInput{
		Destination: &sessdk.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(sendInfo.Recipient),
			},
		},
		Message: &sessdk.Message{
			Body: &sessdk.Body{
				Html: &sessdk.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(htmlStr),
				},
			},
			Subject: &sessdk.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(sendInfo.Subject),
			},
		},
		Source: aws.String(sendInfo.Sender),
	}

	return input, nil
}

func init() {
	rootCmd.AddCommand(sendCmd)

	sendCmd.Flags().StringVarP(&path, "path", "p", "", "send config path (required)")
	sendCmd.MarkFlagRequired("path")
}
