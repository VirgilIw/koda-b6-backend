# koda-b6-backend

REST API backend untuk aplikasi Coffee Shop, dibangun menggunakan **Go** dengan framework **Gin**.

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

### 1. Clone repository

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

Isi file `.env` dengan value yang sesuai:

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

| Variable | Deskripsi |
|---|---|
| `PORT` | Port server berjalan |
| `DB_HOST` | Host PostgreSQL |
| `DB_PORT` | Port PostgreSQL (default: `5432`) |
| `DB_USERNAME` | Username PostgreSQL |
| `DB_PASSWORD` | Password PostgreSQL |
| `DB_NAME` | Nama database |
| `SECRET_KEY` | Secret key untuk JWT |
| `REDIS_HOST` | Host Redis |
| `REDIS_PORT` | Port Redis (default: `6379`) |
| `REDIS_PASSWORD` | Password Redis (kosongkan jika tidak ada) |
| `FRONTEND_URL` | URL frontend untuk CORS |
| `APP_URL` | URL base aplikasi ini |

### 4. Jalankan migrasi database

```bash
# Pastikan database sudah dibuat, lalu jalankan migration
go run cmd/main.go migrate
```

### 5. Jalankan server

```bash
# Development (dengan hot reload menggunakan Air)
air

# Atau tanpa hot reload
go run cmd/main.go
```

Server berjalan di `http://localhost:{PORT}`

---

## Docker

### Build & Run

```bash
# Build image
docker build -t koda-b6-backend .

# Run container
docker run -p 8888:8888 --env-file .env koda-b6-backend
```

Aplikasi akan berjalan di port `8888`.

### Multi-stage Build

Dockerfile menggunakan dua stage:
- **Stage 1 (build)**: Compile Go binary menggunakan `golang:1.25.3-alpine`
- **Stage 2 (run)**: Jalankan binary di image `alpine:latest` yang lebih ringan

---

## API Documentation

Swagger UI tersedia setelah server berjalan:

```
http://localhost:{PORT}/swagger/index.html
```

Untuk regenerate docs setelah ada perubahan:

```bash
swag init -g cmd/main.go
```

---

## License

MIT License

Copyright (c) 2025 VirgilIw

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
