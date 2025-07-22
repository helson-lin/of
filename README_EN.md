# of - Smart File Opener

[English](README_EN.md) | 中文

A powerful and intelligent command-line tool for opening files and directories with the right applications.

## ✨ Features

- 🚀 **Cross-platform**: Works on macOS, Windows, and Linux
- 🧠 **Smart file type detection**: Automatically chooses the best app for each file type
- 🔧 **Customizable**: Configure your own file type mappings and managers
- 📝 **Recent history**: Tracks recently opened paths
- ⚙️ **Easy configuration**: Simple YAML-based configuration
- 🐛 **Debug mode**: Detailed debugging information
- 📦 **Rich commands**: Comprehensive subcommand system
- 🔍 **Smart suggestions**: Auto-corrects typos in application names
- 📋 **Clipboard support**: One-click copy path to system clipboard

## 🚀 Quick Start

### Installation

```bash
# Using Homebrew (macOS)
brew install helson-lin/tap/of

# Or download from releases
# Visit: https://github.com/helson-lin/of/releases
```

### Basic Usage

```bash
# Open current directory
of

# Open specific path
of /path/to/folder

# Open file with flag
of -p /path/to/file

# Specify file manager
of -m finder /path/to/folder

# Enable debug mode
of --debug /path/to/folder

# Copy path to clipboard
of --copy
of --copy /path/to/folder
```

## 📋 Clipboard Functionality

### Copy Path to Clipboard

You can use the `--copy` or `-c` flag to copy a path to the system clipboard:

```bash
# Copy current directory path to clipboard
of --copy

# Copy specified path to clipboard
of --copy /path/to/folder

# Use short flag
of -c /path/to/file
```

This feature is particularly useful in scenarios such as:
- Need to paste path into other applications
- Share file path with colleagues
- Use path in scripts
- Quickly get absolute path

## 📋 Commands

### Main Commands

```bash
# Open files/directories
of [path] [flags]

# Show help
of --help

# Show version
of version
```

### Configuration Commands

```bash
# Show current configuration
of config show

# Add custom file manager
of config add-manager "vscode" "code"

# Set default file manager
of config set-default "finder"

# Clear recent paths
of config clear-recent

# List recent paths
of list
```

### File Type Management

```bash
# Add single file type mapping
of config add-filetype "txt" "TextEdit"
of config add-filetype "py" "code"

# Add file type group (batch mapping)
of config add-filegroup "audio" "IINA"
of config add-filegroup "video" "IINA"
of config add-filegroup "image" "Preview"
of config add-filegroup "code" "vscode"

# List all file type mappings
of config list-filetypes

# Remove file type mapping
of config remove-filetype "txt"
```

## ⚙️ Configuration

Configuration file: `~/.of/config.yaml`

### Example Configuration

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

## 🧠 Smart Features

### Auto-correction

The tool automatically suggests correct application names when you make typos:

```bash
# Typo: "cusor" → Auto-suggests: "Cursor"
of config add-filetype txt cusor
# Output: Application 'cusor' does not exist, but found a similar application 'Cursor'.
#         Please use the correct name: Cursor

# Typo: "previe" → Auto-suggests: "Preview"
of config add-filetype txt previe
# Output: Application 'previe' does not exist, but found a similar application 'Preview'.
#         Please use the correct name: Preview
```

### File Type Groups

Quickly configure multiple file types at once:

```bash
# Audio files (mp3, wav, flac, aac, ogg, m4a, wma)
of config add-filegroup audio IINA

# Video files (mp4, avi, mkv, mov, wmv, flv, webm, m4v, 3gp)
of config add-filegroup video IINA

# Image files (jpg, jpeg, png, gif, bmp, svg, tiff, webp)
of config add-filegroup image Preview

# Code files (py, js, ts, go, java, cpp, c, h, html, css, json, xml, yaml, yml)
of config add-filegroup code vscode

# Document files (pdf, doc, docx, txt, md, rtf)
of config add-filegroup document TextEdit
```

## 🔧 Platform Support

### macOS
- Uses `open -a` for applications
- Checks `/Applications` and `/System/Applications`
- Supports command-line tools in PATH
- Auto-corrects application names

### Windows
- Uses `start` command
- Requires full application paths
- Example: `C:\Program Files\Notepad++\notepad++.exe`

### Linux
- Uses custom manager commands
- Checks PATH for command-line tools
- Falls back to default file manager

## 🐛 Debug Mode

Enable debug mode to see detailed execution information:

```bash
of --debug /path/to/file
```

Debug output includes:
- Operating system detection
- Configuration file loading
- File type detection
- Application selection
- Command execution details

## 📝 Examples

### Development Workflow

```bash
# Open project in VS Code
of config add-filetype py code
of config add-filetype js code
of config add-filetype go code

# Open files
of main.py      # Opens in VS Code
of app.js       # Opens in VS Code
of server.go    # Opens in VS Code
```

### Media Management

```bash
# Configure media players
of config add-filegroup audio IINA
of config add-filegroup video IINA
of config add-filegroup image Preview

# Open media files
of song.mp3     # Opens in IINA
of video.mp4    # Opens in IINA
of photo.jpg    # Opens in Preview
```

### Document Workflow

```bash
# Configure document apps
of config add-filetype pdf Preview
of config add-filetype docx Pages
of config add-filetype xlsx Numbers

# Open documents
of report.pdf   # Opens in Preview
of document.docx # Opens in Pages
of data.xlsx    # Opens in Numbers
```

## 🛠️ Development

### Building

```bash
# Build executable
go build -o of main.go

# Build for specific platform
GOOS=windows GOARCH=amd64 go build -o of.exe main.go
GOOS=darwin GOARCH=arm64 go build -o of main.go
```

### Testing

```bash
# Run tests
go test ./...

# Run with debug
./of --debug test.txt
```

## 📋 File Type Groups

Available file type groups for batch configuration:

| Group | Extensions | Description |
|-------|------------|-------------|
| `audio` | mp3, wav, flac, aac, ogg, m4a, wma | Audio files |
| `video` | mp4, avi, mkv, mov, wmv, flv, webm, m4v, 3gp | Video files |
| `image` | jpg, jpeg, png, gif, bmp, svg, tiff, webp | Image files |
| `document` | pdf, doc, docx, txt, md, rtf | Document files |
| `code` | py, js, ts, go, java, cpp, c, h, html, css, json, xml, yaml, yml | Code files |
| `archive` | zip, rar, 7z, tar, gz, bz2 | Archive files |
| `spreadsheet` | xls, xlsx, csv | Spreadsheet files |
| `presentation` | ppt, pptx | Presentation files |

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📄 License

MIT License - see [LICENSE](LICENSE) file for details.

## 🆕 Changelog

### v0.0.1
- Initial release
- Cross-platform support
- Smart file type detection
- Auto-correction for application names
- File type groups for batch configuration
- Comprehensive configuration system 