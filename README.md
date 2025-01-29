# Short URL (go)

This is a simple URL shortening service built with Go and the Fiber framework (v3). The service allows users to shorten long URLs, track the number of clicks, and manage URL expiration. This project uses SQLite as the database for storing URLs and related metadata.

## Features
- **Shorten URLs**: Convert a long URL into a short one.
- **Track Clicks**: Count and display the number of clicks for each shortened URL.
- **Expiry Date**: Set an expiration date for a short URL. Once expired, the URL will no longer be accessible.
- **API for Stats**: Retrieve statistics for a shortened URL, including the number of clicks and expiry date.
- **URL Redirection**: Redirect the user to the original long URL when they visit the shortened URL.

## Requirements
- Go 1.23.4 
- SQLite3

## Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/FeiZGe/ShortURL-go.git
   cd short-url-service
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run the server:
   ```bash
   go run main.go
   ```
The server will start at `http://localhost:3000`.

## API Endpoints