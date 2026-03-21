package cmd

import "github.com/spf13/cobra"

var v2SyncCmd = &cobra.Command{
	Use: "v2-sync",
	Short: "v2-sync",
}

func init() {
	rootCmd.AddCommand(v2SyncCmd)
}
