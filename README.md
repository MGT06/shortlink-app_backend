# ShortLink App - Backend

![Go Version](https://img.shields.io/badge/Go-1.27.3-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-Framework-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)
![Redis](https://img.shields.io/badge/Redis-DC382D?style=for-the-badge&logo=redis&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-000000?style=for-the-badge&logo=jsonwebtokens&logoColor=white)
![Swagger](https://img.shields.io/badge/Swagger-85EA2D?style=for-the-badge&logo=swagger&logoColor=black)
![License](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)

A robust URL shortener application built with Go, featuring user authentication, link management, and efficient redirection using Redis caching.

## Description

ShortLink App is a modern backend service that allows users to create shortened URLs, manage their links, and track redirections. Built with clean architecture principles, this application provides a fast and reliable way to shorten long URLs while maintaining user privacy and security through JWT-based authentication.

The application uses PostgreSQL for persistent data storage and Redis for caching, ensuring optimal performance and scalability. Each shortened link generates a unique slug that can be shared and tracked.

## Features

- **User Authentication**
  - User registration with secure password hashing
  - JWT-based authentication
  - Session management with Redis
  - Secure logout functionality

- **Link Management**
  - Create shortened URLs with custom or auto-generated slugs
  - View all user-created links
  - Delete links
  - Fast redirection to original URLs

- **Security**
  - Bcrypt password hashing
  - JWT token-based authorization
  - CORS middleware configuration
  - Protected API endpoints

- **Performance**
  - Redis caching for fast lookups
  - Efficient database queries with connection pooling
  - Optimized slug generation using UUID

- **Developer Experience**
  - Swagger API documentation
  - Database migrations support
  - Clean architecture structure
  - Comprehensive error handling

## Tech Stack

### Backend Framework
- **Go** (v1.27.3) - Programming language
- **Gin** - HTTP web framework

### Database
- **PostgreSQL** - Primary database for persistent storage
- **Redis** - In-memory caching and session management

### Libraries & Tools
- **pgx/v5** - PostgreSQL driver and connection pooling
- **go-redis/v9** - Redis client
- **golang-jwt/v5** - JWT implementation
- **bcrypt** - Password hashing
- **UUID** - Unique identifier generation
- **Swaggo** - API documentation generator
- **godotenv** - Environment variable management
- **golang-migrate** - Database migration tool

## Setup Instructions

### Prerequisites

Ensure you have the following installed on your system:
- Go 1.27.3 or higher
- PostgreSQL 12 or higher
- Redis 6 or higher
- golang-migrate CLI tool

### Installation Steps

1. **Clone the repository**
   ```bash
   git clone https://github.com/MGT06/shortlink-app_backend.git
   cd shortlink-app_backend
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Configure environment variables**
   
   Create a `.env` file in the root directory by copying the example:
   ```bash
   cp .env.example .env
   ```
   
   Update the `.env` file with your configuration values.

4. **Setup PostgreSQL**

5. **Run database migrations**
   ```bash
   make migrate-up
   ```

6. **Start Redis server**
   start with default configuration:
   ```bash
   redis-server
   ```

7. **Run the application**
   ```bash
   go run main.go
   ```
   
   The server will start at `http://localhost:5050`

### Development Commands

- **Create a new migration**
  ```bash
  make migrate-create NAME=migration_name
  ```

- **Run migrations**
  ```bash
  make migrate-up
  ```

- **Rollback migrations**
  ```bash
  make migrate-down
  ```

- **Generate Swagger documentation**
  ```bash
  swag init
  ```

## Endpoints

### API Documentation
- **Swagger UI**: `GET /docs/index.html`

### Authentication Endpoints

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| POST | `/api/register` | Register a new user | No |
| POST | `/api/login` | Login user and receive JWT token | No |
| POST | `/api/logout` | Logout user and invalidate token | Yes |

### Link Management Endpoints

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| POST | `/api/links` | Create a new short link | Yes |
| GET | `/api/links` | Get all links for authenticated user | Yes |
| DELETE | `/api/links/:id` | Delete a specific link | Yes |
| GET | `/:slug` | Redirect to original URL | No |

### Authentication

Protected endpoints require a Bearer token in the Authorization header:
Authorization: Bearer <jwt_token>


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Related Project

**Frontend Application**: [ShortLink App Frontend](https://github.com/MGT06/shortlink-app_frontend)

The frontend provides a user-friendly interface to interact with this API, featuring:
- Modern and responsive UI
- Link creation and management dashboard
- User authentication flows
---

