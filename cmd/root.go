package cmd

import (
	"fmt"
	"os"

	"github.com/blackironj/ses-templator/ses"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	keyFile      string
	templateName string
	path         string

	rootCmd = &cobra.Command{
		Use:   "ses-templator",
		Short: "A helper for managing aws-ses email templates",
		Long: `ses-templator is a CLI tool. 
It helps you update / edit / get templates easily`,
	}
)

func init() {
	cobra.OnInitialize(initAccess)

	rootCmd.PersistentFlags().StringVar(&keyFile, "key", "", "access-key file (default is $HOME/.aws-access-key.json)")
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		er(err)
	}
}

func initAccess() {
	if keyFile != "" {
		viper.SetConfigFile(keyFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			er(err)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".aws-access-key")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using access-key file:", viper.ConfigFileUsed())
	}

	ak := viper.GetString("access_key")
	sk := viper.GetString("secret_key")
	region := viper.GetString("region")

	var sessErr error
	ses.EamilServiceSess, sessErr = ses.NewSession(ak, sk, region)
	if sessErr != nil {
		er(sessErr)
	}
}
