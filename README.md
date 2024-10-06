# go-url-shortener
# Go URL Shortener

This is a simple URL shortener built using the Go programming language. It takes long URLs and generates a shortened version. You can use the shortened URL to redirect to the original long URL.

## Features

- **Create Short URLs**: Takes a long URL and creates a shorter, unique version of it.
- **Redirect to Original URL**: The shortened URL will redirect you back to the original long URL.
- **In-Memory Storage**: URLs are stored in memory using a map for simplicity.

## How It Works

1. **Shorten a URL**:  
   Send a `POST` request with the long URL, and get a short URL in response.

2. **Redirect to Original URL**:  
   Use the short URL, and it will redirect you to the original long URL.

## Endpoints

### 1. Root Endpoint

- **URL**: `/`
- **Method**: `GET`
- **Response**: `"Hello World"`

### 2. Create Short URL

- **URL**: `/shorten`
- **Method**: `POST`
- **Request Body (JSON)**:
  ```json
  {
    "url": "https://example.com/your-long-url"
  }
