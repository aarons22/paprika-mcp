package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"paprika/internal/client"
	"paprika/internal/output"
)

var (
	v1AccountLoginCmdBody string
	v1AccountLoginCmdBodyFile string
)

var v1AccountLoginCmd = &cobra.Command{
	Use: "login",
	Short: "Authenticate and obtain a Bearer token",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("PAPRIKA_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if v1AccountLoginCmdBodyFile != "" {
			fileData, err := os.ReadFile(v1AccountLoginCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			v1AccountLoginCmdBody = string(fileData)
		}
		if v1AccountLoginCmdBody != "" {
			if !json.Valid([]byte(v1AccountLoginCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(v1AccountLoginCmdBody), &bodyObj)
			resp, err := c.Do("POST", "/v1/account/login/", pathParams, queryParams, bodyObj)
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
		}
		resp, err := c.Do("POST", "/v1/account/login/", pathParams, queryParams, nil)
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
	v1AccountCmd.AddCommand(v1AccountLoginCmd)
	v1AccountLoginCmd.Flags().StringVar(&v1AccountLoginCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	v1AccountLoginCmd.Flags().StringVar(&v1AccountLoginCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
}
