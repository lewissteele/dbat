package cmd

import (
	"os"

	"github.com/lewissteele/dbat/internal/db"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dbat",
	Short: "DBat is a client for all SQL databases",
	Run: func(cmd *cobra.Command, args []string) {
		if len(db.UserDBNames()) > 0 {
			connectCmd.Run(cmd, args)
			return
		}

		cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
