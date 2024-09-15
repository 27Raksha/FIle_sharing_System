# **21BLC1564 - File Sharing & Management System**

## Introduction

**File Sharing & Management System** is a backend project built using Go that allows users to securely upload, manage, and share files. This system provides user authentication, file metadata management, and caching for efficient file retrieval. Redis is used for caching, and PostgreSQL is used for persistent data storage. Concurrency is handled via Go's goroutines, making the system scalable and efficient for handling large file uploads.

## Features

- **User Authentication & Authorization**: Secure JWT-based authentication system allowing users to manage their own files.
- **File Upload & Metadata Management**: Upload files locally or to S3 with metadata stored in PostgreSQL.
- **File Retrieval & Sharing**: Retrieve file metadata and share files via a public URL.
- **File Search**: Search for files based on metadata (e.g., name, upload date, or file type).
- **Caching Layer**: Uses Redis to cache file metadata, reducing database load.
- **Background Worker for Expired Files**: A worker that periodically deletes expired files from storage and their metadata from the database.
- **Concurrency**: Efficiently processes large file uploads using goroutines.

## Technologies Used

- **Go**: For backend API development.
- **PostgreSQL**: To store file and user metadata.
- **Redis**: For caching file metadata.
- **Docker**: For containerization and environment consistency.
- **Gin**: A web framework for routing HTTP requests.
- **GORM**: For ORM-based database interactions with PostgreSQL.
- **JWT**: For secure user authentication.

## API Endpoints

### Authentication

- `POST /auth/register` - Register a new user.
- `POST /auth/login` - Log in and receive a JWT token.

### File Management

- `POST /api/upload` - Upload a file.
- `GET /api/files` - List all uploaded files.
- `GET /api/share/:file_id` - Get a public shareable link for a file.
- `GET /api/search/files?name=:name` - Search for files by name, upload date, or file type.

## Running the Application

### Prerequisites

- **Go 1.20+**
- **PostgreSQL**
- **Redis**
- **Docker** (optional, for containerized setup)

# Steps for Setting up the File Sharing & Management System

## 1. Clone the Repository
First, clone the GitHub repository to your local machine:
```bash
git clone https://github.com/27Raksha/21BLC1564_Backend.git
cd 21BLC1564_Backend
```
## 2. Set up Enviroment Variable
create a .env file in the root directory and configure your enviroment variables:
```bash
DATABASE_URL=postgresql://username:password@localhost/database_name
REDIS_ADDR=redis_address
REDIS_PASSWORD=your_redis_password
REDIS_DB=0
```
## 3. Install Dependencies:
```bash
go mod download
```
## 4. Run the Application Locally:
```bash
go run main.go
```
## 5. Run with Docker(Optional):
If you want to run the application using Docker, follow these steps:
# Build and Run the Containers
To build and run the app using Docker:
```bash
docker-compose up --build
```
## 6. Running Test:
To ensure everything is working as expected, run the test suite for your project:
```bash
go test ./... -v
```
