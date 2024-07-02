package path

import (
	"os/user"
	"testing"
)

func TestParse(t *testing.T) {
	currentUser, _ := user.Current()
	username := currentUser.Name
	// 1. user@host:path
	result1 := Parse("john@example.com:/path/to/file.txt")
	if result1.User != "john" {
		t.Errorf("Expected user: john, got: %s", result1.User)
	}
	if result1.Host != "example.com" {
		t.Errorf("Expected host: example.com, got: %s", result1.Host)
	}
	if result1.FilePath != "/path/to/file.txt" {
		t.Errorf("Expected file path: /path/to/file.txt, got: %s", result1.FilePath)
	}

	// 2. user@path
	result2 := Parse("jane@localhost:/path/to/file.txt")
	if result2.User != "jane" {
		t.Errorf("Expected user: jane, got: %s", result2.User)
	}
	if result2.Host != "localhost" {
		t.Errorf("Expected host: localhost, got: %s", result2.Host)
	}
	if result2.FilePath != "/path/to/file.txt" {
		t.Errorf("Expected file path: /path/to/file.txt, got: %s", result2.FilePath)
	}

	// 3. host:path
	result3 := Parse("example.com:/path/to/file.txt")
	if result3.User != username {
		t.Errorf("Expected user: %s, got: %s", username, result3.User)
	}
	if result3.Host != "example.com" {
		t.Errorf("Expected host: example.com, got: %s", result3.Host)
	}
	if result3.FilePath != "/path/to/file.txt" {
		t.Errorf("Expected file path: /path/to/file.txt, got: %s", result3.FilePath)
	}

	// 4. path
	result4 := Parse("/path/to/file.txt")
	if result4.User != username {
		t.Errorf("Expected user: %s, got: %s", username, result4.User)
	}
	if result4.Host != "localhost" {
		t.Errorf("Expected host: localhost, got: %s", result4.Host)
	}
	if result4.FilePath != "/path/to/file.txt" {
		t.Errorf("Expected file path: /path/to/file.txt, got: %s", result4.FilePath)
	}

	// 5. Empty path
	result5 := Parse("")
	if result5.User != username {
		t.Errorf("Expected user: %s, got: %s", username, result5.User)
	}
	if result5.Host != "localhost" {
		t.Errorf("Expected host: localhost, got: %s", result5.Host)
	}
	if result5.FilePath != "" {
		t.Errorf("Expected file path: '', got: %s", result5.FilePath)
	}

	// 6. Only user
	result6 := Parse("john")
	if result6.User != username {
		t.Errorf("Expected user: %s, got: %s", username, result6.User)
	}
	if result6.Host != "localhost" {
		t.Errorf("Expected host: localhost, got: %s", result6.Host)
	}
	if result6.FilePath != "john" {
		t.Errorf("Expected file path: john, got: %s", result6.FilePath)
	}

	// 7. Only host
	result7 := Parse("example.com:")
	if result7.User != username {
		t.Errorf("Expected user: %s, got: %s", username, result7.User)
	}
	if result7.Host != "example.com" {
		t.Errorf("Expected host: example.com, got: %s", result7.Host)
	}
	if result7.FilePath != "" {
		t.Errorf("Expected file path: '', got: %s", result7.FilePath)
	}

	// 8. Only @
	result8 := Parse("@localhost:/path/to/file.txt")
	if result8.User != username {
		t.Errorf("Expected user: %s, got: %s", username, result8.User)
	}
	if result8.Host != "localhost" {
		t.Errorf("Expected host: localhost, got: %s", result8.Host)
	}
	if result8.FilePath != "/path/to/file.txt" {
		t.Errorf("Expected file path: /path/to/file.txt, got: %s", result8.FilePath)
	}

	// 9. Only :
	result9 := Parse("example.com:")
	if result9.User != username {
		t.Errorf("Expected user: %s, got: %s", username, result9.User)
	}
	if result9.Host != "example.com" {
		t.Errorf("Expected host: example.com, got: %s", result9.Host)
	}
	if result9.FilePath != "" {
		t.Errorf("Expected file path: '', got: %s", result9.FilePath)
	}

	// 10. Invalid input
	result10 := Parse("invalid_input")
	if result10.User != username {
		t.Errorf("Expected user: %s, got: %s", username, result10.User)
	}
	if result10.Host != "localhost" {
		t.Errorf("Expected host: localhost, got: %s", result10.Host)
	}
	if result10.FilePath != "invalid_input" {
		t.Errorf("Expected file path: invalid_input, got: %s", result10.FilePath)
	}
}
