package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("📦 of - Open File Manager\n")
		fmt.Printf("Version: %s\n", version)
		fmt.Printf("Go Version: %s\n", runtime.Version())
		fmt.Printf("Platform: %s/%s\n", runtime.GOOS, runtime.GOARCH)
		fmt.Printf("Build Time: %s\n", getBuildTime())
	},
}

// getBuildTime 获取构建时间（这里返回编译时间）
func getBuildTime() string {
	// 在实际构建时，可以通过 ldflags 传入构建时间
	// 这里暂时返回一个占位符
	return "unknown"
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
