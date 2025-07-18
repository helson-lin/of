package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	version string = "0.0.1"
	path    string
	debug   bool
	manager string

	// 配置结构体
	config struct {
		DefaultManager string            `mapstructure:"default_manager"`
		CustomManagers map[string]string `mapstructure:"custom_managers"`
		RecentPaths    []string          `mapstructure:"recent_paths"`
		MaxRecent      int               `mapstructure:"max_recent"`
		FileTypeApps   map[string]string `mapstructure:"file_type_apps"`
	}

	rootCmd = &cobra.Command{
		Use:   "of [path]",
		Short: "open your file or directory in file manager",
		Long: `Open files or directories in file manager from terminal.

Examples:
  of                    # 打开当前目录
  of /path/to/folder    # 打开指定路径
  of -p /path/to/file   # 使用标志指定路径
  of -m finder          # 指定文件管理器
  of --debug            # 启用调试模式`,
		Args: cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// 加载配置
			loadConfig()

			// 调试模式
			if debug {
				fmt.Printf("🔍 Debug mode enabled\n")
				fmt.Printf("🔍 OS: %s\n", runtime.GOOS)
				fmt.Printf("🔍 Manager: %s\n", manager)
			}

			// 如果没有提供子命令，显示帮助信息
			if len(args) == 0 && path == "" {
				if err := cmd.Help(); err != nil {
					fmt.Printf("❌ Error: cannot display help: %v\n", err)
					os.Exit(1)
				}
				return
			}

			// 获取要打开的路径
			targetPath := path
			if len(args) > 0 {
				targetPath = args[0]
			}

			// 如果路径为空，使用当前目录
			if targetPath == "" {
				currentDir, err := os.Getwd()
				if err != nil {
					fmt.Printf("❌ Error: cannot get current directory: %v\n", err)
					os.Exit(1)
				}
				targetPath = currentDir
			}

			// 检查路径是否存在
			if !isPathValid(targetPath) {
				fmt.Printf("❌ Error: path does not exist: %s\n", targetPath)
				os.Exit(1)
			}

			// 获取绝对路径
			absPath, err := filepath.Abs(targetPath)
			if err != nil {
				fmt.Printf("❌ Error: cannot get absolute path: %v\n", err)
				os.Exit(1)
			}

			// 如果没有指定管理器，使用默认管理器
			if manager == "" && config.DefaultManager != "" {
				manager = config.DefaultManager
				if debug {
					fmt.Printf("🔍 Using default manager: %s\n", manager)
				}
			}

			// 检查是否为文件，如果是文件则根据文件类型选择应用程序
			if isFile(absPath) {
				appForFile := getAppForFileType(absPath)
				if appForFile != "" {
					if debug {
						fmt.Printf("🔍 File type detected, using app: %s\n", appForFile)
					}
					err = openFileWithApp(absPath, appForFile)
				} else {
					// 没有配置的文件类型使用默认文件管理器
					err = openInFileManager(absPath)
				}
			} else {
				// 文件夹使用默认文件管理器
				err = openInFileManager(absPath)
			}

			if err != nil {
				fmt.Printf("❌ Error: cannot open path: %v\n", err)
				os.Exit(1)
			}

			// 添加到最近使用列表
			addToRecentPaths(absPath)

			// 确定使用的应用程序或文件管理器名称
			var usedApp string
			if isFile(absPath) {
				appForFile := getAppForFileType(absPath)
				if appForFile != "" {
					usedApp = appForFile
				} else {
					usedApp = getFileManagerName()
				}
			} else {
				usedApp = getFileManagerName()
			}

			if manager != "" {
				usedApp = manager
			}

			fmt.Printf("🚀 Opened in %s: %s\n", usedApp, formatPath(absPath))
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().StringVarP(&path, "path", "p", "", "path to file or directory to open")
	rootCmd.Flags().StringVarP(&manager, "manager", "m", "", "specify file manager to use")
	rootCmd.Flags().BoolVarP(&debug, "debug", "d", false, "enable debug mode")
}

