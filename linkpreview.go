package linkpreview

import (
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// CacheEntry holds the cached data for a URL.
type CacheEntry struct {
	Title       string
	Description string
	Image       string
	Expires     time.Time
}

// LinkPreviewer struct holds the cache and a mutex for thread safety.
type LinkPreviewer struct {
	cache map[string]CacheEntry
	mu    sync.Mutex
	agent string
}

// NewLinkPreviewer creates a new LinkPreviewer with the specified user agent.
func NewLinkPreviewer(userAgent string) *LinkPreviewer {
	return &LinkPreviewer{
		cache: make(map[string]CacheEntry),
		agent: userAgent,
	}
}

// GetLinkPreview fetches the title, description, and image from the given URL.
func (lp *LinkPreviewer) GetLinkPreview(url string) (title, description, image string, err error) {
	lp.mu.Lock()
	defer lp.mu.Unlock()

	// Check cache
	if entry, found := lp.cache[url]; found && time.Now().Before(entry.Expires) {
		return entry.Title, entry.Description, entry.Image, nil
	}

	// Fetch from URL
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", lp.agent)

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}

	title = strings.TrimSpace(doc.Find("meta[property='og:title']").AttrOr("content", ""))
	if title == "" {
		title = strings.TrimSpace(doc.Find("title").Text())
	}

	description = strings.TrimSpace(doc.Find("meta[property='og:description']").AttrOr("content", ""))
	image = strings.TrimSpace(doc.Find("meta[property='og:image']").AttrOr("content", ""))

	// Cache the result for 1 hour
	lp.cache[url] = CacheEntry{
		Title:       title,
		Description: description,
		Image:       image,
		Expires:     time.Now().Add(1 * time.Hour),
	}

	return
}
