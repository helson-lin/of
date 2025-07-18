package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "manage configuration",
	Long:  "Manage of tool configuration",
}

var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "show current configuration",
	Run: func(cmd *cobra.Command, args []string) {
		loadConfig()

		home, _ := os.UserHomeDir()
		configFile := filepath.Join(home, ".of", "config.yaml")

		fmt.Printf("📁 Config file: %s\n", configFile)
		fmt.Printf("🔧 Default manager: %s\n", config.DefaultManager)
		fmt.Printf("📊 Recent paths count: %d\n", len(config.RecentPaths))
		fmt.Printf("📈 Max recent paths: %d\n", config.MaxRecent)

		if len(config.CustomManagers) > 0 {
			fmt.Println("🔧 Custom managers:")
			for name, cmd := range config.CustomManagers {
				fmt.Printf("  %s: %s\n", name, cmd)
			}
		}

		if len(config.FileTypeApps) > 0 {
			fmt.Println("📄 File type applications:")
			for ext, app := range config.FileTypeApps {
				fmt.Printf("  .%s: %s\n", ext, app)
			}
		}
	},
}

var configAddManagerCmd = &cobra.Command{
	Use:   "add-manager [name] [command]",
	Short: "add custom file manager",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		loadConfig()

		name := args[0]
		command := args[1]

		if config.CustomManagers == nil {
			config.CustomManagers = make(map[string]string)
		}

		config.CustomManagers[name] = command
		viper.Set("custom_managers", config.CustomManagers)

		if err := viper.WriteConfig(); err != nil {
			fmt.Printf("❌ Error saving config: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("✅ Added custom manager: %s -> %s\n", name, command)
	},
}

var configSetDefaultCmd = &cobra.Command{
	Use:   "set-default [manager]",
	Short: "set default file manager",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		loadConfig()

		managerName := args[0]
		config.DefaultManager = managerName
		viper.Set("default_manager", managerName)

		if err := viper.WriteConfig(); err != nil {
			fmt.Printf("❌ Error saving config: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("✅ Set default manager: %s\n", managerName)
	},
}

var configClearRecentCmd = &cobra.Command{
	Use:   "clear-recent",
	Short: "clear recent paths",
	Run: func(cmd *cobra.Command, args []string) {
		loadConfig()

		config.RecentPaths = []string{}
		viper.Set("recent_paths", config.RecentPaths)

		if err := viper.WriteConfig(); err != nil {
			fmt.Printf("❌ Error saving config: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("✅ Cleared recent paths")
	},
}

var configAddFileTypeCmd = &cobra.Command{
	Use:   "add-filetype [extension] [app]",
	Short: "add file type application mapping",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		loadConfig()

		ext := strings.ToLower(strings.TrimPrefix(args[0], "."))
		app := args[1]

		// 验证应用程序是否存在
		exists, message := validateApp(app)
		if !exists {
			fmt.Printf("❌ %s\n", message)
			os.Exit(1)
		}

		if config.FileTypeApps == nil {
			config.FileTypeApps = make(map[string]string)
		}

		config.FileTypeApps[ext] = app
		viper.Set("file_type_apps", config.FileTypeApps)

		if err := viper.WriteConfig(); err != nil {
			fmt.Printf("❌ Error saving config: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("✅ Added file type mapping: .%s -> %s\n", ext, app)
	},
}

var configRemoveFileTypeCmd = &cobra.Command{
	Use:   "remove-filetype [extension]",
	Short: "remove file type application mapping",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		loadConfig()

		ext := strings.ToLower(strings.TrimPrefix(args[0], "."))

		if config.FileTypeApps == nil {
			fmt.Printf("❌ No file type mappings found\n")
			os.Exit(1)
		}

		if _, exists := config.FileTypeApps[ext]; !exists {
			fmt.Printf("❌ File type .%s not found in mappings\n", ext)
			os.Exit(1)
		}

		delete(config.FileTypeApps, ext)
		viper.Set("file_type_apps", config.FileTypeApps)

		if err := viper.WriteConfig(); err != nil {
			fmt.Printf("❌ Error saving config: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("✅ Removed file type mapping: .%s\n", ext)
	},
}

var configListFileTypesCmd = &cobra.Command{
	Use:   "list-filetypes",
	Short: "list all file type mappings",
	Run: func(cmd *cobra.Command, args []string) {
		loadConfig()

		if len(config.FileTypeApps) == 0 {
			fmt.Println("📄 No file type mappings found")
			return
		}

		fmt.Println("📄 File type mappings:")
		for ext, app := range config.FileTypeApps {
			fmt.Printf("  .%s -> %s\n", ext, app)
		}
	},
}

var configAddFileGroupCmd = &cobra.Command{
	Use:   "add-filegroup [group] [app]",
	Short: "add file type group application mapping",
	Long: `Add file type group application mapping.

Available groups:
  audio     - mp3, wav, flac, aac, ogg, m4a, wma
  video     - mp4, avi, mkv, mov, wmv, flv, webm, m4v, 3gp
  image     - jpg, jpeg, png, gif, bmp, svg, tiff, webp
  document  - pdf, doc, docx, txt, md, rtf
  code      - py, js, ts, go, java, cpp, c, h, html, css, json, xml, yaml, yml
  archive   - zip, rar, 7z, tar, gz, bz2
  spreadsheet - xls, xlsx, csv
  presentation - ppt, pptx`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		loadConfig()

		group := strings.ToLower(args[0])
		app := args[1]

		// 验证应用程序是否存在
		exists, message := validateApp(app)
		if !exists {
			fmt.Printf("❌ %s\n", message)
			os.Exit(1)
		}

		// 定义文件类型组
		fileGroups := map[string][]string{
			"audio":        {"mp3", "wav", "flac", "aac", "ogg", "m4a", "wma"},
			"video":        {"mp4", "avi", "mkv", "mov", "wmv", "flv", "webm", "m4v", "3gp"},
			"image":        {"jpg", "jpeg", "png", "gif", "bmp", "svg", "tiff", "webp"},
			"document":     {"pdf", "doc", "docx", "txt", "md", "rtf"},
			"code":         {"py", "js", "ts", "go", "java", "cpp", "c", "h", "html", "css", "json", "xml", "yaml", "yml"},
			"archive":      {"zip", "rar", "7z", "tar", "gz", "bz2"},
			"spreadsheet":  {"xls", "xlsx", "csv"},
			"presentation": {"ppt", "pptx"},
		}

		extensions, exists := fileGroups[group]
		if !exists {
			fmt.Printf("❌ Unknown file group: %s\n", group)
			fmt.Println("Available groups:")
			for g := range fileGroups {
				fmt.Printf("  %s\n", g)
			}
			os.Exit(1)
		}

		if config.FileTypeApps == nil {
			config.FileTypeApps = make(map[string]string)
		}

		// 添加所有扩展名
		count := 0
		for _, ext := range extensions {
			config.FileTypeApps[ext] = app
			count++
		}

		viper.Set("file_type_apps", config.FileTypeApps)

		if err := viper.WriteConfig(); err != nil {
			fmt.Printf("❌ Error saving config: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("✅ Added file group mapping: %s (%d file types) -> %s\n", group, count, app)
		fmt.Printf("📄 Added extensions: %s\n", strings.Join(extensions, ", "))
	},
}

func init() {
	configCmd.AddCommand(configShowCmd)
	configCmd.AddCommand(configAddManagerCmd)
	configCmd.AddCommand(configSetDefaultCmd)
	configCmd.AddCommand(configClearRecentCmd)
	configCmd.AddCommand(configAddFileTypeCmd)
	configCmd.AddCommand(configRemoveFileTypeCmd)
	configCmd.AddCommand(configListFileTypesCmd)
	configCmd.AddCommand(configAddFileGroupCmd)
	rootCmd.AddCommand(configCmd)
}