// openInFileManager 使用系统默认的文件管理器打开文件或文件夹
func openInFileManager(path string) error {
	// 如果指定了自定义管理器
	if manager != "" {
		if customCmd, exists := getCustomManager(manager); exists {
			if debug {
				fmt.Printf("🔍 Using custom manager: %s -> %s\n", manager, customCmd)
			}
			cmd := exec.Command(customCmd, path)
			return cmd.Run()
		}

		// 尝试直接使用指定的管理器名称
		if debug {
			fmt.Printf("🔍 Trying direct manager: %s\n", manager)
		}
		cmd := exec.Command(manager, path)
		if err := cmd.Run(); err == nil {
			return nil
		}
	}

	switch runtime.GOOS {
	case "darwin":
		// macOS - 使用 Finder
		cmd := exec.Command("open", path)
		return cmd.Run()
	case "windows":
		// Windows - 使用 Explorer
		cmd := exec.Command("explorer", path)
		_ = cmd.Run() // Windows explorer 即使成功打开文件夹也可能返回非零状态码
		// 所以我们忽略错误，因为文件夹实际上已经被打开了
		return nil
	case "linux":
		// Linux - 尝试使用 xdg-open
		cmd := exec.Command("xdg-open", path)
		return cmd.Run()
	default:
		return fmt.Errorf("⚠️ unsupported operating system: %s", runtime.GOOS)
	}
}

// getFileManagerName 获取当前平台的文件管理器名称
func getFileManagerName() string {
	switch runtime.GOOS {
	case "darwin":
		return "Finder"
	case "windows":
		return "Explorer"
	case "linux":
		return "File Manager"
	default:
		return "File Manager"
	}
}

// loadConfig 加载配置文件
func loadConfig() {
	// 设置配置文件路径
	var home string
	var err error

	// 检查是否使用 sudo 运行
	if os.Geteuid() == 0 && os.Getenv("SUDO_USER") != "" {
		// 使用 sudo 运行时，尝试获取原始用户的主目录
		sudoUser := os.Getenv("SUDO_USER")
		if sudoUser != "" {
			// 在 macOS 上，用户主目录通常在 /Users/username
			home = "/Users/" + sudoUser
			if _, err := os.Stat(home); os.IsNotExist(err) {
				// 如果路径不存在，回退到当前用户的主目录
				home, err = os.UserHomeDir()
				if err != nil {
					if debug {
						fmt.Printf("⚠️ Warning: cannot get home directory: %v\n", err)
					}
					return
				}
			}
		} else {
			home, err = os.UserHomeDir()
			if err != nil {
				if debug {
					fmt.Printf("⚠️ Warning: cannot get home directory: %v\n", err)
				}
				return
			}
		}
	} else {
		// 正常运行时，使用当前用户的主目录
		home, err = os.UserHomeDir()
		if err != nil {
			if debug {
				fmt.Printf("⚠️ Warning: cannot get home directory: %v\n", err)
			}
			return
		}
	}

	configDir := filepath.Join(home, ".of")
	configFile := filepath.Join(configDir, "config.yaml")

	if debug {
		fmt.Printf("🔍 Config directory: %s\n", configDir)
		fmt.Printf("🔍 Config file: %s\n", configFile)
	}

	// 创建配置目录
	if err := os.MkdirAll(configDir, 0755); err != nil {
		if debug {
			fmt.Printf("⚠️ Warning: cannot create config directory: %v\n", err)
		}
		return
	}

	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")

	// 设置默认值
	viper.SetDefault("default_manager", "")
	viper.SetDefault("custom_managers", map[string]string{})
	viper.SetDefault("recent_paths", []string{})
	viper.SetDefault("max_recent", 10)
	viper.SetDefault("file_type_apps", map[string]string{})

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		if debug {
			fmt.Printf("🔍 No config file found, using defaults\n")
		}
		// 创建默认配置文件
		if err := viper.SafeWriteConfig(); err != nil {
			if debug {
				fmt.Printf("⚠️ Warning: cannot write config file: %v\n", err)
			}
		}
	}

	// 解析配置到结构体
	if err := viper.Unmarshal(&config); err != nil {
		if debug {
			fmt.Printf("⚠️ Warning: cannot parse config: %v\n", err)
		}
	}
}

