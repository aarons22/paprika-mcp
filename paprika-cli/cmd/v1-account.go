package cmd

import "github.com/spf13/cobra"

var v1AccountCmd = &cobra.Command{
	Use: "v1-account",
	Short: "v1-account",
}

func init() {
	rootCmd.AddCommand(v1AccountCmd)
}
