def add_message(self, message_id: str, lang_input: str, content: str) -> None:
        """Add a new message to the database"""
        try:
            # Convert language name to standard code
            lang_code = get_language_code(lang_input)
            
            # Show the language code being used
            print(f"Using language code: {lang_code}")
            
            self.db.add_message(message_id, lang_code, content)
            print(f"Message '{message_id}' in language '{lang_code}' added successfully.")
        except Exception as e:
            print(f"Error adding message: {e}")
            
    def delete_message(self, lang_input: str, message_id: str) -> None:
        """Delete a message from the database"""
        try:
            # Convert language name to standard code
            lang_code = get_language_code(lang_input)
            
            # Show the language code being used
            print(f"Using language code: {lang_code}")
            
            result = self.db.delete_message(lang_code, message_id)
            if result:
                print(f"Message '{message_id}' in language '{lang_code}' deleted successfully.")
            else:
                print(f"Message '{message_id}' in language '{lang_code}' not found.")
        except Exception as e:
            print(f"Error deleting message: {e}")
            
    def list_messages(self, lang_input: str = None) -> None:
        """List messages in the database"""
        try:
            # Convert language name to standard code if provided
            lang_code = get_language_code(lang_input) if lang_input else None
            #!/usr/bin/env python3
"""
Message Database Admin Tool

This tool provides a command-line interface for managing the message database.
It allows adding, retrieving, deleting, and listing messages.
"""

import argparse
import os
import sys
import time
from typing import List, Dict, Any

# Import the MessageDB class
try:
    from message_db import MessageDB
    from language_util import get_language_code, LANGUAGE_MAP
except ImportError:
    # If imported from the same directory
    sys.path.append(os.path.dirname(os.path.abspath(__file__)))
    from message_db import MessageDB
    from language_util import get_language_code, LANGUAGE_MAP


class MessageDBAdmin:
    """Admin tool for managing the message database"""
    
    def __init__(self, db_path: str):
        """Initialize the admin tool with a database path"""
        self.db_path = db_path
        self.db = MessageDB(db_path)
        
    def add_message(self, message_id: str, lang_code: str, content: str) -> None:
        """Add a new message to the database"""
        try:
            self.db.add_message(message_id, lang_code, content)
            print(f"Message '{message_id}' in language '{lang_code}' added successfully.")
        except Exception as e:
            print(f"Error adding message: {e}")
            
    def get_message(self, lang_code: str, message_id: str) -> None:
        """Retrieve and display a message"""
        try:
            message = self.db.get_message(lang_code, message_id)
            if message:
                print(f"\nMessage content: {message}")
            else:
                print(f"Message '{message_id}' in language '{lang_code}' not found.")
        except Exception as e:
            print(f"Error retrieving message: {e}")
            
    def delete_message(self, lang_code: str, message_id: str) -> None:
        """Delete a message from the database"""
        try:
            result = self.db.delete_message(lang_code, message_id)
            if result:
                print(f"Message '{message_id}' in language '{lang_code}' deleted successfully.")
            else:
                print(f"Message '{message_id}' in language '{lang_code}' not found.")
        except Exception as e:
            print(f"Error deleting message: {e}")
            
    def list_messages(self, lang_code: str = None) -> None:
        """List messages in the database"""
        try:
            messages = self.db.list_messages(lang_code)
            
            if not messages:
                if lang_code:
                    print(f"No messages found for language '{lang_code}'.")
                else:
                    print("No messages found in the database.")
                return
            
            # Format and display messages
            if lang_code:
                print(f"\nMessages for language '{lang_code}':")
            else:
                print("\nAll messages in the database:")
                
            print("-" * 60)
            print(f"{'ID':<20} {'Language':<10} {'Content':<30}")
            print("-" * 60)
            
            for msg in messages:
                print(f"{msg['id']:<20} {msg['lang_code']:<10} {msg['content']:<30}")
                
            print(f"\nTotal: {len(messages)} message(s)")
            
        except Exception as e:
            print(f"Error listing messages: {e}")
            
    def show_stats(self) -> None:
        """Display database statistics"""
        try:
            stats = self.db.get_stats()
            
            print("\nDatabase Statistics")
            print("-" * 30)
            print(f"Database path: {stats['db_path']}")
            print(f"Total messages: {stats['total_messages']}")
            
            if stats['languages']:
                print("\nMessages per language:")
                for lang, count in stats['languages'].items():
                    print(f"  {lang}: {count}")
            
        except Exception as e:
            print(f"Error retrieving statistics: {e}")
            
    def get_system_message(self, message_id: str) -> None:
        """Retrieve and display a message using system language"""
        try:
            message = self.db.get_system_message(message_id)
            if message:
                print(f"\nMessage content (using system language): {message}")
            else:
                print(f"Message '{message_id}' not found in system language or English fallback.")
        except Exception as e:
            print(f"Error retrieving message: {e}")

    def interactive_mode(self) -> None:
        """Run the admin tool in interactive mode"""
        print("\nMessage Database Admin Tool")
        print("=========================")
        
        # Get system language for display
        sys_lang = self.db.get_system_language()
        
        while True:
            print("\nOptions:")
            print(f"  System Language: {sys_lang}")
            print("  1. Add message")
            print("  2. List messages")
            print("  3. Delete message") 
            print("  4. Get message using system language")
            print("  0. Exit")
            
            choice = input("\nEnter your choice (0-3): ").strip()
            
            if choice == "1":
                message_id = input("Enter message ID: ").strip()
                lang_code = input("Enter language code (required): ").strip()
                
                if not lang_code:
                    print("Error: Language code is required")
                    continue
                
                print("Enter message content (type $% on a new line when finished):")
                content_lines = []
                while True:
                    line = input()
                    if line.strip() == "$%":
                        break
                    content_lines.append(line)
                
                # Preserve all line endings (CR, LF, CRLF)
                content = '\n'.join(content_lines)
                
                self.add_message(message_id, lang_code, content)
                
            elif choice == "2":
                lang_code = input("Enter language code (leave empty to list all): ").strip()
                self.list_messages(lang_code)
                
            elif choice == "3":
                lang_code = input("Enter language code: ").strip()
                message_id = input("Enter message ID: ").strip()
                confirm = input(f"Are you sure you want to delete '{message_id}' in language '{lang_code}'? (y/n): ").strip().lower()
                if confirm == 'y':
                    self.delete_message(lang_code, message_id)
                
            elif choice == "0":
                print("Exiting...")
                break
                
            else:
                print("Invalid choice. Please try again.")
    
    def close(self) -> None:
        """Close the database connection"""
        self.db.close()


