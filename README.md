# MessageDB: Lightweight Multi-language Message Database

A high-performance, lightweight message storage system designed for applications that need rapid retrieval (< 2ms) of messages in multiple languages.

## Overview

MessageDB provides a simple, SQLite-based solution for storing and retrieving text messages in multiple languages. It's perfect for:

- Application internationalization (i18n)
- Configurable system messages
- Error message catalogs
- Multi-language UI text
- Template messaging systems

## Features

- **Ultra-fast retrieval**: < 2ms access time for small message sets
- **Multi-language support**: Store messages in any language
- **Flexible language input**: Use language names OR language codes
- **System language detection**: Automatically use the user's system language
- **Cross-platform**: Available in both Go and Python
- **Simple API**: Easy to integrate into any application
- **Admin tools**: Command-line tools for message management
- **JSON import/export**: Tools for bulk operations and backup

## Quick Start

### Go

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/yourusername/messagedb"
)

func main() {
    // Initialize database
    db, err := messagedb.NewMessageDB("messages.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    // Add a message
    err = db.AddMessage("greeting", "en", "Hello, world!")
    if err != nil {
        log.Printf("Error adding message: %v", err)
    }
    
    // Add same message in another language
    err = db.AddMessage("greeting", "es", "¡Hola, mundo!")
    if err != nil {
        log.Printf("Error adding message: %v", err)
    }
    
    // Retrieve a message with specific language
    content, err := db.GetMessage("en", "greeting")
    if err != nil {
        log.Printf("Error: %v", err)
    } else {
        fmt.Printf("Message: %s\n", content)
    }
    
    // Retrieve a message using system language
    content, err = db.GetSystemMessage("greeting")
    if err != nil {
        log.Printf("Error: %v", err)
    } else {
        fmt.Printf("System language message: %s\n", content)
    }
}
```

### Python

```python
from message_db import MessageDB

# Initialize database
db = MessageDB("messages.db")

# Add messages
db.add_message("greeting", "en", "Hello, world!")
db.add_message("greeting", "fr", "Bonjour, monde!")

# Retrieve messages
en_message = db.get_message("en", "greeting")
print(f"English greeting: {en_message}")

# Use system language
sys_message = db.get_system_message("greeting")
print(f"System language greeting: {sys_message}")

# Close connection
db.close()
```

## Installation

### Prerequisites

- SQLite (Version 3.x or higher)
- Go (1.16+) for Go implementation
- Python (3.6+) for Python implementation

### Go Implementation

```bash
# Install the SQLite driver
go get github.com/mattn/go-sqlite3

# Build the admin tool
cd cmd/msgadmin
go build -o msgadmin
```

### Python Implementation

```bash
# No external dependencies required
# Make the admin tool executable
chmod +x message_admin.py
```

## Admin Tool

The admin tool provides an interactive interface for managing messages:

```bash
# Go
./msgadmin

# Python
python message_admin.py -i
```

It offers the following features:
- Add messages with multi-line support
- List messages by language
- Delete messages
- Get messages using system language

## JSON Import/Export

Tools for bulk operations in JSON format:

```bash
# Export
python message_json_tool.py --db messages.db export --output messages.json

# Import
python message_json_tool.py --db messages.db import --input messages.json
```

## Language Support

The system supports any language and provides flexible input options:

```python
# All of these will work
db.get_message("en", "greeting")
db.get_message("english", "greeting")
db.get_message("eng", "greeting")
db.get_message("en_US", "greeting")
```

## System Language Detection

The system automatically detects the user's language from environment variables:

```go
// Order of precedence
// 1. LC_ALL
// 2. LC_MESSAGES
// 3. LANG
// 4. LANGUAGE
// 5. Default to "en" (English)
```

## Performance

MessageDB is designed for high-performance scenarios where message retrieval time is critical:

- Utilizes SQLite's WAL mode for fast reads
- Creates proper indexes for < 2ms retrieval times
- Optimized for small to medium message sets (< 1000 messages)

## Files and Components

### Go Implementation:
- `messagedb.go` - Core database API
- `language_util.go` - Language mapping utilities
- `admin_tool.go` - Interactive admin tool
- `msgdbjson.go` - JSON import/export tool

### Python Implementation:
- `message_db.py` - Core database API
- `language_util.py` - Language mapping utilities
- `message_admin.py` - Interactive admin tool
- `message_json_tool.py` - JSON import/export tool

## License

[MIT License](LICENSE)
