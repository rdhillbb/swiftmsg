# Message Database Installation and Operation Guide

This guide provides step-by-step instructions for installing and running the Message Database system in both Go and Python.

## Table of Contents
- [Prerequisites](#prerequisites)
- [Go Implementation](#go-implementation)
- [Python Implementation](#python-implementation)
- [Using the Admin Tool](#using-the-admin-tool)
- [Common Issues](#common-issues)

## Prerequisites

### Required Software
- SQLite (Version 3.x or higher)
- Go (1.16+) for Go implementation
- Python (3.6+) for Python implementation

## Installing SQLite

### Windows
1. Download the pre-compiled binaries from [SQLite Download Page](https://www.sqlite.org/download.html)
2. Extract the ZIP file to a folder (e.g., `C:\sqlite`)
3. Add the folder to your system PATH:
   - Right-click on 'This PC' > Properties > Advanced System Settings > Environment Variables
   - Add `C:\sqlite` to the PATH variable

### macOS
SQLite comes pre-installed on macOS. To verify the installation:
```bash
sqlite3 --version
```

If needed, install or upgrade using Homebrew:
```bash
brew install sqlite
```

### Linux (Ubuntu/Debian)
```bash
sudo apt-get update
sudo apt-get install sqlite3 libsqlite3-dev
```

### Linux (CentOS/RHEL/Fedora)
```bash
sudo dnf install sqlite sqlite-devel    # For Fedora/RHEL 8+
# OR
sudo yum install sqlite sqlite-devel    # For CentOS/RHEL 7
```

## Go Implementation

### Installation

1. **Install Go:**
   - Download from [Go's official website](https://golang.org/dl/)
   - Follow the installation instructions for your platform

2. **Install the SQLite driver for Go:**
   ```bash
   go get github.com/mattn/go-sqlite3
   ```

3. **Create a new project directory:**
   ```bash
   mkdir -p message-db/cmd/msgadmin
   cd message-db
   ```

4. **Initialize Go module:**
   ```bash
   go mod init github.com/yourusername/messagedb
   ```

5. **Save the Go source files:**
   - Save `messagedb.go` in the root directory
   - Save `language_util.go` in the root directory
   - Save `admin_tool.go` in the `cmd/msgadmin` directory

6. **Build the admin tool:**
   ```bash
   cd cmd/msgadmin
   go build -o msgadmin
   ```

### Running the Admin Tool (Go)

```bash
# Interactive mode
./msgadmin

# Command-line mode
./msgadmin --db=messages.db --list=en
./msgadmin --db=messages.db --add
./msgadmin --db=messages.db --delete
./msgadmin --db=messages.db --sysget
```

## Python Implementation

### Installation

1. **Install Python:**
   - Download from [Python's official website](https://www.python.org/downloads/)
   - Follow the installation instructions for your platform

2. **SQLite is included in Python's standard library, no additional installation needed.**

3. **Create a new project directory:**
   ```bash
   mkdir message-db-python
   cd message-db-python
   ```

4. **Save the Python source files:**
   - Save `message_db.py` in the project directory
   - Save `language_util.py` in the project directory
   - Save `message_admin.py` in the project directory

5. **Make the admin tool executable (Linux/macOS):**
   ```bash
   chmod +x message_admin.py
   ```

### Running the Admin Tool (Python)

```bash
# Interactive mode
./message_admin.py -i
# or
python message_admin.py -i

# Command-line mode
python message_admin.py --db=messages.db list
python message_admin.py --db=messages.db add --id=welcome --lang=english --content="Welcome"
python message_admin.py --db=messages.db delete --id=welcome --lang=english
python message_admin.py --db=messages.db stats
```

## Using the Admin Tool

### Interactive Mode Options

1. **Add message:**
   - Enter a unique message ID
   - Enter language (name or code)
   - Enter multi-line message content, ending with `$$%` on a new line

2. **List messages:**
   - Enter language to filter messages (or leave empty to list all)

3. **Delete message:**
   - Enter message ID and language
   - Confirm deletion

4. **Get message using system language:**
   - Enter message ID
   - System will use detected language with English fallback

### Sample Database Structure

The message database is a single SQLite file that contains:
- One table: `messages`
- Indexes for fast retrieval
- Storage for multiple languages

To view the database directly using SQLite CLI:
```bash
sqlite3 messages.db
> .tables
> SELECT * FROM messages;
> .exit
```

## Common Issues

### Go Implementation Issues

1. **Error: "gcc: command not found"**
   - Install GCC or a C compiler required by cgo:
     - Windows: Install MinGW or TDM-GCC
     - macOS: `xcode-select --install`
     - Linux: `sudo apt install build-essential` or `sudo dnf install gcc`

2. **Error: "sqlite3.h: No such file or directory"**
   - Install SQLite development files:
     - Windows: Included in the MinGW or TDM-GCC installation
     - macOS: `brew install sqlite`
     - Linux: `sudo apt install libsqlite3-dev` or `sudo dnf install sqlite-devel`

### Python Implementation Issues

1. **Error: "sqlite3.OperationalError: unable to open database file"**
   - Check permissions on the directory
   - Ensure the path to the database file is valid
   - Try using an absolute path to the database file

2. **Error: "ImportError: No module named 'sqlite3'"**
   - Reinstall Python with SQLite support
   - On some Linux distributions: `sudo apt install python3-sqlite3`
