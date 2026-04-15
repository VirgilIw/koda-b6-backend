# koda-b6-backend
REST API backend for a Coffee Shop application, built using **Go** with the **Gin** framework.

---

## Tech Stack
- **Language**: Go 1.25.3
- **Framework**: Gin
- **Database**: PostgreSQL (via `pgx/v5`)
- **Cache**: Redis (`go-redis/v9`)
- **Auth**: JWT (`golang-jwt/jwt/v5`)
- **Password Hashing**: Argon2 (`matthewhartstonge/argon2`)
- **API Docs**: Swagger (`swaggo/swag`)
- **Containerization**: Docker (multi-stage build)
- **Hot Reload (dev)**: Air

---

## Project Structure
```
koda-b6-backend/
├── cmd/
│   └── main.go               # Entry point
├── internal/                 # Business logic (handler, service, repository)
├── migrations/               # Database migration files
├── docs/                     # Swagger generated docs
├── images/                   # Static assets
├── .air.toml                 # Air hot reload config
├── .env.example              # Environment variable template
├── Dockerfile                # Multi-stage Docker build
└── go.mod
```

---

## Getting Started

### 1. Clone the repository
```bash
git clone https://github.com/VirgilIw/koda-b6-backend.git
cd koda-b6-backend
```

### 2. Install dependencies
```bash
go mod tidy
```

### 3. Setup environment variables
```bash
cp .env.example .env
```

Fill in the `.env` file with the appropriate values:
```env
PORT=
DB_HOST=
DB_PORT=
DB_USERNAME=
DB_PASSWORD=
DB_NAME=
SECRET_KEY=
REDIS_HOST=
REDIS_PORT=
REDIS_PASSWORD=
FRONTEND_URL=
APP_URL=
```

| Variable | Description |
|---|---|
| `PORT` | Port the server runs on |
| `DB_HOST` | PostgreSQL host |
| `DB_PORT` | PostgreSQL port (default: `5432`) |
| `DB_USERNAME` | PostgreSQL username |
| `DB_PASSWORD` | PostgreSQL password |
| `DB_NAME` | Database name |
| `SECRET_KEY` | Secret key for JWT |
| `REDIS_HOST` | Redis host |
| `REDIS_PORT` | Redis port (default: `6379`) |
| `REDIS_PASSWORD` | Redis password (leave empty if not set) |
| `FRONTEND_URL` | Frontend URL for CORS |
| `APP_URL` | Base URL of this application |

### 4. Run database migrations
```bash
# Make sure the database has been created, then run the migration
go run cmd/main.go migrate
```

### 5. Run the server
```bash
# Development (with hot reload using Air)
air

# Or without hot reload
go run cmd/main.go
```

Server runs at `http://localhost:{PORT}`

---

## Docker

### Build & Run
```bash
# Build image
docker build -t koda-b6-backend .

# Run container
docker run -p 8888:8888 --env-file .env koda-b6-backend
```

The application will run on port `8888`.

### Multi-stage Build
The Dockerfile uses two stages:
- **Stage 1 (build)**: Compiles the Go binary using `golang:1.25.3-alpine`
- **Stage 2 (run)**: Runs the binary on a lighter `alpine:latest` image

---

## API Documentation

Swagger UI is available once the server is running:
```
http://localhost:{PORT}/swagger/index.html
```

To regenerate docs after making changes:
```bash
swag init -g cmd/main.go
```

---

## License

MIT License

Copyright (c) 2026 VirgilIw

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
