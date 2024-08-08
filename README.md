# Link Preview Package

The `linkpreview` package provides a simple and efficient way to fetch the title, description, and image from a given URL. It includes caching functionality to minimize repeated requests and allows for a customizable user agent.

## Installation

To install the package, run:

```bash
go get github.com/AmanuelCh/linkpreview
```

## Usage

Import the package in your Go file:

```go
import "github.com/AmanuelCh/linkpreview"
```

### Creating a LinkPreviewer

You can create a new LinkPreviewer instance with a custom user agent:

```go
lp := linkpreview.NewLinkPreviewer("MyCustomUserAgent/1.0")
```

### Fetching Link Previews

Use the `GetLinkPreview` function to fetch the metadata from a URL:

```go
url := "https://github.com/AmanuelCh/hahu"
title, description, image, err := lp.GetLinkPreview(url)
if err != nil {
    // Handle error
}

fmt.Printf("Title: %s\nDescription: %s\nImage: %s\n", title, description, image)
```

The `GetLinkPreview` function returns the following values:

- title: The title of the webpage.
- description: The description of the webpage.
- image: The URL of the main image on the webpage.
- err: An error, if any occurred during the process.

### Caching

The package caches the results for one hour. If the same URL is requested within this time frame, the cached data will be returned instead of making a new HTTP request.

### Custom User Agent

You can specify a custom user agent when creating the `LinkPreviewer` instance. This can help in identifying the requests made by your application.

### Example code

```go
package main

import (
    "fmt"
    "log"
    "github.com/AmanuelCh/linkpreview"
)

func main() {
    lp := linkpreview.NewLinkPreviewer("MyCustomUserAgent/1.0")
    url := "https://github.com/AmanuelCh/hahu"
    title, description, image, err := lp.GetLinkPreview(url)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Title: %s\nDescription: %s\nImage: %s\n", title, description, image)
}

```