package cmd

import (
	"github.com/atpuxiner/gin-layout/cmd/internal/runserver"
	"github.com/spf13/cobra"
)

// runserverCmd represents the runserver command
var runserverCmd = &cobra.Command{
	Use:   "runserver",
	Short: "runserver",
	Long:  `runserver`,
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetString("port")
		defer runserver.Clean()
		runserver.Start(port)
	},
}

func init() {
	rootCmd.AddCommand(runserverCmd)
	// add flags
	runserverCmd.Flags().StringP("port", "p", "", "服务端口")
}
