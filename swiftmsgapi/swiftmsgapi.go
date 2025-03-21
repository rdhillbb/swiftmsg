package swiftmsgapi 

import (
	"database/sql"
	"errors"
	"log"
	"time"
        "strings"
	_ "github.com/mattn/go-sqlite3"
)

// MessageDB handles all database operations for the message storage system
type MessageDB struct {
	db *sql.DB
}

// Message represents a stored message
type Message struct {
	ID        string
	LangCode  string
	Content   string
	CreatedAt int64
}

// NewMessageDB creates and initializes a new message database
func NewMessageDB(dbPath string) (*MessageDB, error) {
	if dbPath == "" {
		dbPath = ":memory:"
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Set pragmas for faster performance
	_, err = db.Exec("PRAGMA journal_mode = WAL")
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("PRAGMA synchronous = NORMAL")
	if err != nil {
		return nil, err
	}

	msgDB := &MessageDB{db: db}
	if err := msgDB.initDB(); err != nil {
		return nil, err
	}

	return msgDB, nil
}

// initDB sets up the database schema
func (m *MessageDB) initDB() error {
	// Create messages table
	_, err := m.db.Exec(`
		CREATE TABLE IF NOT EXISTS messages (
			id TEXT NOT NULL,
			lang_code TEXT NOT NULL,
			content TEXT NOT NULL,
			created_at INTEGER DEFAULT (strftime('%s', 'now')),
			PRIMARY KEY (id, lang_code)
		)
	`)
	if err != nil {
		return err
	}

	// Create index for fast retrieval
	_, err = m.db.Exec("CREATE INDEX IF NOT EXISTS idx_lang_id ON messages(lang_code, id)")
	return err
}

// AddMessage adds or updates a message in the database
func (m *MessageDB) AddMessage(id, langCode, content string) error {
	if id == "" || langCode == "" {
		return errors.New("id and langCode cannot be empty")
	}

	stmt, err := m.db.Prepare("INSERT OR REPLACE INTO messages (id, lang_code, content, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, langCode, content, time.Now().Unix())
	return err
}

// GetMessage retrieves a message by language code and message ID
func (m *MessageDB) GetMessage(langCode, id string) (string, error) {
	if id == "" {
		return "", errors.New("id cannot be empty")
	}
	
	// Convert language name/variant to standard code and default to English if empty
	standardCode := GetLanguageCode(langCode)
	
	var content string
	err := m.db.QueryRow(
		"SELECT content FROM messages WHERE lang_code = ? AND id = ? LIMIT 1",
		standardCode, id,
	).Scan(&content)

	if err == sql.ErrNoRows {
		// If message not found in requested language, try English
		if standardCode != "en" {
			err = m.db.QueryRow(
				"SELECT content FROM messages WHERE lang_code = ? AND id = ? LIMIT 1",
				"en", id,
			).Scan(&content)
			
			if err == nil {
				return content, nil
			}
		}
		return "", errors.New("message not found")
	}

	return content, err
}

// GetSystemMessage retrieves a message using the system's language settings
func (m *MessageDB) GetSystemMessage(id string) (string, error) {
	if id == "" {
		return "", errors.New("id cannot be empty")
	}

	// Get system language from environment variables
	langCode := getSystemLanguage()
	
	return m.GetMessage(langCode, id)
}

// getSystemLanguage determines the system language from environment variables
func getSystemLanguage() string {
	// Check environment variables in order of precedence
	for _, envVar := range []string{"LC_ALL", "LC_MESSAGES", "LANG", "LANGUAGE"} {
		locale := os.Getenv(envVar)
		if locale != "" {
			// Extract language code (e.g., "en" from "en_US.UTF-8")
			if strings.Contains(locale, "_") {
				return strings.Split(locale, "_")[0]
			}
			if strings.Contains(locale, ".") {
				return strings.Split(locale, ".")[0]
			}
			return locale
		}
	}
	
	// Default to English if no locale is set
	return "en"
}

// DeleteMessage removes a message from the database
func (m *MessageDB) DeleteMessage(langCode, id string) error {
	if id == "" || langCode == "" {
		return errors.New("id and langCode cannot be empty")
	}

	stmt, err := m.db.Prepare("DELETE FROM messages WHERE lang_code = ? AND id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(langCode, id)
	return err
}

// ListMessages returns all messages for a specific language
func (m *MessageDB) ListMessages(langCode string) ([]Message, error) {
	if langCode == "" {
		return nil, errors.New("langCode cannot be empty")
	}

	rows, err := m.db.Query(
		"SELECT id, lang_code, content, created_at FROM messages WHERE lang_code = ?",
		langCode,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var msg Message
		if err := rows.Scan(&msg.ID, &msg.LangCode, &msg.Content, &msg.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	return messages, nil
}

// ListAllMessages returns all messages in the database
func (m *MessageDB) ListAllMessages() ([]Message, error) {
	rows, err := m.db.Query(
		"SELECT id, lang_code, content, created_at FROM messages",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var msg Message
		if err := rows.Scan(&msg.ID, &msg.LangCode, &msg.Content, &msg.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	return messages, nil
}

// Close closes the database connection
func (m *MessageDB) Close() error {
	return m.db.Close()
}
