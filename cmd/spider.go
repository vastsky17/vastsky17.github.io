package cmd

import (
	"github.com/spf13/cobra"
	"gojekyll/spiderhn"
	"log"
)

// spiderCmd represents the spiderHN command
var spiderCmd = &cobra.Command{
	Use:   "spider",
	Short: "爬取翻译hacknews",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		spiderhn.LoadSQLiteDb(false)
		if err := spiderhn.SpiderHackNews(); err != nil {
			log.Println(err)
		}
		if err := spiderhn.SpiderHackShows(); err != nil {
			log.Println(err)
		}
		if err := spiderhn.OutputMarkdown("."); err != nil {
			log.Println(err)
		}
		log.Println("work is done")
	},
}

func init() {
	rootCmd.AddCommand(spiderCmd)

}
