## Language Support

The system includes comprehensive language support with the following features:

1. **Flexible language input**:
   - Supports full language names: "English", "Spanish", "French"
   - Supports ISO language codes: "en", "es", "fr"
   - Supports common variations: "eng", "español", "français"
   - Supports locale variants: "en_US", "fr_CA", "es_MX"

2. **System language detection**:
   - Automatically detects the user's system language
   - Uses environment variables in order: LC_ALL, LC_MESSAGES, LANG, LANGUAGE
   - Falls back to English if no system language is detected

3. **Language fallback**:
   - If a message isn't available in the requested language, tries English
   - Only returns "not found" if the message doesn't exist in any language

## Supported Languages

The system supports these languages (and many more):

| Language | Full Name | Variants |
|----------|-----------|----------|
| en | English | eng, english, en_US, en_GB |
| es | Spanish | esp, español, espanol, spanish |
| fr | French | fra, français, francais, french |
| de | German | deu, deutsch, german |
| it | Italian | ita, italiano, italian |
| pt | Portuguese | por, português, portugues |
| zh | Chinese | zho, mandarin, chinese |
| ja | Japanese | jpn, japanese, 日本語 |
| ko | Korean | kor, korean, 한국어 |
| ru | Russian | rus, russian, русский |

...and many more languages are supported.# Message Database Usage Guide

This lightweight message database system is designed for ultra-fast message retrieval (under 2 milliseconds), with support for multiple languages, system language detection, and preservation of message formatting.

## Admin Tool Instructions

Both Go and Python implementations provide an interactive terminal interface with the following features:

### Main Menu Options

```
Message Database Admin Tool
----------------------------
System Language: en
1. Add message
2. List messages
3. Delete message
4. Get message using system language
0. Exit
```

### Adding Messages

When adding messages:

1. You'll be prompted for:
   - Message ID: A unique identifier for the message
   - Language: Either the language name or code (e.g., "English" or "en")
   - Message content: The actual message text

2. **Language Input Flexibility**:
   - You can enter full language names ("English", "Spanish", "French", etc.)
   - You can enter ISO codes ("en", "es", "fr", etc.)
   - You can enter common variations ("eng", "español", "français", etc.)

3. **Multi-line message input**:
   - Enter your message text
   - The tool preserves all line endings (`\r`, `\n`, or `\r\n`)
   - Type `$%` on a new line to finish message input

Example:
```
Enter message ID: welcome_message
Enter language (name or code, e.g. 'english' or 'en'): spanish
Using language code: es
Enter message content (type $% on a new line when finished):
¡Bienvenido a nuestro sistema!
Esta es una mensaje de múltiples líneas
con formato preservado.
$%
Message added successfully!
```

### Adding Messages

When adding messages:

1. You'll be prompted for:
   - Message ID: A unique identifier for the message
   - Language code (required): Standard ISO language code (en, fr, de, etc.)
   - Message content: The actual message text

2. **Multi-line message input**:
   - Enter your message text
   - The tool preserves all line endings (`\r`, `\n`, or `\r\n`)
   - Type `$$%` on a new line to finish message input

Example:
```
Enter message ID: welcome_message
Enter language code (required): en
Enter message content (type $$% on a new line when finished):
Welcome to our system!
This is a multi-line message
with preserved formatting.
$$%
Message added successfully!
```

### Listing Messages

When listing messages:
- Enter a language code to filter messages by language
- Leave empty to list all messages

### Deleting Messages

When deleting messages:
- Enter the language code and message ID
- Confirm deletion when prompted

## API Reference

### Go API Usage

```go
// Initialize database
db, err := messagedb.NewMessageDB("messages.db")
if err != nil {
    log.Fatal(err)
}
defer db.Close()

// Retrieve a message with specific language
content, err := db.GetMessage("en", "welcome_message")
if err != nil {
    log.Printf("Error: %v", err)
} else {
    fmt.Printf("Message: %s\n", content)
}

// Retrieve a message using system language
content, err = db.GetSystemMessage("welcome_message")
if err != nil {
    log.Printf("Error: %v", err)
} else {
    fmt.Printf("System language message: %s\n", content)
}
```

### Python API Usage

```python
from message_db import MessageDB

# Initialize database
db = MessageDB("messages.db")

# Retrieve a message with specific language
content = db.get_message("en", "welcome_message")
if content:
    print(f"Message: {content}")
else:
    print("Message not found")

# Retrieve a message using system language
system_content = db.get_system_message("welcome_message")
if system_content:
    print(f"System language message: {system_content}")
else:
    print("Message not found in system language")

# Close the connection
db.close()
```

## System Language Detection

Both implementations automatically detect the system language by checking environment variables in this order:

1. `LC_ALL`
2. `LC_MESSAGES`
3. `LANG`
4. `LANGUAGE`

If none of these variables are set, English (`en`) is used as the default language. The system will always fall back to English messages if a message isn't available in the detected language.


## Notes on Message Formatting

1. The database preserves all line endings:
   - Unix-style (`\n`)
   - Windows-style (`\r\n`)
   - Classic Mac-style (`\r`)

2. The admin tools handle multi-line input with the `$$%` terminator to make it easy to enter complex messages with preserved formatting.
