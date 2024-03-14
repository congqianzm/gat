package cmd

import (
	"gat/pkg"
	"log/slog"

	"github.com/spf13/cobra"
)

var path string

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Info("run called")
		cases, err := pkg.PathToCases(path)
		if err != nil {
			pkg.Logger.Error(err.Error())
			return
		}
		suite := pkg.Suite{Cases: cases}
		suite.Run()
		pkg.GenerateReportFromJson("template/summary.html", "template/data.json")
	},
}

func init() {
	runCmd.Flags().StringVarP(&path, "path", "p", "", "Path to run the command")
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
