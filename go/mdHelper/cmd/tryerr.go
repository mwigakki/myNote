package cmd

import (
	"io"

	"github.com/spf13/cobra"
)

var tryerr = &cobra.Command{
	Use:   "tryerr",
	Short: "测试错误处理 Short",
	Long:  "测试错误处理 Long",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := io.EOF; err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(tryerr)
}
