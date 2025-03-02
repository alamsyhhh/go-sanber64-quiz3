# Book API Documentation

## 📌 Overview

This is the backend API for managing books. The API allows users to create, update, delete, and retrieve books.

## 📖 API Documentation

The API documentation can be accessed at:

🔗 [API Documentation](https://go-sanber64-quiz3.onrender.com/api-docs/index.html)

## 🔑 Authentication

To use the API, you need to register or log in with an existing account.

### Test Account:

```json
{
  "username": "test 3",
  "password": "test12345"
}
```

You can also register a new account through the API if needed.

## 🚀 Getting Started

### 1. Clone the Repository

```sh
git clone <repository_url>
cd <project_directory>
```

### 2. Install Dependencies

```sh
go mod tidy
```

### 3. Set Environment Variables

Ensure you have a `.env` file configured with the required environment variables, including `CLOUDINARY_URL`.

### 4. Run the Server

```sh
go run main.go
```

### 5. Access the API

Once the server is running, open your browser and go to:

```
http://localhost:8080/api-docs/index.html#/
```

## 🛠 Technologies Used

- Golang
- Gin Framework
- Goqu
- Cloudinary (for image storage)
- PostgreSQL (Database)
