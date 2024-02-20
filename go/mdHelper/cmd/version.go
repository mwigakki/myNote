package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",                              // 此子命令的名称
	Short: "命令的版本",                                // 命令的简短描述
	Long:  `print the version number of this app`, // 命令的完整描述
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("the version number of this app: 0.1")

		fmt.Println(Dir)
	},
	/* Run 属性是一个函数，当不带任何参数直接运行appname时会调用这个函数。
	因此这个函数与一般是用来解释说明这个app是干嘛的
	*/
}

func init() { // 每个子命令的init也会在命令开始时被调用
	rootCmd.AddCommand(versionCmd) // 把version这个子命令加入到跟命令中去
}
