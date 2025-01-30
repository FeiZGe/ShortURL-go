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
### 1. POST /api/shorten
   
   *Description*: Shorten a long URL.
   - **Request Body**:
      ```json
      {
         "long_url": "https://www.example.com",
         "expiry_days": 5
      }
      ```
   - **Response**:
     ```json
      {
         "short_url": "http://localhost:3000/abCDef"
      }
      ```

### 2. GET /:short
   *Description*: Redirect to the original long URL using the shortened URL.
   - **Example**: http://localhost:3000/abCDef
   - **Response**: Redirect to the original long URL.

### 3. GET /stats/:short
   *Description*: Get statistics for a shortened URL.
   - **Example**: http://localhost:3000/stats/abCDef
   - **Response**:
      ```json
      {
          "click_count": 3,
          "expiry_date": "2025-02-06T13:37:15.803299+07:00",
          "long_url": "https://example.com",
          "short_url": "http://localhost:3000/abCDef"
      }
      ```

## Database Schema
The service uses an SQLite database with the following table structure:

### **URL table**:
- `ID`: Primary key, auto-incremented.
- `ShortURL`: The generated short URL (unique).
- `LongURL`: The original long URL.
- `ClickCount`: The number of clicks on the shortened URL.
- `ExpiryDate`: The date and time when the short URL expires.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.