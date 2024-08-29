package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cmmarslender/fj2/pkg/jinja2"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fj2",
	Short: "Jinja2 CLI templating in a single static binary",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		inputContent, err := os.ReadFile(args[0])
		if err != nil {
			log.Fatalf("Error reading input file: %s\n", err.Error())
		}

		result, err := jinja2.ExpandTemplate(string(inputContent))
		if err != nil {
			log.Fatalf("Error expanding template: %s\n", err.Error())
		}

		fmt.Print(result)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("output", "o", "", "Output file (optional)")
	cobra.CheckErr(viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output")))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match
}
