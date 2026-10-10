# Go API Client — Website Data Demo

A small, working Go command-line application that retrieves JSON data from a public REST API.

## What it does
- Retrieves and displays 10 posts.
- Retrieves one post by ID.
- Filters posts by user ID.
- Retrieves user profile information by ID.
- Handles HTTP errors, invalid JSON, and network timeouts.
- Uses only Go's standard library; no API key or third-party packages are required.

## Data source
JSONPlaceholder: https://jsonplaceholder.typicode.com

This is a public demo API intended for learning and testing. Its posts and users are sample data, not a live news feed or a real production website's private data.

## Requirements
- Go 1.20+ (the code uses standard library APIs and `any`).
- Internet connection.

## Run on Windows
1. Extract the ZIP archive.
2. Open PowerShell in the extracted folder.
3. Run:

```powershell
go run main.go
```

To build a Windows executable:

```powershell
go build -o go-api-client.exe main.go
```

Then run:

```powershell
.\go-api-client.exe
```

## Menu
1. Show the 10 first posts returned by the API.
2. Find a post by ID.
3. Show posts by user ID.
4. Show user information by ID.
0. Exit.


