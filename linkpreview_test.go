package linkpreview

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestGetLinkPreview tests the GetLinkPreview function.
func TestGetLinkPreview(t *testing.T) {
	// Mock server to simulate a response with metadata
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, `
			<html>
				<head>
					<meta property="og:title" content="Mock Title"/>
					<meta property="og:description" content="Mock Description"/>
					<meta property="og:image" content="https://example.com/image.jpg"/>
				</head>
				<body>
					<title>Fallback Title</title>
				</body>
			</html>
		`)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	lp := NewLinkPreviewer("TestUserAgent/1.0")

	// Test successful retrieval of link preview
	title, description, image, err := lp.GetLinkPreview(server.URL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if title != "Mock Title" {
		t.Errorf("Expected title 'Mock Title', got '%s'", title)
	}
	if description != "Mock Description" {
		t.Errorf("Expected description 'Mock Description', got '%s'", description)
	}
	if image != "https://example.com/image.jpg" {
		t.Errorf("Expected image 'https://example.com/image.jpg', got '%s'", image)
	}

	// Test caching mechanism
	title2, description2, image2, err := lp.GetLinkPreview(server.URL)
	if err != nil {
		t.Fatalf("Expected no error on cache hit, got %v", err)
	}
	if title2 != title || description2 != description || image2 != image {
		t.Errorf("Expected cached values to match original, got title '%s', description '%s', image '%s'", title2, description2, image2)
	}
}

// TestGetLinkPreviewError tests the GetLinkPreview function for error handling.
func TestGetLinkPreviewError(t *testing.T) {
	lp := NewLinkPreviewer("TestUserAgent/1.0")

	// Test with an invalid URL
	_, _, _, err := lp.GetLinkPreview("http://invalid.url")
	if err == nil {
		t.Fatal("Expected an error for invalid URL, got none")
	}
}
