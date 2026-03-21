package cmd

import (
	"bytes"
	"io"
	"mime/multipart"
	"fmt"
	"os"
	"path/filepath"
	"github.com/spf13/cobra"
	"paprika/internal/client"
	"paprika/internal/output"
)

var (
	v2SyncUpsertGroceryItemsCmd_data string
)

var v2SyncUpsertGroceryItemsCmd = &cobra.Command{
	Use: "upsertGroceryItems",
	Short: "Create or update grocery items",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("PAPRIKA_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		var _mpBuf bytes.Buffer
		_mpWriter := multipart.NewWriter(&_mpBuf)
		var _mpErr error
		{
			var _mpFileBytes []byte
			_mpFileBytes, _mpErr = os.ReadFile(filepath.Clean(v2SyncUpsertGroceryItemsCmd_data))
			if _mpErr != nil {
				return fmt.Errorf("reading file: %w", _mpErr)
			}
			var _mpPart io.Writer
			_mpPart, _mpErr = _mpWriter.CreateFormFile("data", filepath.Base(v2SyncUpsertGroceryItemsCmd_data))
			if _mpErr != nil {
				return fmt.Errorf("creating form file: %w", _mpErr)
			}
			if _, _mpErr = _mpPart.Write(_mpFileBytes); _mpErr != nil {
				return fmt.Errorf("writing file content: %w", _mpErr)
			}
		}
		if _mpErr = _mpWriter.Close(); _mpErr != nil {
			return fmt.Errorf("closing multipart writer: %w", _mpErr)
		}
		resp, err := c.DoMultipart("POST", "/v2/sync/groceries/", pathParams, queryParams, &_mpBuf, _mpWriter.FormDataContentType())
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
	v2SyncCmd.AddCommand(v2SyncUpsertGroceryItemsCmd)
	v2SyncUpsertGroceryItemsCmd.Flags().StringVar(&v2SyncUpsertGroceryItemsCmd_data, "data", "", "")
	v2SyncUpsertGroceryItemsCmd.MarkFlagRequired("data")
}
