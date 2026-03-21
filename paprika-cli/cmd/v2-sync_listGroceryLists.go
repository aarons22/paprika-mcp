package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"paprika/internal/client"
	"paprika/internal/output"
)

var v2SyncListGroceryListsCmd = &cobra.Command{
	Use: "listGroceryLists",
	Short: "List all grocery lists",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("PAPRIKA_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		resp, err := c.Do("GET", "/v2/sync/grocerylists/", pathParams, queryParams, nil)
		if err != nil {
			return err
		}
		jsonMode, _ := cmd.Root().PersistentFlags().GetBool("json")
		noColor, _ := cmd.Root().PersistentFlags().GetBool("no-color")
		if jsonMode {
			fmt.Printf("%s\n", string(resp))
		} else {
			if err := output.PrintTable(resp, noColor); err != nil {
				fmt.Println(string(resp))
			}
		}
		return nil
	},
}

func init() {
	v2SyncCmd.AddCommand(v2SyncListGroceryListsCmd)
}
