package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "show recent paths",
	Long:  "Display recently opened paths",
	Run: func(cmd *cobra.Command, args []string) {
		loadConfig()

		if len(config.RecentPaths) == 0 {
			fmt.Println("📝 No recent paths found")
			return
		}

		fmt.Println("📝 Recent paths:")
		for i, path := range config.RecentPaths {
			if isPathValid(path) {
				fmt.Printf("  %d. %s\n", i+1, formatPath(path))
			} else {
				// 移除无效路径
				config.RecentPaths = append(config.RecentPaths[:i], config.RecentPaths[i+1:]...)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
