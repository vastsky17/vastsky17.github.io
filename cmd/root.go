package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "felix",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(bTime, gHash string) {
	gitHash = gHash
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var buildTime, gitHash string
var verbose, isShowVersion bool

func init() {
	rootCmd.Flags().BoolVarP(&isShowVersion, "version", "V", false, "show binary build information")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "verbose")
}
