package cmd

import (
	"fmt"

	"github.com/lewissteele/dbat/internal/db"
	"github.com/lewissteele/dbat/internal/list"
	"github.com/lewissteele/dbat/internal/model"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove database connection",
	Run: func(cmd *cobra.Command, args []string) {
		var database string

		if len(args) > 0 {
			database = args[0]
		}

		if len(database) == 0 {
			database = list.RenderConnectionSelection()
		}

		db.LocalDB.Where("name = ?", database).Delete(&model.Database{})
		fmt.Println("removed database connection")
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
