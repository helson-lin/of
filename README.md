# of - 智能文件打开器

[English](README_EN.md) | 中文

一个强大而智能的命令行工具，用于使用正确的应用程序打开文件和目录。

## ✨ 功能特性

- 🚀 **跨平台支持**: 支持 macOS、Windows 和 Linux
- 🧠 **智能文件类型检测**: 自动为每种文件类型选择最佳应用程序
- 🔧 **高度可定制**: 配置您自己的文件类型映射和管理器
- 📝 **最近历史记录**: 跟踪最近打开的路径
- ⚙️ **简单配置**: 基于 YAML 的简单配置
- 🐛 **调试模式**: 详细的调试信息
- 📦 **丰富命令**: 全面的子命令系统
- 🔍 **智能建议**: 自动纠正应用程序名称中的拼写错误

## 🚀 快速开始

### 安装

```bash
# 使用 Homebrew (macOS)
brew install helson-lin/tap/of

# 或从发布页面下载
# 访问: https://github.com/helson-lin/of/releases
```

### 基本用法

```bash
# 打开当前目录
of

# 打开指定路径
of /path/to/folder

# 使用标志打开文件
of -p /path/to/file

# 指定文件管理器
of -m finder /path/to/folder

# 启用调试模式
of --debug /path/to/folder
```

## 📋 命令

### 主要命令

```bash
# 打开文件/目录
of [path] [flags]

# 显示帮助
of --help

# 显示版本
of version
```

### 配置命令

```bash
# 显示当前配置
of config show

# 添加自定义文件管理器
of config add-manager "vscode" "code"

# 设置默认文件管理器
of config set-default "finder"

# 清除最近路径
of config clear-recent

# 列出最近路径
of list
```

### 文件类型管理

```bash
# 添加单个文件类型映射
of config add-filetype "txt" "TextEdit"
of config add-filetype "py" "code"

# 添加文件类型组（批量映射）
of config add-filegroup "audio" "IINA"
of config add-filegroup "video" "IINA"
of config add-filegroup "image" "Preview"
of config add-filegroup "code" "vscode"

# 列出所有文件类型映射
of config list-filetypes

# 移除文件类型映射
of config remove-filetype "txt"
```

## ⚙️ 配置

配置文件: `~/.of/config.yaml`

### 配置示例

```yaml
default_manager: ""
custom_managers: {}
recent_paths:
  - "/Users/username/Documents"
  - "/Users/username/Downloads"
max_recent: 10
file_type_apps:
  txt: "TextEdit"
  md: "vscode"
  py: "code"
  jpg: "Preview"
  mp4: "IINA"
  mp3: "IINA"
```

## 🧠 智能功能

### 自动纠正

当您输入拼写错误时，工具会自动建议正确的应用程序名称：

```bash
# 拼写错误: "cusor" → 自动建议: "Cursor"
of config add-filetype txt cusor
# 输出: Application 'cusor' does not exist, but found a similar application 'Cursor'.
#        Please use the correct name: Cursor

# 拼写错误: "previe" → 自动建议: "Preview"
of config add-filetype txt previe
# 输出: Application 'previe' does not exist, but found a similar application 'Preview'.
#        Please use the correct name: Preview
```

### 文件类型组

快速配置多种文件类型：

```bash
# 音频文件 (mp3, wav, flac, aac, ogg, m4a, wma)
of config add-filegroup audio IINA

# 视频文件 (mp4, avi, mkv, mov, wmv, flv, webm, m4v, 3gp)
of config add-filegroup video IINA

# 图片文件 (jpg, jpeg, png, gif, bmp, svg, tiff, webp)
of config add-filegroup image Preview

# 代码文件 (py, js, ts, go, java, cpp, c, h, html, css, json, xml, yaml, yml)
of config add-filegroup code vscode

# 文档文件 (pdf, doc, docx, txt, md, rtf)
of config add-filegroup document TextEdit
```

## 🔧 平台支持

### macOS
- 使用 `open -a` 打开应用程序
- 检查 `/Applications` 和 `/System/Applications`
- 支持 PATH 中的命令行工具
- 自动纠正应用程序名称

### Windows
- 使用 `start` 命令
- 需要完整的应用程序路径
- 示例: `C:\Program Files\Notepad++\notepad++.exe`

### Linux
- 使用自定义管理器命令
- 检查 PATH 中的命令行工具
- 回退到默认文件管理器

## 🐛 调试模式

启用调试模式查看详细的执行信息：

```bash
of --debug /path/to/file
```

调试输出包括：
- 操作系统检测
- 配置文件加载
- 文件类型检测
- 应用程序选择
- 命令执行详情

## 📝 示例

### 开发工作流

```bash
# 在 VS Code 中打开项目
of config add-filetype py code
of config add-filetype js code
of config add-filetype go code

# 打开文件
of main.py      # 在 VS Code 中打开
of app.js       # 在 VS Code 中打开
of server.go    # 在 VS Code 中打开
```

### 媒体管理

```bash
# 配置媒体播放器
of config add-filegroup audio IINA
of config add-filegroup video IINA
of config add-filegroup image Preview

# 打开媒体文件
of song.mp3     # 在 IINA 中打开
of video.mp4    # 在 IINA 中打开
of photo.jpg    # 在 Preview 中打开
```

### 文档工作流

```bash
# 配置文档应用程序
of config add-filetype pdf Preview
of config add-filetype docx Pages
of config add-filetype xlsx Numbers

# 打开文档
of report.pdf   # 在 Preview 中打开
of document.docx # 在 Pages 中打开
of data.xlsx    # 在 Numbers 中打开
```

## 🛠️ 开发

### 构建

```bash
# 构建可执行文件
go build -o of main.go

# 为特定平台构建
GOOS=windows GOARCH=amd64 go build -o of.exe main.go
GOOS=darwin GOARCH=arm64 go build -o of main.go
```

### 测试

```bash
# 运行测试
go test ./...

# 使用调试模式运行
./of --debug test.txt
```

## 📋 文件类型组

可用于批量配置的文件类型组：

| 组 | 扩展名 | 描述 |
|-------|------------|-------------|
| `audio` | mp3, wav, flac, aac, ogg, m4a, wma | 音频文件 |
| `video` | mp4, avi, mkv, mov, wmv, flv, webm, m4v, 3gp | 视频文件 |
| `image` | jpg, jpeg, png, gif, bmp, svg, tiff, webp | 图片文件 |
| `document` | pdf, doc, docx, txt, md, rtf | 文档文件 |
| `code` | py, js, ts, go, java, cpp, c, h, html, css, json, xml, yaml, yml | 代码文件 |
| `archive` | zip, rar, 7z, tar, gz, bz2 | 压缩文件 |
| `spreadsheet` | xls, xlsx, csv | 电子表格文件 |
| `presentation` | ppt, pptx | 演示文件 |

## 🤝 贡献

1. Fork 仓库
2. 创建功能分支
3. 进行您的更改
4. 如果适用，添加测试
5. 提交拉取请求

## 📄 许可证

MIT 许可证 - 详情请参阅 [LICENSE](LICENSE) 文件。

## 🆕 更新日志

### v0.0.1
- 初始发布
- 跨平台支持
- 智能文件类型检测
- 应用程序名称自动纠正
- 文件类型组批量配置
- 全面的配置系统 