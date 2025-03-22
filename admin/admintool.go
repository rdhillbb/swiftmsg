package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	messagedb "github.com/rdhillbb/swiftmsg/swiftmsgapi" // Alias the import to messagedb
)

func main() {
	// Command line arguments
	dbPath := flag.String("db", "messages.db", "Path to SQLite database file")
	listLang := flag.String("list", "", "List all messages for the specified language code")
	listAll := flag.Bool("listall", false, "List all messages in the database")
	addMsg := flag.Bool("add", false, "Add a new message")
	deleteMsg := flag.Bool("delete", false, "Delete a message")
	getMessage := flag.Bool("get", false, "Get a specific message")
	getSysMsg := flag.Bool("sysget", false, "Get a message using system language")

	flag.Parse()

	// Initialize database
	db, err := messagedb.NewMessageDB(*dbPath)
	if err != nil {
		fmt.Printf("Error opening database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	// Execute the requested command
	if *listAll {
		listAllMessages(db)
	} else if *listLang != "" {
		listMessagesForLang(db, *listLang)
	} else if *addMsg {
		addMessage(db)
	} else if *deleteMsg {
		deleteMessage(db)
	} else if *getMessage {
		getMessageContent(db)
	} else if *getSysMsg {
		getSystemMessageContent(db)
	} else {
		// Interactive mode if no specific command is provided
		interactiveMode(db)
	}
}

func listAllMessages(db *messagedb.MessageDB) {
	messages, err := db.ListAllMessages()
	if err != nil {
		fmt.Printf("Error listing messages: %v\n", err)
		return
	}

	if len(messages) == 0 {
		fmt.Println("No messages found in the database.")
		return
	}

	fmt.Println("All messages in the database:")
	fmt.Println("-----------------------------")
	for _, msg := range messages {
		fmt.Printf("ID: %s | Language: %s | Content: %s\n", msg.ID, msg.LangCode, msg.Content)
	}
}

func listMessagesForLang(db *messagedb.MessageDB, langInput string) {
	// Convert language name to standard code
	langCode := messagedb.GetLanguageCode(langInput)

	messages, err := db.ListMessages(langCode)
	if err != nil {
		fmt.Printf("Error listing messages: %v\n", err)
		return
	}

	if len(messages) == 0 {
		fmt.Printf("No messages found for language '%s' (code: %s).\n", langInput, langCode)
		return
	}

	fmt.Printf("Messages for language '%s' (code: %s):\n", langInput, langCode)
	fmt.Println("-----------------------------")
	for _, msg := range messages {
		fmt.Printf("ID: %s | Content: %s\n", msg.ID, msg.Content)
	}
}

func addMessage(db *messagedb.MessageDB) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter message ID: ")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)

	fmt.Print("Enter language code (required): ")
	langCode, _ := reader.ReadString('\n')
	langCode = strings.TrimSpace(langCode)

	if langCode == "" {
		fmt.Println("Error: Language code is required")
		return
	}

	fmt.Println("Enter message content (type $% on a new line when finished):")
	var contentBuilder strings.Builder
	for {
		line, _ := reader.ReadString('\n')
		if strings.TrimSpace(line) == "$%" {
			break
		}
		contentBuilder.WriteString(line)
	}
	content := contentBuilder.String()

	if id == "" || content == "" {
		fmt.Println("Error: Message ID and content must be provided")
		return
	}

	// Preserve all line endings (CR, LF, CRLF)
	err := db.AddMessage(id, langCode, content)
	if err != nil {
		fmt.Printf("Error adding message: %v\n", err)
		return
	}

	fmt.Println("Message added successfully!")
}

func deleteMessage(db *messagedb.MessageDB) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter message ID to delete: ")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)

	if id == "" {
		fmt.Println("Error: Message ID must be provided")
		return
	}

	// First verify the message exists
	content, err := db.GetMessage("en", id) // Try English first
	if err != nil {
		fmt.Printf("Debug: Message lookup error: %v\n", err)
	} else {
		fmt.Printf("Debug: Found message with content: %s\n", content)
	}

	fmt.Printf("Are you sure you want to delete message '%s'? (y/n): ", id)
	confirm, _ := reader.ReadString('\n')
	confirm = strings.TrimSpace(confirm)

	if strings.ToLower(confirm) != "y" {
		fmt.Println("Deletion cancelled.")
		return
	}

	err = db.DeleteMessageByID(id)
	if err != nil {
		fmt.Printf("Error deleting message: %v\n", err)
		return
	}

	// Verify deletion
	content, err = db.GetMessage("en", id)
	if err != nil {
		fmt.Println("Message deleted successfully!")
	} else {
		fmt.Printf("Warning: Message still exists after deletion attempt. Content: %s\n", content)
		fmt.Println("This means the deletion operation did not affect any rows in the database.")
		fmt.Println("Please check that you're using the correct message ID.")
	}
}

func getMessageContent(db *messagedb.MessageDB) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter message ID: ")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)

	fmt.Print("Enter language code: ")
	langCode, _ := reader.ReadString('\n')
	langCode = strings.TrimSpace(langCode)

	if id == "" || langCode == "" {
		fmt.Println("Error: Both ID and language code must be provided")
		return
	}

	content, err := db.GetMessage(langCode, id)
	if err != nil {
		fmt.Printf("Error retrieving message: %v\n", err)
		return
	}

	fmt.Printf("Message content: %s\n", content)
}

func getSystemMessageContent(db *messagedb.MessageDB) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter message ID: ")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)

	if id == "" {
		fmt.Println("Error: Message ID must be provided")
		return
	}

	content, err := db.GetSystemMessage(id)
	if err != nil {
		fmt.Printf("Error retrieving message: %v\n", err)
		return
	}

	fmt.Printf("Message content (using system language): %s\n", content)
}

func interactiveMode(db *messagedb.MessageDB) {
	reader := bufio.NewReader(os.Stdin)

	// Get system language for display
	sysLang := messagedb.GetSystemLanguage()

	for {
		fmt.Println("\nMessage Database Admin Tool")
		fmt.Println("----------------------------")
		fmt.Printf("System Language: %s\n", sysLang)
		fmt.Println("1. Add message")
		fmt.Println("2. List messages")
		fmt.Println("3. Delete message")
		fmt.Println("4. Get message using system language")
		fmt.Println("0. Exit")
		fmt.Print("\nEnter your choice: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			addMessage(db)
		case "2":
			fmt.Print("Enter language code (leave empty to list all): ")
			langCode, _ := reader.ReadString('\n')
			langCode = strings.TrimSpace(langCode)
			if langCode == "" {
				listAllMessages(db)
			} else {
				listMessagesForLang(db, langCode)
			}
		case "3":
			deleteMessage(db)
		case "4":
			fmt.Print("Enter message ID: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)
			if id != "" {
				content, err := db.GetSystemMessage(id)
				if err != nil {
					fmt.Printf("Error: %v\n", err)
				} else {
					fmt.Printf("Message content: %s\n", content)
				}
			}
		case "0":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
