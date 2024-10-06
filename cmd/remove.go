package cmd

import (
	"fmt"

	"github.com/lewissteele/dbat/internal/db"
	"github.com/lewissteele/dbat/internal/model"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove database connection",
	Run: func(cmd *cobra.Command, args []string) {
		host := args[0]

		db.LocalDB.Where("host = ?", host).Delete(&model.Database{})

		fmt.Println("removed database connection", host)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
