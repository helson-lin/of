# of - Open File Manager

English | [中文](README.md)

A simple yet powerful command-line tool for opening files and directories in the system's default file manager.

## ✨ Features

- 🚀 **Cross-platform Support**: Supports macOS, Windows, and Linux
- 🔧 **Custom File Managers**: Support for configuring custom file managers
- 📝 **Recent Usage History**: Automatically records recently opened paths
- ⚙️ **Configuration Management**: Support for configuration file management
- 🐛 **Debug Mode**: Provides detailed debug information
- 📦 **Subcommand System**: Rich subcommand functionality
- 📄 **File Type Recognition**: Automatically select opening programs based on file type

## 🚀 Quick Start

### Basic Usage

```bash
# Open current directory
of

# Open specified path
of /path/to/folder

# Use flag to specify path
of -p /path/to/file

# Specify file manager
of -m finder /path/to/folder

# Enable debug mode
of --debug /path/to/folder
```

### Subcommands

#### View recently used paths
```bash
of list
```

#### Configuration Management
```bash
# View current configuration
of config show

# Add custom file manager
of config add-manager "vscode" "code"

# Set default file manager
of config set-default "finder"

# Clear recent usage history
of config clear-recent

# Add file type mapping
of config add-filetype "txt" "vscode"
of config add-filetype "xlsx" "wps"

# View file type mappings
of config list-filetypes

# Remove file type mapping
of config remove-filetype "txt"
```

#### Version Information
```bash
of version
```

## ⚙️ Configuration

Configuration file location: `~/.of/config.yaml`

### Configuration Example

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

## 🔧 Custom File Managers

### Adding Custom Managers

```bash
# Add VS Code as file manager
of config add-manager "vscode" "code"

# Add Sublime Text
of config add-manager "sublime" "subl"

# Add Nautilus (Linux)
of config add-manager "nautilus" "nautilus"
```

### Using Custom Managers

```bash
# Open folder with VS Code
of -m vscode /path/to/folder

# Set as default manager
of config set-default "vscode"
```

## 📄 File Type Recognition

The `of` tool can automatically select the most appropriate application to open files based on their type:

### Default File Type Mappings

- **Text Files**: `.txt`, `.md`, `.json`, `.yaml`, `.yml`, `.xml`, `.html`, `.css`, `.js`, `.ts`, `.py`, `.go`, `.java`, `.cpp`, `.c`, `.h` → VS Code
- **Office Files**: `.xlsx`, `.xls`, `.docx`, `.doc`, `.pptx`, `.ppt` → WPS
- **Media Files**: `.pdf`, `.jpg`, `.jpeg`, `.png`, `.gif`, `.bmp`, `.svg` → Preview Program

### Custom File Type Mappings

```bash
# Add new file type mappings
of config add-filetype "log" "vscode"
of config add-filetype "csv" "excel"
of config add-filetype "mp4" "vlc"

# View all file type mappings
of config list-filetypes

# Remove file type mapping
of config remove-filetype "log"
```

### How It Works

1. **File Detection**: Tool checks if the target path is a file
2. **Extension Recognition**: Extracts file extension and converts to lowercase
3. **Application Lookup**: Searches for corresponding application in configuration
4. **Program Launch**: Opens file with the found application
5. **Fallback Mechanism**: Uses system default program if not configured

## 🐛 Debug Mode

Enable debug mode to view detailed execution information:

```bash
of --debug /path/to/folder
```

Debug information includes:
- Operating system information
- File manager being used
- Configuration file loading status
- Command execution details

## 📋 Supported File Managers

### macOS
- Finder (default)
- VS Code
- Sublime Text
- Terminal

### Windows
- Explorer (default)
- VS Code
- Total Commander
- Directory Opus

### Linux
- xdg-open (default)
- Nautilus
- Dolphin
- Thunar

## 🛠️ Building

```bash
# Build executable
go build -o of

# Install to system
go install
```

## 📝 Changelog

### v0.0.1
- Basic functionality implementation
- Cross-platform support
- Configuration file support
- Custom file managers
- Recent usage history
- Debug mode
- Subcommand system
- File type recognition

## 🤝 Contributing

Welcome to submit Issues and Pull Requests!

## 📄 License

MIT License 