package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// editProjectCmd represents the editProject command
var editProjectCmd = &cobra.Command{
	Use:   "edit-project",
	Short: "Edit existing project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("editProject called")
	},
}

func init() {
	rootCmd.AddCommand(editProjectCmd)
}
