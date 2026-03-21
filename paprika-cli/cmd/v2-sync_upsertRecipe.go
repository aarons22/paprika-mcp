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
	v2SyncUpsertRecipeCmd_data string
)

var v2SyncUpsertRecipeCmd = &cobra.Command{
	Use: "upsertRecipe <uid>",
	Short: "Create or update a recipe",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("PAPRIKA_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		pathParams["uid"] = args[0]
		queryParams := map[string]string{}
		var _mpBuf bytes.Buffer
		_mpWriter := multipart.NewWriter(&_mpBuf)
		var _mpErr error
		{
			var _mpFileBytes []byte
			_mpFileBytes, _mpErr = os.ReadFile(filepath.Clean(v2SyncUpsertRecipeCmd_data))
			if _mpErr != nil {
				return fmt.Errorf("reading file: %w", _mpErr)
			}
			var _mpPart io.Writer
			_mpPart, _mpErr = _mpWriter.CreateFormFile("data", filepath.Base(v2SyncUpsertRecipeCmd_data))
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
		resp, err := c.DoMultipart("POST", "/v2/sync/recipe/{uid}/", pathParams, queryParams, &_mpBuf, _mpWriter.FormDataContentType())
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
	v2SyncCmd.AddCommand(v2SyncUpsertRecipeCmd)
	v2SyncUpsertRecipeCmd.Flags().StringVar(&v2SyncUpsertRecipeCmd_data, "data", "", "")
	v2SyncUpsertRecipeCmd.MarkFlagRequired("data")
}
