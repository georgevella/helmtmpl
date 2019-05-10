package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "helm-deploy",
	Short: "Helm deploy is an template based, helm chart deployment tool.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			err := cmd.Help()
			if err != nil {
				panic(err)
			}

			os.Exit(0)
		}
	},
}

func init()  {
	cobra.EnableCommandSorting = true

	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&configFilePath, "config", "", "Path to the configuration file")
}

var configFilePath string

func initConfig() {
	viper.SetConfigType("yaml")

	// Don't forget to read config either from cfgFile or from home directory!
	if configFilePath != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configFilePath)
	} else {
		dir, e := os.Getwd()
		if e != nil {
			log.Fatal(e)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(dir)
		viper.SetConfigName("helmdeploy.yaml")
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}