// addToRecentPaths 添加路径到最近使用列表
func addToRecentPaths(path string) {
	// 移除已存在的相同路径
	for i, recentPath := range config.RecentPaths {
		if recentPath == path {
			config.RecentPaths = append(config.RecentPaths[:i], config.RecentPaths[i+1:]...)
			break
		}
	}

	// 添加到开头
	config.RecentPaths = append([]string{path}, config.RecentPaths...)

	// 限制数量
	if len(config.RecentPaths) > config.MaxRecent {
		config.RecentPaths = config.RecentPaths[:config.MaxRecent]
	}

	// 保存配置
	viper.Set("recent_paths", config.RecentPaths)
	if err := viper.WriteConfig(); err != nil && debug {
		fmt.Printf("⚠️ Warning: cannot save recent paths: %v\n", err)
	}
}

// getCustomManager 获取自定义文件管理器命令
func getCustomManager(managerName string) (string, bool) {
	if cmd, exists := config.CustomManagers[managerName]; exists {
		return cmd, true
	}
	return "", false
}

// isPathValid 检查路径是否有效
func isPathValid(path string) bool {
	if path == "" {
		return false
	}

	// 检查路径是否存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

// formatPath 格式化路径显示
func formatPath(path string) string {
	home, err := os.UserHomeDir()
	if err != nil {
		return path
	}

	// 将用户主目录替换为 ~
	if strings.HasPrefix(path, home) {
		return "~" + strings.TrimPrefix(path, home)
	}

	return path
}

// getFileExtension 获取文件扩展名
func getFileExtension(path string) string {
	ext := filepath.Ext(path)
	if ext == "" {
		return ""
	}
	// 移除点号并转换为小写
	return strings.ToLower(strings.TrimPrefix(ext, "."))
}

// isFile 检查路径是否为文件
func isFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// getAppForFileType 根据文件类型获取对应的应用程序
func getAppForFileType(filePath string) string {
	if !isFile(filePath) {
		return "" // 文件夹使用默认文件管理器
	}

	ext := getFileExtension(filePath)
	if ext == "" {
		return "" // 没有扩展名的文件使用默认程序
	}

	// 从配置中获取应用程序
	if app, exists := config.FileTypeApps[ext]; exists {
		return app
	}

	return "" // 没有配置的文件类型使用默认程序
}

// findSimilarApp 查找相似的应用程序
func findSimilarApp(appName string) (string, bool) {
	switch runtime.GOOS {
	case "darwin":
		// 检查 /Applications 和 /System/Applications 目录
		appDirs := []string{"/Applications", "/System/Applications"}

		for _, appDir := range appDirs {
			entries, err := os.ReadDir(appDir)
			if err != nil {
				continue
			}

			for _, entry := range entries {
				if entry.IsDir() && strings.HasSuffix(entry.Name(), ".app") {
					appNameWithoutExt := strings.TrimSuffix(entry.Name(), ".app")

					// 精确匹配（忽略大小写）
					if strings.EqualFold(appNameWithoutExt, appName) {
						return appNameWithoutExt, true
					}

					// 编辑距离匹配（处理拼写错误，距离不超过2）
					if editDistance(strings.ToLower(appNameWithoutExt), strings.ToLower(appName)) <= 2 {
						return appNameWithoutExt, true
					}

					// 包含匹配（但要求至少3个字符匹配）
					if len(appName) >= 3 && len(appNameWithoutExt) >= 3 {
						if strings.Contains(strings.ToLower(appNameWithoutExt), strings.ToLower(appName)) ||
							strings.Contains(strings.ToLower(appName), strings.ToLower(appNameWithoutExt)) {
							return appNameWithoutExt, true
						}
					}
				}
			}
		}

		// 检查 PATH 中的命令行工具
		if path, err := exec.LookPath(appName); err == nil {
			return filepath.Base(path), true
		}

		// 在 PATH 中查找相似的工具
		pathDirs := strings.Split(os.Getenv("PATH"), ":")
		for _, pathDir := range pathDirs {
			entries, err := os.ReadDir(pathDir)
			if err != nil {
				continue
			}

			for _, entry := range entries {
				if !entry.IsDir() {
					fileName := entry.Name()

					// 精确匹配（忽略大小写）
					if strings.EqualFold(fileName, appName) {
						return fileName, true
					}

					// 编辑距离匹配（距离不超过2）
					if editDistance(strings.ToLower(fileName), strings.ToLower(appName)) <= 2 {
						return fileName, true
					}

					// 包含匹配（但要求至少3个字符匹配）
					if len(appName) >= 3 && len(fileName) >= 3 {
						if strings.Contains(strings.ToLower(fileName), strings.ToLower(appName)) ||
							strings.Contains(strings.ToLower(appName), strings.ToLower(fileName)) {
							return fileName, true
						}
					}
				}
			}
		}

		return "", false
	default:
		// 其他系统只检查 PATH
		if path, err := exec.LookPath(appName); err == nil {
			return filepath.Base(path), true
		}
		return "", false
	}
}

// editDistance 计算两个字符串的编辑距离
func editDistance(s1, s2 string) int {
	len1, len2 := len(s1), len(s2)

	// 创建矩阵
	matrix := make([][]int, len1+1)
	for i := range matrix {
		matrix[i] = make([]int, len2+1)
	}

	// 初始化第一行和第一列
	for i := 0; i <= len1; i++ {
		matrix[i][0] = i
	}
	for j := 0; j <= len2; j++ {
		matrix[0][j] = j
	}

	// 填充矩阵
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if s1[i-1] == s2[j-1] {
				matrix[i][j] = matrix[i-1][j-1]
			} else {
				matrix[i][j] = min(matrix[i-1][j]+1, matrix[i][j-1]+1, matrix[i-1][j-1]+1)
			}
		}
	}

	return matrix[len1][len2]
}

