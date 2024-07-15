# URL Shortener

Welcome to the URL Shortener project made with Golang! This project uses SQLite for data storage and aims to provide a simple, readable, and user-friendly URL shortening service.

## Features

- **Readable and Friendly URLs**: Shortened URLs are easy to read and remember, using meaningful words.
- **Easy to Use**: Simple setup and usage, ideal for quick and efficient URL shortening.
- **SQLite Integration**: Utilizes SQLite for lightweight and hassle-free database management.
- **Safety Check**: Allows you to preview the destination URL before redirection, ensuring safety.

## Getting Started

### Prerequisites

Ensure you have the following installed on your machine:

- Go (version 1.21)

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/viniciusamc/url-shortener
    cd url-shortener
    ```

2. Build the project:
    ```sh
    go run ./cmd/*.go
    ```

3. The project will run using port 8123

### Configuration

The project uses an SQLite database for storing URLs. By default, it will create and use a database file named `database.sqlite3` in /internal/database.

## Usage

Once the server is running, you can use the following endpoints:

- **Create Short URL**: `POST /url`
  - Request Body:
    ```json
    {
        "link": "https://example.com",
        "duration": "2030-01-01 00:00:00"
    }
    ```
  - Response:
    ```json
    {
      "shortened_url": "http://localhost:8080/abacate-abacate"
    }
    ```

- **Redirect to Original URL**: `GET /{shortened_url}`
  - Redirects to the original URL associated with the shortened URL.

- **Preview Original URL**: `GET /{shortened_url}+`
  - Response:
    ```json
    {
      "link": "http://example.com"
    }
    ```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
