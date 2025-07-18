# of - Open File Manager

[English](README_EN.md) | 中文

一个简单而强大的命令行工具，用于在系统默认文件管理器中打开文件或文件夹。

## ✨ 功能特性

- 🚀 **跨平台支持**: 支持 macOS、Windows 和 Linux
- 🔧 **自定义文件管理器**: 支持配置自定义文件管理器
- 📝 **最近使用记录**: 自动记录最近打开的路径
- ⚙️ **配置管理**: 支持配置文件管理
- 🐛 **调试模式**: 提供详细的调试信息
- 📦 **子命令系统**: 提供丰富的子命令功能
- 📄 **文件类型识别**: 根据文件类型自动选择打开程序

## 🚀 快速开始

### 基本用法

```bash
# 打开当前目录
of

# 打开指定路径
of /path/to/folder

# 使用标志指定路径
of -p /path/to/file

# 指定文件管理器
of -m finder /path/to/folder

# 启用调试模式
of --debug /path/to/folder
```

### 子命令

#### 查看最近使用的路径
```bash
of list
```

#### 配置管理
```bash
# 查看当前配置
of config show

# 添加自定义文件管理器
of config add-manager "vscode" "code"

# 设置默认文件管理器
of config set-default "finder"

# 清除最近使用记录
of config clear-recent

# 添加文件类型映射
of config add-filetype "txt" "vscode"
of config add-filetype "xlsx" "wps"

# 查看文件类型映射
of config list-filetypes

# 移除文件类型映射
of config remove-filetype "txt"
```

#### 版本信息
```bash
of version
```

## ⚙️ 配置

配置文件位置: `~/.of/config.yaml`

### 配置示例

```yaml
default_manager: "finder"
custom_managers:
  vscode: "code"
  sublime: "subl"
  nautilus: "nautilus"
recent_paths:
  - "/Users/username/Documents"
  - "/Users/username/Downloads"
max_recent: 10
file_type_apps:
  txt: "vscode"
  md: "vscode"
  json: "vscode"
  xlsx: "wps"
  docx: "wps"
  pdf: "preview"
  jpg: "preview"
```

## 🔧 自定义文件管理器

### 添加自定义管理器

```bash
# 添加 VS Code 作为文件管理器
of config add-manager "vscode" "code"

# 添加 Sublime Text
of config add-manager "sublime" "subl"

# 添加 Nautilus (Linux)
of config add-manager "nautilus" "nautilus"
```

### 使用自定义管理器

```bash
# 使用 VS Code 打开文件夹
of -m vscode /path/to/folder

# 设置为默认管理器
of config set-default "vscode"
```

## 📄 文件类型识别

`of` 工具能够根据文件类型自动选择最合适的应用程序打开文件：

### 默认文件类型映射

- **文本文件**: `.txt`, `.md`, `.json`, `.yaml`, `.yml`, `.xml`, `.html`, `.css`, `.js`, `.ts`, `.py`, `.go`, `.java`, `.cpp`, `.c`, `.h` → VS Code
- **Office 文件**: `.xlsx`, `.xls`, `.docx`, `.doc`, `.pptx`, `.ppt` → WPS
- **媒体文件**: `.pdf`, `.jpg`, `.jpeg`, `.png`, `.gif`, `.bmp`, `.svg` → 预览程序

### 自定义文件类型映射

```bash
# 添加新的文件类型映射
of config add-filetype "log" "vscode"
of config add-filetype "csv" "excel"
of config add-filetype "mp4" "vlc"

# 查看所有文件类型映射
of config list-filetypes

# 移除文件类型映射
of config remove-filetype "log"
```

### 工作原理

1. **文件检测**: 工具会检查目标路径是否为文件
2. **扩展名识别**: 提取文件扩展名并转换为小写
3. **应用查找**: 在配置中查找对应的应用程序
4. **程序启动**: 使用找到的应用程序打开文件
5. **回退机制**: 如果没有配置，使用系统默认程序打开

## 🐛 调试模式

启用调试模式可以查看详细的执行信息：

```bash
of --debug /path/to/folder
```

调试信息包括：
- 操作系统信息
- 使用的文件管理器
- 配置文件加载状态
- 命令执行详情

## 📋 支持的文件管理器

### macOS
- Finder (默认)
- VS Code
- Sublime Text
- Terminal

### Windows
- Explorer (默认)
- VS Code
- Total Commander
- Directory Opus

### Linux
- xdg-open (默认)
- Nautilus
- Dolphin
- Thunar

## 🛠️ 构建

```bash
# 构建可执行文件
go build -o of

# 安装到系统
go install
```

## 📝 更新日志

### v0.0.1
- 基础功能实现
- 跨平台支持
- 配置文件支持
- 自定义文件管理器

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## �� 许可证

MIT License 