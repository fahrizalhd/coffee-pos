# Coffee POS Dev

Modern coffee shop POS system built with Golang, PostgreSQL, Vue, and Docker.

---

## Tech Stack

### Backend

* Golang
* Gin
* GORM
* PostgreSQL
* Air (Hot Reload)

### Frontend

* Vue 3
* Vite
* Axios

### Infrastructure

* Docker
* Docker Compose

---

# Project Structure

```text
coffee-pos/
│
├── apps/
│   ├── api/        # Golang backend API
│   └── web-pos/    # Vue frontend POS
│
├── docker-compose.yml
└── README.md
```

---

# Backend Structure

```text
apps/api/
│
├── cmd/
│   └── server/
│
├── internal/
│   ├── config/
│   ├── database/
│   ├── middleware/
│   └── modules/
│
├── pkg/
├── .air.toml
├── Dockerfile
└── go.mod
```

---

# Requirements

* Docker Desktop
* Git
* Go 1.25+ (optional for local development)

---

# Getting Started

## Clone Repository

```bash
git clone git@github.com:fahrizalhd/coffee-pos.git
cd coffee-pos
```

---

# Environment Setup

Create:

```text
apps/api/.env
```

Example:

```env
APP_PORT=8080

DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=coffee_pos
DB_SSLMODE=disable
```

---

# Run Project

## Start Docker Services

```bash
docker compose up --build
```

Backend API:

```text
http://localhost:8080
```

PostgreSQL:

```text
localhost:5432
```

---

# Development Workflow

## Start Containers

```bash
docker compose up
```

## Stop Containers

```bash
docker compose down
```

## Rebuild Containers

```bash
docker compose up --build
```

---

# Git Branch Strategy

| Branch | Purpose                  |
| ------ | ------------------------ |
| main   | Stable production branch |
| dev    | Main development branch  |

---

# Planned Features

* Authentication & Authorization
* Product Management
* Order Management
* Payment System
* Inventory Tracking
* Sales Report
* Multi-role Access
* Kitchen Display System (KDS)
* QR Order Support

---

# Author

Developed by fahrizalhd.
