package cmd


import (
	"fmt"
	"github.com/georgevella/helmtmpl/config"
	"github.com/spf13/cobra"
	"strings"
)

var environment string

func init() {
	rootCmd.AddCommand(renderCmd)

	rootCmd.Flags().StringVarP(&environment, "environment", "e", "", "Name of environment to render")
}

var renderCmd = &cobra.Command{
	Use:   "render [CHART] [NAME]",
	Short: "Renders a Helm Chart with the selected environment variables and templates",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadConfiguration()


		for _, env := range cfg.Environments {
			if len(environment) == 0 {
				// determine environment from selectors specified in config
				for _, selector := range env.Selector {
					switch strings.ToLower(selector.Type) {
					case "branch":

					}
				}
			} else {
				if env.Name == environment {
					renderEnvironment(&env, cfg)
				}
			}
		}
	},
}

func renderEnvironment(env *config.Environment, cfg *config.Config) {
	fmt.Print(env)
}