def main():
    """Main entry point for the admin tool"""
    parser = argparse.ArgumentParser(description="Message Database Admin Tool")
    
    parser.add_argument("--db", default="messages.db", help="Path to the database file")
    parser.add_argument("--interactive", "-i", action="store_true", help="Run in interactive mode")
    
    subparsers = parser.add_subparsers(dest="command", help="Command to execute")
    
    # Add command
    add_parser = subparsers.add_parser("add", help="Add a new message")
    add_parser.add_argument("--id", required=True, help="Message ID")
    add_parser.add_argument("--lang", required=True, help="Language code")
    add_parser.add_argument("--content", required=True, help="Message content")
    
    # Get command
    get_parser = subparsers.add_parser("get", help="Get a message")
    get_parser.add_argument("--lang", required=True, help="Language code")
    get_parser.add_argument("--id", required=True, help="Message ID")
    
    # Delete command
    delete_parser = subparsers.add_parser("delete", help="Delete a message")
    delete_parser.add_argument("--lang", required=True, help="Language code")
    delete_parser.add_argument("--id", required=True, help="Message ID")
    
    # List command
    list_parser = subparsers.add_parser("list", help="List messages")
    list_parser.add_argument("--lang", help="Filter by language code")
    
    # Stats command
    subparsers.add_parser("stats", help="Show database statistics")
    
    args = parser.parse_args()
    
    # Create admin instance
    admin = MessageDBAdmin(args.db)
    
    try:
        # Interactive mode
        if args.interactive or not args.command:
            admin.interactive_mode()
        # Command mode
        elif args.command == "add":
            admin.add_message(args.id, args.lang, args.content)
        elif args.command == "get":
            admin.get_message(args.lang, args.id)
        elif args.command == "delete":
            admin.delete_message(args.lang, args.id)
        elif args.command == "list":
            admin.list_messages(args.lang)
        elif args.command == "stats":
            admin.show_stats()
    finally:
        admin.close()


if __name__ == "__main__":
    main()
