package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

// techMojoCmd represents the taskrm command
var jekyllServeCmd = &cobra.Command{
	Use:   "run",
	Short: "run bundle execu jekyll $1",
	Long:  `只能在我家里面的电脑使用`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		jekyllDir := `D:\code\libragen.cn`

		thisCmd := exec.Command("bundle", "exec", "jekyll", args[0])
		thisCmd.Dir = jekyllDir
		o, err := thisCmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(o))

	},
}

func init() {
	rootCmd.AddCommand(jekyllServeCmd)
}