// min 返回三个整数中的最小值
func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

// validateApp 验证应用程序是否存在
func validateApp(appName string) (bool, string) {
	switch runtime.GOOS {
	case "darwin":
		// macOS 检查 /Applications 和 /System/Applications 目录
		appPaths := []string{
			filepath.Join("/Applications", appName+".app"),
			filepath.Join("/System/Applications", appName+".app"),
		}

		for _, appPath := range appPaths {
			if _, err := os.Stat(appPath); err == nil {
				return true, appPath
			}
		}

		// 检查是否在 PATH 中的命令行工具
		if _, err := exec.LookPath(appName); err == nil {
			return true, appName
		}

		// 尝试查找相似的应用程序
		if similarApp, found := findSimilarApp(appName); found {
			return false, fmt.Sprintf("Application '%s' does not exist, but found a similar application '%s'. \nPlease use the correct name: %s", appName, similarApp, similarApp)
		}

		return false, fmt.Sprintf("Application '%s' does not exist. Please ensure:\n  1. The application is installed in the /Applications or /System/Applications directory\n  2. Command line tools are added to the PATH", appName)
	case "windows":
		// Windows 需要完整路径，这里只做基本检查
		if strings.Contains(appName, "/") || strings.Contains(appName, "\\") {
			// 如果是路径，检查文件是否存在
			if _, err := os.Stat(appName); err == nil {
				return true, appName
			}
			return false, fmt.Sprintf("Application path '%s' does not exist", appName)
		}
		// 如果不是路径，提示用户需要完整路径
		return false, "Windows requires a full path for applications. \nFor example: C:\\Program Files\\Notepad++\\notepad++.exe"
	default:
		// Linux 和其他系统检查 PATH
		if _, err := exec.LookPath(appName); err == nil {
			return true, appName
		}
		return false, fmt.Sprintf("Application '%s' does not exist in PATH", appName)
	}
}

// openFileWithApp 使用指定应用程序打开文件
func openFileWithApp(filePath string, appName string) error {
	if debug {
		fmt.Printf("🔍 Opening file with app: %s -> %s\n", filePath, appName)
	}

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		// macOS 使用 open -a 命令
		cmd = exec.Command("open", "-a", appName, filePath)
	case "windows":
		// Windows 使用 start 命令
		cmd = exec.Command("start", appName, filePath)
	default:
		// Linux 和其他系统，尝试使用自定义管理器
		if customCmd, exists := getCustomManager(appName); exists {
			if debug {
				fmt.Printf("🔍 Using custom app: %s -> %s\n", appName, customCmd)
			}
			cmd = exec.Command(customCmd, filePath)
		} else {
			// 使用默认文件管理器
			return openInFileManager(filePath)
		}
	}

	return cmd.Run()
}
