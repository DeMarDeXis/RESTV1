package sqlite_test

import (
	"testing"
	"url-shortener/internal/storage/sqlite"
)

func TestGetURL(t *testing.T) {
	storage, err := sqlite.New("C:/Users/User/Documents/GitHub/GoRestAPI/storage/storage.db")
	if err != nil {
		t.Fatalf("Failed to initialize storage: %v", err)
	}

	// Add data to the database
	alias := "ejemplo"
	url := "https://vk.com"
	_, err = storage.SaveURL(url, alias)
	if err != nil {
		t.Fatalf("Failed to save URL: %v", err)
	}

	// Test getting the URL
	resURL, err := storage.GetURL(alias)
	if err != nil {
		t.Fatalf("Failed to get URL: %v", err)
	}

	expectedURL := "https://vk.com"
	if resURL != expectedURL {
		t.Fatalf("Expected URL %s, got %s", expectedURL, resURL)
	}
}

func TestStorage_DeleteURL(t *testing.T) {
	storage, err := sqlite.New("C:/Users/User/Documents/GitHub/GoRestAPI/storage/storage.db")
	if err != nil {
		t.Fatalf("Failed to initialize storage: %v", err)
	}

	// Save a URL to delete later
	_, err = storage.SaveURL("vk.com", "VK")
	if err != nil {
		t.Fatalf("Failed to save URL: %v", err)
	}

	// Delete the URL
	err = storage.DeleteURL("testalias")
	if err != nil {
		t.Fatalf("Failed to delete URL: %v", err)
	}

	// Try to get the deleted URL to ensure it's gone
	_, err = storage.GetURL("testalias")
	if err == nil {
		t.Fatalf("Expected URL to be deleted, but it was found")
	}
}
