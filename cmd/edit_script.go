package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// editScriptCmd represents the editScript command
var editScriptCmd = &cobra.Command{
	Use:   "edit-script",
	Short: "Edit project scripts",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("editScript called")
	},
}

func init() {
	rootCmd.AddCommand(editScriptCmd)
}
