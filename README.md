# Koda B6 — Coffee Shop Backend

A RESTful API backend for a coffee shop application, built with **Go** and the **Gin** framework. This project follows a layered architecture with PostgreSQL, Redis, JWT authentication, and Swagger API documentation.

---

## Tech Stack

| Technology | Description |
|---|---|
| **Go 1.25** | Primary language |
| **Gin** | HTTP web framework |
| **PostgreSQL (pgx/v5)** | Main database |
| **Redis** | Caching & session management |
| **JWT (golang-jwt/v5)** | Token-based authentication |
| **Argon2** | Password hashing |
| **Swagger (swaggo)** | API documentation |
| **Docker** | Containerization |
| **Air** | Hot reload for development |

---

## Project Structure

```
koda-b6-backend/
├── cmd/                  # Application entry point (main.go)
├── internal/             # Core application logic (handler, service, repository)
├── migrations/           # Database migration files
├── docs/                 # Auto-generated Swagger documentation
├── images/               # Image assets
├── .github/workflows/    # CI/CD pipeline
├── .air.toml             # Air configuration (hot reload)
├── .env.example          # Environment variable template
├── Dockerfile            # Multi-stage Docker build configuration
├── go.mod
└── go.sum
```

---

## Environment Configuration

Copy the `.env.example` file to `.env` and fill in the appropriate values:

```bash
cp .env.example .env
```

```env
PORT=                  # Application port (e.g. 8888)

DB_HOST=               # PostgreSQL host
DB_PORT=               # Database port (e.g. 5432)
DB_USERNAME=           # Database username
DB_PASSWORD=           # Database password
DB_NAME=               # Database name

SECRET_KEY=            # Secret key for JWT signing

REDIS_HOST=            # Redis host
REDIS_PORT=            # Redis port (e.g. 6379)
REDIS_PASSWORD=        # Redis password (leave empty if none)

FRONTEND_URL=          # Frontend URL (used for CORS)
APP_URL=               # Backend application URL
```

---

## Running the Application

### Option 1: Local (Development)

Make sure you have [Go](https://go.dev/dl/) and [Air](https://github.com/air-verse/air) installed.

```bash
# Install dependencies
go mod tidy

# Run with hot reload
air

# Or run directly
go run cmd/main.go
```

### Option 2: Docker

```bash
# Build the image
docker build -t koda-b6-backend .

# Run the container
docker run -p 8888:8888 --env-file .env koda-b6-backend
```

The application will run on port **8888**.

---

## 📖 API Documentation

Swagger documentation is available once the application is running. Access it at:

```
http://localhost:8888/swagger/index.html
```

To regenerate Swagger docs:

```bash
swag init -g cmd/main.go
```

---

## Database Migration

Migration files are located in the `migrations/` directory. Run the migrations before starting the application for the first time.

---

## Contributing

1. Fork this repository
2. Create a new feature branch: `git checkout -b feature/your-feature-name`
3. Commit your changes: `git commit -m 'feat: add new feature'`
4. Push to the branch: `git push origin feature/your-feature-name`
5. Open a Pull Request

---

## License

```
MIT License

Copyright (c) 2026 VirgilIw

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
